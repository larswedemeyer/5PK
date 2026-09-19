// the following Code is used for the esp32 in order to connect it with the raspberry pi and take pictures

/*
#include "esp_camera.h"
#include <WiFi.h>
#include <HTTPClient.h>

// WLAN Daten
const char* ssid = "infopi";
const char* password = "ilovegolang";

// URL des Raspberry Pi Servers
const char* serverUrl = "http://10.42.0.1:5000/upload";

// AI THINKER ESP32-CAM
#define PWDN_GPIO_NUM     32
#define RESET_GPIO_NUM    -1
#define XCLK_GPIO_NUM      0
#define SIOD_GPIO_NUM     26
#define SIOC_GPIO_NUM     27

#define Y9_GPIO_NUM       35
#define Y8_GPIO_NUM       34
#define Y7_GPIO_NUM       39
#define Y6_GPIO_NUM       36
#define Y5_GPIO_NUM       21
#define Y4_GPIO_NUM       19
#define Y3_GPIO_NUM       18
#define Y2_GPIO_NUM        5

#define VSYNC_GPIO_NUM    25
#define HREF_GPIO_NUM     23
#define PCLK_GPIO_NUM     22

const int LED_BUILTIN = 33;

void setup() {

  Serial.begin(115200);

  pinMode(LED_BUILTIN, OUTPUT);
  digitalWrite(LED_BUILTIN, HIGH);

  camera_config_t config;

  config.ledc_channel = LEDC_CHANNEL_0;
  config.ledc_timer = LEDC_TIMER_0;

  config.pin_d0 = Y2_GPIO_NUM;
  config.pin_d1 = Y3_GPIO_NUM;
  config.pin_d2 = Y4_GPIO_NUM;
  config.pin_d3 = Y5_GPIO_NUM;
  config.pin_d4 = Y6_GPIO_NUM;
  config.pin_d5 = Y7_GPIO_NUM;
  config.pin_d6 = Y8_GPIO_NUM;
  config.pin_d7 = Y9_GPIO_NUM;

  config.pin_xclk = XCLK_GPIO_NUM;
  config.pin_pclk = PCLK_GPIO_NUM;
  config.pin_vsync = VSYNC_GPIO_NUM;
  config.pin_href = HREF_GPIO_NUM;

  config.pin_sscb_sda = SIOD_GPIO_NUM;
  config.pin_sscb_scl = SIOC_GPIO_NUM;

  config.pin_pwdn = PWDN_GPIO_NUM;
  config.pin_reset = RESET_GPIO_NUM;

  config.xclk_freq_hz = 20000000;
  config.pixel_format = PIXFORMAT_JPEG;

  config.frame_size = FRAMESIZE_QQVGA;
  config.jpeg_quality = 25;
  config.fb_count = 1;

  Serial.println("Initialisiere Kamera...");

  esp_err_t err = esp_camera_init(&config);

  if (err != ESP_OK) {
    Serial.printf("Kamerafehler: 0x%x\n", err);
    while (true) {
      delay(1000);
    }
  }

  Serial.println("Kamera OK");

  Serial.print("Verbinde mit WLAN ");

  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print(".");
  }

  Serial.println();
  Serial.println("WLAN verbunden");

  Serial.print("ESP32 IP: ");
  Serial.println(WiFi.localIP());
}

void loop() {

  digitalWrite(LED_BUILTIN, LOW);

  camera_fb_t *fb = esp_camera_fb_get();

  if (!fb) {
    Serial.println("Foto fehlgeschlagen");
    digitalWrite(LED_BUILTIN, HIGH);
    delay(1000);
    return;
  }

  Serial.print("Bildgröße: ");
  Serial.print(fb->len);
  Serial.println(" Bytes");

  HTTPClient http;

  http.begin(serverUrl);
  http.addHeader("Content-Type", "image/jpeg");

  int httpCode = http.POST(fb->buf, fb->len);

  Serial.print("HTTP-Code: ");
  Serial.println(httpCode);

  if (httpCode > 0) {
    String response = http.getString();

    Serial.print("Antwort: ");
    Serial.println(response);
  }

  http.end();

  esp_camera_fb_return(fb);

  digitalWrite(LED_BUILTIN, HIGH);

  delay(1000);
}
*/