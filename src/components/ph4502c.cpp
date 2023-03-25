#include "ph4502c.h"

namespace aquareo {

const auto samples = 10;

PH4502PHSensor::PH4502PHSensor(Adafruit_ADS1115& ads) : ads{ads} {}

void PH4502PHSensor::setup() { ads.begin(); }

void PH4502PHSensor::loop(unsigned long tick)
{
    if (tick - lastUpdate >= 3000) {
        lastUpdate       = tick;
        int measurements = 0;

        for (auto i = 0; i < samples; i++) {
            measurements += ads.readADC_SingleEnded(0);
            delay(10);
        }
        float adc0    = measurements / samples;
        float voltage = ads.computeVolts(adc0);
        Serial.print("Voltage: ");
        Serial.println(voltage);
        float ph = -5.64 * voltage + 21.55;
        // float ph = 7 + ((2.5 - voltage) / 0.1875);

        Serial.println(ph);

        currentVal = ph;
    }
}

uint8_t PH4502PHSensor::getDeviceCount() const { return 1; }

float PH4502PHSensor::getCurrentValueByIndex(uint8_t idx) const { return currentVal; }

} // namespace aquareo
