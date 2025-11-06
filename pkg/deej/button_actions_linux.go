//go:build linux
// +build linux

package deej

import (
	"os/exec"
)

// toggleMuteMaster muta/desmuta o volume master usando pactl (PulseAudio)
// Não requer permissões especiais
func (d *Deej) toggleMuteMaster() {
	cmd := exec.Command("pactl", "set-sink-mute", "@DEFAULT_SINK@", "toggle")

	if err := cmd.Run(); err != nil {
		d.logger.Warnw("Failed to toggle mute", "error", err)
		return
	}

	d.logger.Debug("Toggled master mute")
}

// mediaPlayPause envia play/pause usando playerctl (MPRIS)
// Não requer permissões especiais
func (d *Deej) mediaPlayPause() {
	cmd := exec.Command("playerctl", "play-pause")

	if err := cmd.Run(); err != nil {
		d.logger.Warnw("Failed to send play/pause", "error", err)
		return
	}

	d.logger.Debug("Sent play/pause")
}

// mediaNext envia próxima faixa usando playerctl (MPRIS)
// Não requer permissões especiais
func (d *Deej) mediaNext() {
	cmd := exec.Command("playerctl", "next")

	if err := cmd.Run(); err != nil {
		d.logger.Warnw("Failed to send next", "error", err)
		return
	}

	d.logger.Debug("Sent next track")
}

// mediaPrevious envia faixa anterior usando playerctl (MPRIS)
// Não requer permissões especiais
func (d *Deej) mediaPrevious() {
	cmd := exec.Command("playerctl", "previous")

	if err := cmd.Run(); err != nil {
		d.logger.Warnw("Failed to send previous", "error", err)
		return
	}

	d.logger.Debug("Sent previous track")
}