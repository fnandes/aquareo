#include "aquareo.h"
#include <esp_err.h>
#include <esp_log.h>
#include <esp_wifi.h>

static c_wifi_stat_e _stat = WF_NOT_INITIALIZED;

static void c_wifi_event_handler(void* arg, esp_event_base_t event_base, int32_t event_id, void* data)
{
    switch (event_id) {
    case WIFI_EVENT_STA_START:
        _stat = WF_CONNECTING;
        ESP_ERROR_CHECK(esp_wifi_connect());
        break;

    case WIFI_EVENT_STA_CONNECTED:
        _stat = WF_CONNECTED;
        break;

    case WIFI_EVENT_STA_DISCONNECTED:
        _stat = WF_ERROR;
        break;
    default:
        break;
    }
}

esp_err_t c_wifi_init(const char* ssid, const char* pass)
{
    ESP_ERROR_CHECK(esp_event_handler_register(WIFI_EVENT, ESP_EVENT_ANY_ID, &c_wifi_event_handler, NULL));
    ESP_ERROR_CHECK(esp_netif_init());

    wifi_config_t wifi_config = {
        .sta =
            {
                .ssid               = ssid,
                .password           = pass,
                .threshold.authmode = WIFI_AUTH_WPA_PSK,
            },
    };
    wifi_init_config_t init_config = WIFI_INIT_CONFIG_DEFAULT();
    ESP_ERROR_CHECK(esp_wifi_init(&init_config));
    ESP_ERROR_CHECK(esp_wifi_set_mode(WIFI_MODE_STA));
    ESP_ERROR_CHECK(esp_wifi_set_config(ESP_IF_WIFI_STA, &wifi_config));

    return esp_wifi_start();
}

inline c_wifi_stat_e c_wifi_get_stat() { return _stat; }