#include "mqtt.h"
#include "config.h"
#include <esp_log.h>
#include <esp_timer.h>
#include <mqtt_client.h>
#include <stdbool.h>
#include <stdio.h>

static const char* TAG = "MQ";

static esp_mqtt_client_handle_t mqtt_client = NULL;
static bool                     connected   = false;

static void mqtt_event_handler(void* handler_args, esp_event_base_t base, int32_t event_id, void* event_data)
{
    ESP_LOGD(TAG, "mqtt_event_handler");
    // esp_mqtt_event_handle_t event = event_data;
    // esp_mqtt_client_handle_t client = event->client;

    switch ((esp_mqtt_event_id_t)event_id) {
    case MQTT_EVENT_CONNECTED:
        ESP_LOGI(TAG, "MQTT_EVENT_CONNECTED");
        connected = true;
        break;

    case MQTT_EVENT_DISCONNECTED:
        ESP_LOGW(TAG, "MQTT_EVENT_DISCONNECTED");
        connected = false;
        break;

    default:
        break;
    }
}

void mqtt_init()
{
    ESP_LOGI(TAG, "init");

    esp_mqtt_client_config_t mqtt_cfg = {
        .broker.address.uri = AQ_MQTT_URI,
        .credentials =
            {
                .username                = AQ_MQTT_USER,
                .authentication.password = AQ_MQTT_PASS,
            },
    };
    mqtt_client = esp_mqtt_client_init(&mqtt_cfg);
    esp_mqtt_client_register_event(mqtt_client, ESP_EVENT_ANY_ID, mqtt_event_handler, NULL);
    esp_mqtt_client_start(mqtt_client);
}

bool mqtt_is_connected() { return connected; }

static void mqtt_publish_sensor_topic(const char* topic_name, const float value)
{
    int msg_id;
    if (mqtt_is_connected()) {
        char v_buff[10];
        snprintf(v_buff, sizeof(v_buff), "%f", value);

        msg_id = esp_mqtt_client_publish(mqtt_client, topic_name, v_buff, 0, 1, 0);
        ESP_LOGI(TAG, "sent publish successful, msg_id=%d", msg_id);
    }
}

void mqtt_publish_state(const aquareo_state_t* state)
{
    ESP_LOGI(TAG, "publish state");
    for (int i = 0; i < AQ_ONEWIRE_MAX_DS18B20; i++) {
        if (state->current_temp[i] > 0) {
            char topic_str[20];
            snprintf(topic_str, sizeof(topic_str), "temperature_%d", i);
            mqtt_publish_sensor_topic(topic_str, state->current_temp[i]);
        }
    }
}