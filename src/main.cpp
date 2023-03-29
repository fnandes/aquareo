#include "components/ph4502c.h"
#include "components/temperature.h"
#include "components/tft_display.h"
#include "configuration.h"
#include "controller.h"
#include <Adafruit_ADS1X15.h>
#include <DallasTemperature.h>
#include <OneWire.h>
#include <PubSubClient.h>
#include <SPI.h>
#include <U8g2lib.h>
#include <WiFi.h>
#include <Wire.h>
#include <esp32-hal.h>

const char* ssid     = AQ_WIFI_SSID;
const char* password = AQ_WIFI_PWD;

using namespace aquareo;

OneWire           bus(AQ_TP_SENSOR_BUS_PIN);
DallasTemperature ds18b20(&bus);
TemperatureSensor tempSensor0(ds18b20, "temperature_0", 0);
TemperatureSensor tempSensor1(ds18b20, "temperature_1", 1);

U8G2_SH1106_128X64_NONAME_1_HW_I2C sh1106(U8G2_R0, /* reset=*/U8X8_PIN_NONE);
TFTDisplay                         display(sh1106);

WiFiClient   wifi;
PubSubClient pubSub(wifi);

// PH
Adafruit_ADS1115 ads;
PH4502PHSensor   phSensor(ads);

Sensor* sensors[]{&tempSensor0, &tempSensor1, &phSensor};

aquareo::Controller controller(pubSub, display, sensors);

void setup()
{
    Serial.begin(115200);

    ds18b20.begin();
    Wire.begin();
    WiFi.begin(ssid, password);

    controller.setup();
}

void loop()
{
    unsigned long tick = millis();
    controller.loop(tick);
}