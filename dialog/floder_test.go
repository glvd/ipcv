package dialog

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"fyne.io/fyne"
	"fyne.io/fyne/test"
	"fyne.io/fyne/widget"
)

func TestEffectiveStartingDir2(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skipf("os.UserHomeDir) failed, cannot run this test on this system, error was '%s'", err)
	}

	dialog := &FolderDialog{}

	// test that we get $HOME when running with the default struct values
	res := dialog.effectiveStartingDir()
	expect := home
	if !comparePaths(t, res, expect) {
		t.Errorf("Expected effectiveStartingDir() to be '%s', but it was '%s'",
			expect, res)
	}

}

func TestShowFolderOpen(t *testing.T) {
	var chosen string
	var showErr error
	win := test.NewWindow(widget.NewLabel("Content"))
	ShowFolderOpen(func(s string, err error) {
		chosen = s
		showErr = err
	}, win)

	popup := win.Canvas().Overlays().Top().(*widget.PopUp)
	defer win.Canvas().Overlays().Remove(popup)
	assert.NotNil(t, popup)

	ui := popup.Content.(*fyne.Container)
	title := ui.Objects[1].(*widget.Label)
	assert.Equal(t, "Open File", title.Text)

	nameLabel := ui.Objects[2].(*fyne.Container).Objects[1].(*widget.ScrollContainer).Content.(*widget.Label)
	buttons := ui.Objects[2].(*fyne.Container).Objects[0].(*widget.Box)
	open := buttons.Children[1].(*widget.Button)

	files := ui.Objects[3].(*fyne.Container).Objects[1].(*widget.ScrollContainer).Content.(*fyne.Container)
	assert.Greater(t, len(files.Objects), 0)

	fileName := files.Objects[0].(*fileDialogItem).name
	assert.Equal(t, "(Parent)", fileName)
	assert.True(t, open.Disabled())

	var target *fileDialogItem
	for _, icon := range files.Objects {
		if icon.(*fileDialogItem).dir == false {
			target = icon.(*fileDialogItem)
		}
	}

	if target == nil {
		log.Println("Could not find a file in the default directory to tap :(")
		return
	}

	// This will only execute if we have a file in the home path.
	// Until we have a way to set the directory of an open file dialog.
	test.Tap(target)
	assert.Equal(t, filepath.Base(target.path), nameLabel.Text)
	assert.False(t, open.Disabled())

	test.Tap(open)
	assert.Nil(t, win.Canvas().Overlays().Top())
	assert.Nil(t, showErr)
	assert.Equal(t, target.path, chosen)
}
