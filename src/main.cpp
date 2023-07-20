#include <Arduino.h>
#include <OneWire.h>
#include <DallasTemperature.h>
#include <WiFi.h>
#include <LiquidCrystal_I2C.h>
#include "temperature.h"
#include "configuration.h"
#include "types.h"
#include "display.h"
#include "WiFi.h"
#include "PubSubClient.h"

OneWire bus(AQ_TEMP_SENSOR_BUS);
DallasTemperature ds18b20(&bus);

// lcd
LiquidCrystal_I2C lcd(0x27, 16, 2);

// network
WiFiClient wifi;
PubSubClient pubSub(wifi);

void setup()
{
    Serial.begin(115200);
    WiFi.begin(WIFI_SSID, WIFI_PASS);
    ds18b20.begin();

    init_temperature_sensors(&ds18b20);
    init_display(&lcd);
}

void loop()
{
    const sensorData_t data{
        .temperature = get_temperature(&ds18b20),
        .ph = 7, // TODO: make this work
    };

    update_display_data(&lcd, data);

    delay(2000);
}