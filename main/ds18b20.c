#include "tempsensor.h"
#include "config.h"
#include <ds18b20.h>
#include <esp_err.h>
#include <esp_log.h>
#include <onewire_bus.h>

static const char* TAG = "TP";

static int                     ds18b20_device_count = 0;
static ds18b20_device_handle_t ds18b20s_devices[AQ_ONEWIRE_MAX_DS18B20];

void tempsensor_init()
{
    ESP_LOGI(TAG, "init");

    onewire_bus_handle_t bus        = NULL;
    onewire_bus_config_t bus_config = {
        .bus_gpio_num = AQ_ONEWIRE_BUS_GPIO,
    };
    onewire_bus_rmt_config_t rmt_config = {
        .max_rx_bytes = 10, // 1byte ROM command + 8byte ROM number + 1byte device command
    };
    ESP_ERROR_CHECK(onewire_new_bus_rmt(&bus_config, &rmt_config, &bus));

    onewire_device_iter_handle_t iter = NULL;
    ESP_ERROR_CHECK(onewire_new_device_iter(bus, &iter));

    onewire_device_t next_onewire_device;
    ds18b20_config_t ds_cfg        = {};
    esp_err_t        search_result = ESP_OK;
    int              attempts      = 0;

    do {
        search_result = onewire_device_iter_get_next(iter, &next_onewire_device);
        if (search_result == ESP_OK) {
            if (ds18b20_new_device(&next_onewire_device, &ds_cfg, &ds18b20s_devices[ds18b20_device_count]) == ESP_OK) {
                ds18b20_device_count++;
            }
        }
    } while (search_result != ESP_ERR_NOT_FOUND && ++attempts <= AQ_ONEWIRE_MAX_SEARCH_ATTEMPTS);
    ESP_LOGI(TAG, "%d devices found", ds18b20_device_count);
    ESP_ERROR_CHECK(onewire_del_device_iter(iter));
}

float tempsensor_read_temperature(int index)
{
    float temperature = -1;

    if (index > ds18b20_device_count)
        return temperature;

    ESP_LOGI(TAG, "reading temp");

    ESP_ERROR_CHECK(ds18b20_trigger_temperature_conversion(ds18b20s_devices[index]));
    ESP_ERROR_CHECK(ds18b20_get_temperature(ds18b20s_devices[index], &temperature));
    return temperature;
}

int tempsensor_device_count() { return ds18b20_device_count; }