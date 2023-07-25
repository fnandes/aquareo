#pragma once

// system
#define AQ_SYS_CLOCK 5000
#define AQ_ONEWIRE_BUS_GPIO 4
#define AQ_ONEWIRE_MAX_DS18B20 2

// wifi
#define AQ_WIFI_SSID "FRITZ!Box 7530 SM_EXT"
#define AQ_WIFI_PASS "72229058186761679170"
// mqtt
#define AQ_MQTT_URI "mqtt://raspberrypi.local:1883"
#define AQ_MQTT_USER "mqtt-user"
#define AQ_MQTT_PASS "piranha"

typedef struct {
    float current_temp[AQ_ONEWIRE_MAX_DS18B20];
} aquareo_state_t;