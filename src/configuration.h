#pragma once

// base settings
#define AQ_SENSOR_COUNT 3
#define AQ_MAIN_LOOP_TIME 500
// temperature sensor
#define AQ_TP_SENSOR_TIME 3000
#define AQ_TP_SENSOR_BUS_PIN 4
// ph sensor
#define AQ_PH_SENSOR_TIME 10000
// TFT display
#define AQ_DISPLAY_UPDATE_TIME 5000
// WiFi & MQTT
#define AQ_WIFI_SSID "changeit"
#define AQ_WIFI_PWD "changeit"
#define AQ_MQTT_PUBLISH_TIME 10000
#define AQ_MQTT_STATE_TIME 5000
#define AQ_MQTT_DISCOVERY_PREFIX = "homeassistant"
#define AQ_MQTT_DISCOVERY_SEND_ATTEMPT_DELAY 5000
#define AQ_MQTT_CONN_ID "ESP32"
#define AQ_MQTT_CONN_HOST "homeassistant.local"
#define AQ_MQTT_CONN_PORT 1883
