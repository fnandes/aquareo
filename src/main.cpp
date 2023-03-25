#include "components/ph4502c.h"
#include "components/temperature.h"
#include "components/tft_display.h"
#include "components/wifi_mqtt_client.h"
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
TemperatureSensor tempSensor(ds18b20);

U8G2_SH1106_128X64_NONAME_1_HW_I2C sh1106(U8G2_R0, /* reset=*/U8X8_PIN_NONE);
TFTDisplay                         display(sh1106);

WiFiClient     wifi;
PubSubClient   pubSub(wifi);
WiFiMQTTClient mqtt(pubSub);

// PH
Adafruit_ADS1115 ads;
PH4502PHSensor   phSensor(ads);

aquareo::Controller controller(mqtt, display, tempSensor, phSensor);

void setup()
{
    Serial.begin(115200);

    Wire.begin();
    WiFi.begin(ssid, password);

    controller.setup();
}

void loop()
{
    unsigned long tick = millis();
    controller.loop(tick);
}