#include "components/temperature.h"

namespace aquareo {

const int SENSOR_UPDATE_INTERVAL = 5000;

void TemperatureSensor::setup()
{
    sensors.begin();
    deviceCount = sensors.getDeviceCount();
}

void TemperatureSensor::loop(unsigned long tick)
{
    if (tick - lastUpdate >= SENSOR_UPDATE_INTERVAL) {
        sensors.requestTemperatures();

        for (size_t i = 0; i < 2; i++) {
            float temperature = sensors.getTempCByIndex(i);
            currentTemperatures[i] = temperature;
        }
    }
}

float TemperatureSensor::getCurrentValueByIndex(uint8_t idx) const
{
    return currentTemperatures[idx];
}

uint8_t TemperatureSensor::getDeviceCount() const { return deviceCount; }

} // namespace aquareo
