//+build windows

package main

import "os"

func env() {
	os.Setenv("FYNE_FONT", "C:\\Windows\\Fonts\\simkai.ttf")
}
