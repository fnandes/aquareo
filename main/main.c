#include "aquareo.h"
#include <freertos/FreeRTOS.h>
#include <freertos/task.h>

#define WIFI_SSID ""
#define WIFI_PASS ""

c_controller_t _state;

esp_err_t c_controller_init()
{
    ESP_ERROR_CHECK(c_wifi_init(WIFI_SSID, WIFI_PASS));

    ESP_ERROR_CHECK(c_mqtt_init());
    ESP_ERROR_CHECK(c_display_init());
    ESP_ERROR_CHECK(c_http_server_init());

    return ESP_OK;
}

esp_err_t c_controller_run(c_controller_t* state) { return ESP_OK; }

void app_main(void)
{
    ESP_ERROR_CHECK(c_controller_init());

    for (;;) {
        ESP_ERROR_CHECK(c_controller_run(&_state));

        vTaskDelay(1000 / portTICK_PERIOD_MS);
    }
}