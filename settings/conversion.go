package settings

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/config"
	"github.com/glvd/ipcv/dialog"
)

func (s *Settings) makeConversionTab(w fyne.Window) *widget.TabItem {
	acc := s.makeAccSetting(s.config.Conversion.HardwareAcceleration)
	ffmpeg := s.makeFFMpeg(w)
	return widget.NewTabItem(s.lang.Conversion.Title, widget.NewVBox(acc, ffmpeg))
}
func (s *Settings) makeAccSetting(name string) fyne.CanvasObject {
	label := widget.NewLabel(s.lang.System.Accelerate)
	names := []string{config.CPUAcceleration, config.AMDAcceleration, config.NvidiaAcceleration, config.MacAcceleration}
	slt := widget.NewSelect(names, func(v string) {
		s.chooseAcc(v)
	})
	slt.SetSelected(name)
	return fyne.NewContainerWithLayout(layout.NewGridLayout(2), label, slt)
}
func (s *Settings) chooseAcc(v string) {
	s.config.Conversion.HardwareAcceleration = v
}

func (s *Settings) makeFFMpeg(w fyne.Window) fyne.CanvasObject {
	//label := widget.NewLabel(c.lang.Input.Label)
	text := widget.NewEntry()
	name := s.config.Conversion.FFMPEG
	if len(s.config.Conversion.FFMPEG) > 60 {
		name = name[0:60] + "..."
	}
	text.SetText(name)
	text.Disable()
	button := widget.NewButton(s.lang.Conversion.FFMpegButton, func() {
		dialog.ShowFileOpen(func(str string) {
			s.config.Conversion.FFMPEG = str
			name := s.config.Conversion.FFMPEG
			if len(s.config.Conversion.FFMPEG) > 60 {
				name = name[0:60] + "..."
			}
			text.SetText(name)
		}, w)
	})
	//box := widget.NewHBox(layout.NewSpacer(), button)
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), text, button)
}

func (s *Settings) makeFFmpegVersionLabel(w fyne.Window) fyne.CanvasObject {
	text := widget.NewEntry()
	name := s.config.Conversion.FFMPEG
	if len(s.config.Conversion.FFMPEG) > 60 {
		name = name[0:60] + "..."
	}
	text.SetText(name)
	text.Disable()
	button := widget.NewButton(s.lang.Conversion.FFMpegButton, func() {
		dialog.ShowFileOpen(func(str string) {
			s.config.Conversion.FFMPEG = str
			name := s.config.Conversion.FFMPEG
			if len(s.config.Conversion.FFMPEG) > 60 {
				name = name[0:60] + "..."
			}
			text.SetText(name)
		}, w)
	})
	//box := widget.NewHBox(layout.NewSpacer(), button)
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), text, button)
}
