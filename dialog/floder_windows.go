package dialog

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"os"
)

func (f *folderDialog) loadPlaces() []fyne.CanvasObject {
	var places []fyne.CanvasObject

	for _, drive := range listDrives() {
		driveRoot := drive + string(os.PathSeparator) // capture loop var
		places = append(places, widget.NewButton(drive, func() {
			f.setDirectory(driveRoot)
		}))
	}
	return places
}

func folderOpenOSOverride(*FolderDialog) bool {
	return false
}

func folderSaveOSOverride(*FolderDialog) bool {
	return false
}
