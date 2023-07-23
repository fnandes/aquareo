#include "Display.h"
#include <Arduino.h>

void Display::setup() {}

void Display::loop(const unsigned long ticks) {}

void Display::setData(SensorReading_t* data)
{
    Serial.println("setData");
    char tempStr[10], phStr[10];

    snprintf(tempStr, sizeof(tempStr), "%f", data->currentTemp);
    snprintf(phStr, sizeof(phStr), "%f", data->currentPh);

    Serial.println(tempStr);
    Serial.println(phStr);
}