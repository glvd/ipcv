package settings

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/config"
)

func (s *Settings) makeConversionTab(w fyne.Window) *widget.TabItem {
	acc := s.makeAccSetting(s.config.Conversion.HardwareAcceleration)
	return widget.NewTabItem(s.lang.Conversion.Title, widget.NewVBox(acc))
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
