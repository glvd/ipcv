package dialog

import (
	"testing"

	"fyne.io/fyne/widget"
	"github.com/stretchr/testify/assert"
)

func TestFolderDialog_LoadPlaces(t *testing.T) {
	f := &folderDialog{}
	driveLetters := listDrives()
	places := f.loadPlaces()

	assert.Equal(t, len(driveLetters), len(places))
	assert.Equal(t, driveLetters[0], places[0].(*widget.Button).Text)
}
