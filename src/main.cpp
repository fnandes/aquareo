#include "components/temperature.h"
#include "components/tft_display.h"
#include "configuration.h"
#include "core/controller.h"
#include <Arduino.h>
#include <DallasTemperature.h>
#include <OneWire.h>
#include <SPI.h>
#include <U8g2lib.h>
#include <Wire.h>

using namespace aquareo;

OneWire bus(4);
DallasTemperature ds18b20(&bus);
TemperatureSensor tempSensor(ds18b20);
TemperatureSensor phSensor(ds18b20); // TODO: use PH sensor

U8G2_SH1106_128X64_NONAME_1_HW_I2C sh1106(U8G2_R0, /* reset=*/U8X8_PIN_NONE);
TFTDisplay display(sh1106);

aquareo::Controller controller(display, tempSensor, phSensor);

void setup()
{
    Serial.begin(115200);

    controller.setup();
}

void loop()
{
    unsigned long tick = millis();
    controller.update(tick);
}