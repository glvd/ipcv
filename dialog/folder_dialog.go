package dialog

import (
	"fyne.io/fyne/dialog"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type textWidget interface {
	fyne.Widget
	SetText(string)
}

type folderDialog struct {
	file       *FolderDialog
	fileName   textWidget
	dismiss    *widget.Button
	open       *widget.Button
	breadcrumb *widget.Box
	files      *fyne.Container
	fileScroll *widget.ScrollContainer

	win      *widget.PopUp
	selected *folderDialogItem
	dir      string
}

// FolderDialog is a dialog containing a file picker for use in opening or saving files.
type FolderDialog struct {
	//save             bool
	callback         interface{}
	onClosedCallback func(bool)
	filter           storage.FileFilter
	parent           fyne.Window
	dialog           *folderDialog
	dismissText      string
}

// Declare conformity to Dialog interface
var _ dialog.Dialog = (*FolderDialog)(nil)

func (f *folderDialog) makeUI() fyne.CanvasObject {

	f.fileName = widget.NewLabel("")

	label := "Open"
	f.open = widget.NewButton(label, func() {
		if f.file.callback == nil {
			f.win.Hide()
			if f.file.onClosedCallback != nil {
				f.file.onClosedCallback(false)
			}
			return
		}

		if f.selected != nil {
			callback := f.file.callback.(func(string, error))
			f.win.Hide()
			if f.file.onClosedCallback != nil {
				f.file.onClosedCallback(true)
			}
			callback(f.selected.path, nil)
		}
	})
	f.open.Style = widget.PrimaryButton
	f.open.Disable()
	dismissLabel := "Cancel"
	if f.file.dismissText != "" {
		dismissLabel = f.file.dismissText
	}
	f.dismiss = widget.NewButton(dismissLabel, func() {
		f.win.Hide()
		if f.file.onClosedCallback != nil {
			f.file.onClosedCallback(false)
		}
		if f.file.callback != nil {
			f.file.callback.(func(string, error))("", nil)
		}
	})
	buttons := widget.NewHBox(f.dismiss, f.open)
	footer := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, buttons),
		buttons, widget.NewHScrollContainer(f.fileName))

	f.files = fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(fileIconCellWidth,
		fileIconSize+theme.Padding()+fileTextSize)),
	)
	f.fileScroll = widget.NewScrollContainer(f.files)
	verticalExtra := int(float64(fileIconSize) * 0.25)
	f.fileScroll.SetMinSize(fyne.NewSize(fileIconCellWidth*2+theme.Padding(),
		(fileIconSize+fileTextSize)+theme.Padding()*2+verticalExtra))

	f.breadcrumb = widget.NewHBox()
	scrollBread := widget.NewScrollContainer(f.breadcrumb)
	body := fyne.NewContainerWithLayout(layout.NewBorderLayout(scrollBread, nil, nil, nil),
		scrollBread, f.fileScroll)
	header := widget.NewLabelWithStyle(label+" File", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	favorites := widget.NewGroup("Favorites", f.loadFavorites()...)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(header, footer, favorites, nil),
		favorites, header, footer, body)
}

func (f *folderDialog) loadFavorites() []fyne.CanvasObject {
	home, _ := os.UserHomeDir()
	places := []fyne.CanvasObject{
		widget.NewButton("Home", func() {
			f.setDirectory(home)
		}),
		widget.NewButton("Documents", func() {
			f.setDirectory(filepath.Join(home, "Documents"))
		}),
		widget.NewButton("Downloads", func() {
			f.setDirectory(filepath.Join(home, "Downloads"))
		}),
	}

	places = append(places, f.loadPlaces()...)
	return places
}

func (f *folderDialog) refreshDir(dir string) {
	f.files.Objects = nil

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fyne.LogError("Unable to read path "+dir, err)
		return
	}

	var icons []fyne.CanvasObject
	parent := filepath.Dir(dir)
	if parent != dir {
		fi := &folderDialogItem{picker: f, icon: canvas.NewImageFromResource(theme.FolderOpenIcon()),
			name: "(Parent)", path: filepath.Dir(dir)}
		fi.ExtendBaseWidget(fi)
		icons = append(icons, fi)
	}
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		itemPath := filepath.Join(dir, file.Name())
		icons = append(icons, f.newFolderItem(itemPath))

	}

	f.files.Objects = icons
	f.files.Refresh()
	f.fileScroll.Offset = fyne.NewPos(0, 0)
	f.fileScroll.Refresh()
}

func (f *folderDialog) setDirectory(dir string) {
	f.setSelected(nil)
	f.dir = dir

	f.breadcrumb.Children = nil
	buildDir := filepath.VolumeName(dir)
	for i, d := range strings.Split(dir, string(filepath.Separator)) {
		if d == "" {
			if i > 0 { // what we get if we split "/"
				break
			}
			buildDir = "/"
			d = "/"
		} else if i > 0 {
			buildDir = filepath.Join(buildDir, d)
		} else {
			d = buildDir
			buildDir = d + string(os.PathSeparator)
		}

		newDir := buildDir
		f.breadcrumb.Append(
			widget.NewButton(d, func() {
				f.setDirectory(newDir)
			}),
		)
	}

	f.refreshDir(dir)
}

func (f *folderDialog) setSelected(file *folderDialogItem) {
	if f.selected != nil {
		f.selected.isCurrent = false
		f.selected.Refresh()
	}
	if file != nil && file.isDirectory() {
		f.setDirectory(file.path)
	}
	f.selected = file

	if file == nil || file.path == "" {
		f.fileName.SetText("")
		f.open.Disable()
	} else {
		file.isCurrent = true
		f.fileName.SetText(filepath.Base(file.path))
		f.open.Enable()
	}
}

// effectiveStartingDir calculates the directory at which the file dialog
// should open, based on the values of  CWD, home, and any error conditions
// which occur.
//
// Order of precedence is:
//
// * os.UserHomeDir()
// * "/" (should be filesystem root on all supported platforms)
func (f *FolderDialog) effectiveStartingDir() string {

	// Try home dir
	dir, err := os.UserHomeDir()
	if err == nil {
		return dir
	}
	fyne.LogError("Could not load user home dir", err)

	return "/"
}

func showFile(file *FolderDialog) *folderDialog {
	d := &folderDialog{file: file}
	ui := d.makeUI()

	d.setDirectory(file.effectiveStartingDir())

	size := ui.MinSize().Add(fyne.NewSize(fileIconCellWidth*2+theme.Padding()*4,
		(fileIconSize+fileTextSize)+theme.Padding()*4))

	d.win = widget.NewModalPopUp(ui, file.parent.Canvas())
	d.win.Resize(size)

	d.win.Show()
	return d
}

// Show shows the file dialog.
func (f *FolderDialog) Show() {
	if fileOpenOSOverride(f) {
		return
	}
	if f.dialog != nil {
		f.dialog.win.Show()
		return
	}
	f.dialog = showFile(f)
}

// Hide hides the file dialog.
func (f *FolderDialog) Hide() {
	if f.dialog == nil {
		return
	}
	f.dialog.win.Hide()
	if f.onClosedCallback != nil {
		f.onClosedCallback(false)
	}
}

// SetDismissText allows custom text to be set in the confirmation button
func (f *FolderDialog) SetDismissText(label string) {
	if f.dialog == nil {
		return
	}
	f.dialog.dismiss.SetText(label)
	widget.Refresh(f.dialog.win)
}

// SetOnClosed sets a callback function that is called when
// the dialog is closed.
func (f *FolderDialog) SetOnClosed(closed func()) {
	if f.dialog == nil {
		return
	}
	// If there is already a callback set, remember it and call both.
	originalCallback := f.onClosedCallback

	f.onClosedCallback = func(response bool) {
		closed()
		if originalCallback != nil {
			originalCallback(response)
		}
	}
}

// SetFilter sets a filter for limiting files that can be chosen in the file dialog.
func (f *FolderDialog) SetFilter(filter storage.FileFilter) {
	f.filter = filter
	if f.dialog != nil {
		f.dialog.refreshDir(f.dialog.dir)
	}
}

// NewFolderOpen creates a file dialog allowing the user to choose a file to open.
// The dialog will appear over the window specified when Show() is called.
func NewFolderOpen(callback func(string, error), parent fyne.Window) *FolderDialog {
	dialog := &FolderDialog{callback: callback, parent: parent}
	return dialog
}

// ShowFloderOpen creates and shows a file dialog allowing the user to choose a file to open.
// The dialog will appear over the window specified.
func ShowFloderOpen(callback func(string, error), parent fyne.Window) {
	dialog := NewFolderOpen(callback, parent)
	if fileOpenOSOverride(dialog) {
		return
	}
	dialog.Show()
}
