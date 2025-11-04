//go:build windows
// +build windows

package deej

import (
	"syscall"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	keybd_event      = user32.NewProc("keybd_event")
)

const (
	KEYEVENTF_EXTENDEDKEY = 0x0001
	KEYEVENTF_KEYUP       = 0x0002
)

// sendMediaKey envia uma tecla de mídia do Windows
func (d *Deej) sendMediaKey(vkCode byte) {
	// Key down
	keybd_event.Call(
		uintptr(vkCode),
		0,
		uintptr(KEYEVENTF_EXTENDEDKEY),
		0,
	)
	
	// Key up
	keybd_event.Call(
		uintptr(vkCode),
		0,
		uintptr(KEYEVENTF_EXTENDEDKEY|KEYEVENTF_KEYUP),
		0,
	)
	
	d.logger.Debugw("Sent media key", "vkCode", vkCode)
}

// toggleMuteMaster muta/desmuta o volume master
func (d *Deej) toggleMuteMaster() {
	// Usa a tecla de mute do volume (VK_VOLUME_MUTE = 0xAD)
	d.sendMediaKey(0xAD)
	d.logger.Debug("Toggled master mute")
}