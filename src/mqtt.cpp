#include "mqtt.h"
#include "configuration.h"
#include "types.h"
#include <Arduino.h>
#include <PubSubClient.h>
#include <WiFi.h>
#include <esp_log.h>

static const char* TAG = "MQTT";

static PubSubClient* client;
static unsigned long last_run = 0;

void init_mqtt(WiFiClient* wifi)
{
    ESP_LOGI(TAG, "init");

    client = new PubSubClient(*wifi);

    if (!wifi->connected()) {
        ESP_LOGE(TAG, "WIFI not connected!");
        return;
    }

    if (client->connect(AQ_MQTT_CONN_ID)) {
        ESP_LOGI(TAG, "connected");
    }
}

void publish_mqtt_sensor_data(const sensorData_t data)
{
    unsigned long ticks = millis();
    if (ticks - last_run <= 5000) {
        last_run = ticks;

        if (client->connected()) {
            char tmp_str[8], ph_str[8];
            dtostrf(data.temperature, 1, 2, tmp_str);
            dtostrf(data.ph, 1, 2, ph_str);

            client->publish("aquareo/sensor/temperature", tmp_str);
            client->publish("aquareo/sensor/ph", ph_str);
        }
    }
}