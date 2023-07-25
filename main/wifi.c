#include "wifi.h"
#include "config.h"
#include <esp_err.h>
#include <esp_log.h>
#include <esp_wifi.h>
#include <stdbool.h>

static const char* TAG       = "WF";
static bool        connected = false;

static void wifi_event_handler(void* arg, esp_event_base_t event_base, int32_t event_id, void* event_data)
{
    switch (event_id) {
    case WIFI_EVENT_STA_START:
        ESP_ERROR_CHECK(esp_wifi_connect());
        break;

    case WIFI_EVENT_STA_CONNECTED:
        ESP_LOGI(TAG, "WIFI_EVENT_STA_CONNECTED");
        connected = true;
        break;

    case WIFI_EVENT_STA_DISCONNECTED:
        ESP_LOGW(TAG, "WIFI_EVENT_STA_DISCONNECTED");
        connected = false;
        ESP_ERROR_CHECK(esp_wifi_connect());
        break;

    default:
        break;
    }
}

void wifi_init()
{
    ESP_ERROR_CHECK(esp_event_handler_register(WIFI_EVENT, ESP_EVENT_ANY_ID, &wifi_event_handler, NULL));

    ESP_ERROR_CHECK(esp_netif_init());
    esp_netif_create_default_wifi_sta();

    wifi_config_t wifi_config = {
        .sta =
            {
                .ssid               = AQ_WIFI_SSID,
                .password           = AQ_WIFI_PASS,
                .threshold.authmode = WIFI_AUTH_WPA2_PSK,
            },
    };
    wifi_init_config_t init_cfg = WIFI_INIT_CONFIG_DEFAULT();
    esp_wifi_init(&init_cfg);
    esp_wifi_set_mode(WIFI_MODE_STA);
    esp_wifi_set_config(ESP_IF_WIFI_STA, &wifi_config);
    ESP_ERROR_CHECK(esp_wifi_start());
}

bool wifi_is_connected() { return connected; }