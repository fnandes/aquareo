#include "config.h"
#include "mqtt.h"
#include "tempsensor.h"
#include "wifi.h"
#include <esp_event.h>
#include <esp_log.h>
#include <esp_timer.h>
#include <freertos/FreeRTOS.h>
#include <freertos/task.h>
#include <nvs_flash.h>

static const char* TAG = "AQ";

static aquareo_state_t state = {
    .current_temp = {0, 0},
};

static int64_t mqtt_last_publish = 0;

static void app_update_sensor_reading()
{
    for (int i = 0; i < tempsensor_device_count(); i++) {
        state.current_temp[i] = tempsensor_read_temperature(i);
    }
}

void app_main(void)
{
    ESP_ERROR_CHECK(esp_event_loop_create_default());
    ESP_ERROR_CHECK(nvs_flash_init());
    ESP_ERROR_CHECK(esp_timer_early_init());

    tempsensor_init();
    wifi_init();
    mqtt_init();

    for (;;) {
        int64_t ticks = esp_timer_get_time() / 1000;

        app_update_sensor_reading();
        ESP_LOGI(TAG, "temp_0: %f", state.current_temp[0]);
        ESP_LOGI(TAG, "temp_1: %f", state.current_temp[1]);

        if (ticks - mqtt_last_publish >= 30000) {
            // publish every 30s
            mqtt_last_publish = ticks;
            mqtt_publish_state(&state);
        }

        vTaskDelay(AQ_SYS_CLOCK / portTICK_PERIOD_MS);
    }
}
