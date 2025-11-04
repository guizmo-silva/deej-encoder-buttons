package deej

import (
	"fmt"
	"strings"
)

// handleButtonPress processa os comandos de botão recebidos do Arduino
func (d *Deej) handleButtonPress(buttonCommand string) {
	// Remove espaços em branco
	buttonCommand = strings.TrimSpace(buttonCommand)
	
	// Verifica se é um comando de botão (formato: BTN:0, BTN:1, etc)
	if !strings.HasPrefix(buttonCommand, "BTN:") {
		return
	}
	
	// Extrai o número do botão
	buttonNumStr := strings.TrimPrefix(buttonCommand, "BTN:")
	buttonNum := 0
	
	// Converte string para int
	if _, err := fmt.Sscanf(buttonNumStr, "%d", &buttonNum); err != nil {
		d.logger.Warnw("Failed to parse button number", "command", buttonCommand, "error", err)
		return
	}

	// DEBUG: Mostra o mapeamento completo
	d.logger.Debugw("Button mapping loaded", "mapping", d.config.ButtonMapping)
	
	// Busca a ação mapeada para este botão
	action, exists := d.config.ButtonMapping[buttonNum]
	if !exists {
		d.logger.Debugw("Button pressed but not mapped", "button", buttonNum)
		return
	}
	
	// Executa a ação
	d.executeButtonAction(action)
}

// executeButtonAction executa a ação configurada para o botão
func (d *Deej) executeButtonAction(action string) {
	d.logger.Debugw("Executing button action", "action", action)
	
	switch action {
	case "mute_master":
		d.toggleMuteMaster()
	case "media_play_pause":
		d.sendMediaKey(0xB3) // VK_MEDIA_PLAY_PAUSE
	case "media_next":
		d.sendMediaKey(0xB0) // VK_MEDIA_NEXT_TRACK
	case "media_previous":
		d.sendMediaKey(0xB1) // VK_MEDIA_PREV_TRACK
	default:
		d.logger.Warnw("Unknown button action", "action", action)
	}
}