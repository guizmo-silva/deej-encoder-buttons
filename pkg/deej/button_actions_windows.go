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

	// Virtual Key Codes para controle de volume
	VK_VOLUME_MUTE = 0xAD
	VK_VOLUME_DOWN = 0xAE
	VK_VOLUME_UP   = 0xAF
)

// SendMediaKey envia uma tecla de mídia do Windows
func (d *Deej) SendMediaKey(vkCode byte) {
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
	d.SendMediaKey(VK_VOLUME_MUTE)
	d.logger.Debug("Toggled master mute")
}

// mediaPlayPause envia play/pause usando tecla de mídia do Windows
func (d *Deej) mediaPlayPause() {
	d.SendMediaKey(0xB3) // VK_MEDIA_PLAY_PAUSE
	d.logger.Debug("Sent play/pause")
}

// mediaNext envia próxima faixa usando tecla de mídia do Windows
func (d *Deej) mediaNext() {
	d.SendMediaKey(0xB0) // VK_MEDIA_NEXT_TRACK
	d.logger.Debug("Sent next track")
}

// mediaPrevious envia faixa anterior usando tecla de mídia do Windows
func (d *Deej) mediaPrevious() {
	d.SendMediaKey(0xB1) // VK_MEDIA_PREV_TRACK
	d.logger.Debug("Sent previous track")
}

// sendVolumeKey envia uma tecla de volume (sem método receiver)
// Útil para ser chamado de outros pacotes como session_windows.go
func sendVolumeKey(vkCode byte) {
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
}