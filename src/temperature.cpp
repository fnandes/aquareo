#include "temperature.h"
#include "Arduino.h"

static uint8_t device_addr[8];
static uint8_t device_addr_2[8];

void init_temperature_sensors(DallasTemperature *sensors)
{
    if (sensors->getAddress(device_addr, 0))
    {
        Serial.println("detected temperature sensor");
    }
}

float get_temperature(DallasTemperature *sensors)
{
    return sensors->getTempC(device_addr);
}