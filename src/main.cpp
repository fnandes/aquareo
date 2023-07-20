#include "configuration.h"
#include "display.h"
#include "mqtt.h"
#include "temperature.h"
#include "types.h"
#include <Arduino.h>
#include <WiFi.h>
#include <esp_log.h>

static const char* TAG = "MAIN";

// network
WiFiClient wifi;

void setup()
{
    Serial.begin(115200);
    WiFi.begin(WIFI_SSID, WIFI_PASS);

    init_temperature_sensors();
    init_display();
    init_mqtt(&wifi);
}

void loop()
{
    const sensorData_t data{
        .temperature = get_temperature(),
        .ph          = 7, // TODO: make this work
    };

    ESP_LOGI(TAG, "t1: %.1f%%", data.temperature);
    ESP_LOGI(TAG, "ph: %.1f%%", data.ph);

    update_display_data(data);
    publish_mqtt_sensor_data(data);

    delay(2000);
}