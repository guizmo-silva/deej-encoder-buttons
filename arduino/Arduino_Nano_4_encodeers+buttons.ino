// ========================================
// DEEJ WITH 4 ROTARY ENCODERS + BUTTONS
// ========================================

const int NUM_SLIDERS = 4;
const int analogInputs[NUM_SLIDERS] = {A0, A1, A2, A3};

// PINOS DOS ENCODERS
#define ENC1_CLK 2
#define ENC1_DT  3
#define ENC1_SW  4

#define ENC2_CLK 5
#define ENC2_DT  6
#define ENC2_SW  7

#define ENC3_CLK 8
#define ENC3_DT  9
#define ENC3_SW  10

#define ENC4_CLK 11
#define ENC4_DT  12
#define ENC4_SW  A0

// ====================================
// GLOBAL VARIABLES
// ====================================
int enc1Pos = 512; 
int enc2Pos = 512;
int enc3Pos = 512;
int enc4Pos = 512;

int enc1LastCLK, enc2LastCLK, enc3LastCLK, enc4LastCLK;

// Debounce variables
unsigned long lastBtn1Press = 0;
unsigned long lastBtn2Press = 0;
unsigned long lastBtn3Press = 0;
unsigned long lastBtn4Press = 0;
const unsigned long debounceDelay = 300;

// ====================================
// ENCODERS READING FUNCTIONS
// ====================================

void readEncoder1() {
  int currentCLK = digitalRead(ENC1_CLK);

  if (currentCLK != enc1LastCLK && currentCLK == LOW) {
    if (digitalRead(ENC1_DT) != currentCLK) {
      enc1Pos += 10;
      if (enc1Pos > 1023) enc1Pos = 1023;
    } else {
      enc1Pos -= 10;
      if (enc1Pos < 0) enc1Pos = 0;
    }
  }
  enc1LastCLK = currentCLK;
}

void readEncoder2() {
  int currentCLK = digitalRead(ENC2_CLK);

  if (currentCLK != enc2LastCLK && currentCLK == LOW) {
    if (digitalRead(ENC2_DT) != currentCLK) {
      enc2Pos += 10;
      if (enc2Pos > 1023) enc2Pos = 1023;
    } else {
      enc2Pos -= 10;
      if (enc2Pos < 0) enc2Pos = 0;
    }
  }
  enc2LastCLK = currentCLK;
}

void readEncoder3() {
  int currentCLK = digitalRead(ENC3_CLK);

  if (currentCLK != enc3LastCLK && currentCLK == LOW) {
    if (digitalRead(ENC3_DT) != currentCLK) {
      enc3Pos += 10;
      if (enc3Pos > 1023) enc3Pos = 1023;
    } else {
      enc3Pos -= 10;
      if (enc3Pos < 0) enc3Pos = 0;
    }
  }
  enc3LastCLK = currentCLK;
}

void readEncoder4() {
  int currentCLK = digitalRead(ENC4_CLK);

  if (currentCLK != enc4LastCLK && currentCLK == LOW) {
    if (digitalRead(ENC4_DT) != currentCLK) {
      enc4Pos += 10;
      if (enc4Pos > 1023) enc4Pos = 1023;
    } else {
      enc4Pos -= 10;
      if (enc4Pos < 0) enc4Pos = 0;
    }
  }
  enc4LastCLK = currentCLK;
}

// ====================================
// BUTTONS FUNCTIONS
// ====================================

void checkButtons() {
  if (digitalRead(ENC1_SW) == LOW) {
    if (millis() - lastBtn1Press > debounceDelay) {
      Serial.println("BTN:0");
      lastBtn1Press = millis();
    }
  }
  
  if (digitalRead(ENC2_SW) == LOW) {
    if (millis() - lastBtn2Press > debounceDelay) {
      Serial.println("BTN:1");
      lastBtn2Press = millis();
    }
  }
  
  if (digitalRead(ENC3_SW) == LOW) {
    if (millis() - lastBtn3Press > debounceDelay) {
      Serial.println("BTN:2");
      lastBtn3Press = millis();
    }
  }
  
  if (digitalRead(ENC4_SW) == LOW) {
    if (millis() - lastBtn4Press > debounceDelay) {
      Serial.println("BTN:3");
      lastBtn4Press = millis();
    }
  }
}

// ====================================
// SETUP
// ====================================
void setup() {
  Serial.begin(9600);
  
  pinMode(ENC1_CLK, INPUT_PULLUP);
  pinMode(ENC1_DT, INPUT_PULLUP);
  pinMode(ENC1_SW, INPUT_PULLUP);
  
  pinMode(ENC2_CLK, INPUT_PULLUP);
  pinMode(ENC2_DT, INPUT_PULLUP);
  pinMode(ENC2_SW, INPUT_PULLUP);
  
  pinMode(ENC3_CLK, INPUT_PULLUP);
  pinMode(ENC3_DT, INPUT_PULLUP);
  pinMode(ENC3_SW, INPUT_PULLUP);
  
  pinMode(ENC4_CLK, INPUT_PULLUP);
  pinMode(ENC4_DT, INPUT_PULLUP);
  pinMode(ENC4_SW, INPUT_PULLUP);
  
  enc1LastCLK = digitalRead(ENC1_CLK);
  enc2LastCLK = digitalRead(ENC2_CLK);
  enc3LastCLK = digitalRead(ENC3_CLK);
  enc4LastCLK = digitalRead(ENC4_CLK);
  
  delay(1000);
  for (int i = 0; i < 10; i++) {
    Serial.println("512|512|512|512");
    delay(50);
  }
}

// ====================================
// LOOP
// ====================================
void loop() {
  static int lastEnc1 = -1;
  static int lastEnc2 = -1;
  static int lastEnc3 = -1;
  static int lastEnc4 = -1;
  
  // Reads encoders
  readEncoder1();
  readEncoder2();
  readEncoder3();
  readEncoder4();
  
  // Check buttons
  checkButtons();
  
  // Send values only if changed
  if (enc1Pos != lastEnc1 || enc2Pos != lastEnc2 || 
      enc3Pos != lastEnc3 || enc4Pos != lastEnc4) {
    
    String output = String(enc1Pos) + "|" + 
                    String(enc2Pos) + "|" + 
                    String(enc3Pos) + "|" + 
                    String(enc4Pos);
    
    Serial.println(output);
    
    lastEnc1 = enc1Pos;
    lastEnc2 = enc2Pos;
    lastEnc3 = enc3Pos;
    lastEnc4 = enc4Pos;
  }
  
  delay(1);
}