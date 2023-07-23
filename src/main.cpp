#include "Controller.h"
#include "Display.h"
#include "SensorMonitor.h"
#include "config.h"
#include <Arduino.h>
#include <DallasTemperature.h>
#include <OneWire.h>
#include <WiFi.h>

WiFiClient wifi;

OneWire           bus(AQ_TEMP_SENSOR_GPIO_BUS);
DallasTemperature ds18b20(&bus);

SensorMonitor sensors(&ds18b20);
Display       display;
Controller    controller(&sensors, &display);

void setup()
{
    Serial.begin(115200);

    WiFi.begin(AQ_WIFI_SSID, AQ_WIFI_SSID);

    controller.setup();
}

void loop()
{
    const unsigned long ticks = millis();
    controller.loop(ticks);
    delay(1000);
}