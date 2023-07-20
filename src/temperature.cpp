#include "temperature.h"
#include <Arduino.h>
#include <DallasTemperature.h>
#include <OneWire.h>
#include <esp_log.h>

static const char* TAG = "TEMP";

OneWire           bus(AQ_TEMP_SENSOR_BUS);
DallasTemperature ds18b20(&bus);

static uint8_t device_addr[8];
static uint8_t device_addr_2[8];

void init_temperature_sensors()
{
    ESP_LOGI(TAG, "init");

    ds18b20.begin();

    if (ds18b20.getAddress(device_addr, 0)) {
        ESP_LOGI(TAG, "detected temperature sensor");
    }
}

float get_temperature()
{
    if (device_addr != nullptr) {
        ESP_LOGI(TAG, "reading temp");
        return ds18b20.getTempC(device_addr);
    }
    return 0.0f;
}