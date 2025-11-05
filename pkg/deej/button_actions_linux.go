//go:build linux
// +build linux

package deej

import (
	"os/exec"
)

// sendMediaKey envia uma tecla de mídia usando xdotool
func (d *Deej) sendMediaKey(keyName string) {
	cmd := exec.Command("xdotool", "key", keyName)
	
	if err := cmd.Run(); err != nil {
		d.logger.Warnw("Failed to send media key", "key", keyName, "error", err)
		return
	}
	
	d.logger.Debugw("Sent media key", "key", keyName)
}

// toggleMuteMaster muta/desmuta o volume master
func (d *Deej) toggleMuteMaster() {
	d.sendMediaKey("XF86AudioMute")
	d.logger.Debug("Toggled master mute")
}

// mediaPlayPause envia play/pause
func (d *Deej) mediaPlayPause() {
	d.sendMediaKey("XF86AudioPlay")
	d.logger.Debug("Sent play/pause")
}

// mediaNext envia próxima faixa
func (d *Deej) mediaNext() {
	d.sendMediaKey("XF86AudioNext")
	d.logger.Debug("Sent next track")
}

// mediaPrevious envia faixa anterior
func (d *Deej) mediaPrevious() {
	d.sendMediaKey("XF86AudioPrev")
	d.logger.Debug("Sent previous track")
}