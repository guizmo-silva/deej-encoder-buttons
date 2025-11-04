# Deej with Rotary Encoder Button Support

This is a fork of the [original Deej](https://github.com/omriharel/deej) with additional functionality for **rotary encoder buttons**.

## 📝 About this Fork

This project adds native support for rotary encoder buttons, allowing you to control media functions (play/pause, next, previous, mute) directly from your Deej device.

**⚠️ Important Notice:**
- I am not a professional developer
- This version was created with AI assistance
- Tested on: **Windows 11** with **Arduino Nano v3.0 (atmega328p/168p)**
- If you encounter bugs or issues, feel free to fork and adapt it to your needs

## ✨ New Features

- **Button configuration via YAML** - Define button actions without reprogramming the Arduino
- **Media commands** - Play/pause, next track, previous track, master mute
- **Easy customization** - Change functions by editing only the `config.yaml`

## 🔧 Configuration

### 1. Arduino Code

An example file is included in the project: `arduino/Arduino_Nano_4_encodeers+buttons.ino`

**Important:** Buttons must send commands in this format:
```
BTN:0  // Encoder 1 button
BTN:1  // Encoder 2 button
BTN:2  // Encoder 3 button
BTN:3  // Encoder 4 button
```

### 2. config.yaml Configuration

Add this section to your `config.yaml`:

```yaml
button_mapping:
  0: mute_master        # Mute/unmute system
  1: media_play_pause   # Play/Pause
  2: media_previous     # Previous track
  3: media_next         # Next track
```

## ⚠️ IMPORTANT - Arduino Not Recognized Issues

If Deej doesn't recognize your Arduino, the problem might be with the **CH340 driver**:

1. [Download and install the CH340 driver](https://learn.sparkfun.com/tutorials/how-to-install-ch340-drivers/all)
2. If already installed and still not working: **completely uninstall and reinstall**
3. Restart your computer after installation

## 💬 Questions about assembly?

For questions about building the Deej hardware, join the [official Deej Discord](https://discord.com/invite/nf88NJu)

#### Special thanks to _TheGiantSwede_ form the Deej Discord who helped me with the Arduino code with encoders

## 📄 License

This project maintains the MIT license from the original Deej.

---

 # 🇧🇷 Deej com Suporte a Botões para Rotary Encoders

Este é um fork do [Deej original](https://github.com/omriharel/deej) com funcionalidade adicional para **botões de rotary encoders**.

## 📝 Sobre este Fork

Este projeto adiciona suporte nativo aos botões dos rotary encoders, permitindo que você controle funções de mídia (play/pause, next, previous, mute) diretamente pelo seu dispositivo Deej.

**⚠️ Aviso Importante:**
- Não sou desenvolvedor profissional
- Esta versão foi criada com ajuda de IA
- Testada em: **Windows 11** com **Arduino Nano v3.0 (atmega328p/168p)**
- Se encontrar bugs ou problemas, sinta-se à vontade para fazer um fork e adaptar às suas necessidades

## ✨ Novas Funcionalidades

- **Configuração de botões via YAML** - Defina as ações dos botões sem reprogramar o Arduino
- **Comandos de mídia** - Play/pause, próxima música, música anterior, mute master
- **Fácil personalização** - Mude as funções editando apenas o `config.yaml`

## 🔧 Configuração

### 1. Código do Arduino

Um arquivo de exemplo está incluído no projeto: `arduino/Arduino_Nano_4_encodeers+buttons.ino`

**Importante:** Os botões devem enviar comandos no formato:
```
BTN:0  // Botão do encoder 1
BTN:1  // Botão do encoder 2
BTN:2  // Botão do encoder 3
BTN:3  // Botão do encoder 4
```

### 2. Configuração do config.yaml

Adicione esta seção ao seu `config.yaml`:

```yaml
button_mapping:
  0: mute_master        # Muta/desmuta o sistema
  1: media_play_pause   # Play/Pause
  2: media_previous     # Música anterior
  3: media_next         # Próxima música
```

## ⚠️ IMPORTANTE - Problemas com Arduino não reconhecido

Se o Deej não reconhecer seu Arduino, o problema pode ser com o **driver CH340**:

1. [Baixe e instale o driver CH340](https://learn.sparkfun.com/tutorials/how-to-install-ch340-drivers/all)
2. Se já tiver instalado e ainda assim não funcionar: **desinstale completamente e reinstale**
3. Reinicie o computador após a instalação

## 💬 Dúvidas sobre montagem?

Para dúvidas sobre como montar o hardware do Deej, entre no [Discord oficial do Deej](https://discord.com/invite/nf88NJu)

#### Agradecimento especial ao user _TheGiantSwede_ do Discord do Deej pela ajuda com o código de rotary encoders no Arduino

## 📄 Licença

Este projeto mantém a licença MIT do Deej original.