package dialog

import (
	"image/color"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const (
	folderIconSize      = 64
	folderTextSize      = 24
	folderIconCellWidth = folderIconSize * 1.25
)

type folderDialogItem struct {
	widget.BaseWidget
	picker    *folderDialog
	isCurrent bool

	icon fyne.CanvasObject
	name string
	path string
}

func (i *folderDialogItem) Tapped(_ *fyne.PointEvent) {
	i.picker.setSelected(i)
	i.Refresh()
}

func (i *folderDialogItem) TappedSecondary(_ *fyne.PointEvent) {
}

func (i *folderDialogItem) CreateRenderer() fyne.WidgetRenderer {
	text := widget.NewLabelWithStyle(i.name, fyne.TextAlignCenter, fyne.TextStyle{})
	text.Wrapping = fyne.TextTruncate

	return &folderItemRenderer{item: i,
		img: i.icon, text: text, objects: []fyne.CanvasObject{i.icon, text}}
}

func (i *folderDialogItem) isDirectory() bool {
	return true
}

func (f *folderDialog) newFolderItem(path string) *folderDialogItem {
	var icon fyne.CanvasObject
	icon = canvas.NewImageFromResource(theme.FolderIcon())
	name := filepath.Base(path)
	ret := &folderDialogItem{
		picker: f,
		icon:   icon,
		name:   name,
		path:   path,
	}
	ret.ExtendBaseWidget(ret)
	return ret
}

type folderItemRenderer struct {
	item *folderDialogItem

	img     fyne.CanvasObject
	text    *widget.Label
	objects []fyne.CanvasObject
}

func (s folderItemRenderer) Layout(size fyne.Size) {
	iconAlign := (size.Width - folderIconSize) / 2
	s.img.Resize(fyne.NewSize(folderIconSize, folderIconSize))
	s.img.Move(fyne.NewPos(iconAlign, 0))

	s.text.Resize(fyne.NewSize(size.Width, folderTextSize))
	s.text.Move(fyne.NewPos(0, folderIconSize+theme.Padding()))
}

func (s folderItemRenderer) MinSize() fyne.Size {
	return fyne.NewSize(folderIconSize, folderIconSize+folderTextSize+theme.Padding())
}

func (s folderItemRenderer) Refresh() {
	canvas.Refresh(s.item)
}

func (s folderItemRenderer) BackgroundColor() color.Color {
	if s.item.isCurrent {
		return theme.PrimaryColor()
	}
	return theme.BackgroundColor()
}

func (s folderItemRenderer) Objects() []fyne.CanvasObject {
	return s.objects
}

func (s folderItemRenderer) Destroy() {
}
