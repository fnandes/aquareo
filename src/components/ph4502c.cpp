#include "ph4502c.h"

namespace aquareo {

const auto samples = 10;

PH4502PHSensor::PH4502PHSensor(Adafruit_ADS1115& ads) : ads_{ads} {}

const char* PH4502PHSensor::get_device_class() { return nullptr; }

const char* PH4502PHSensor::get_name() { return "ph"; }

const char* PH4502PHSensor::get_unique_id() { return "ph"; }

const char* PH4502PHSensor::get_unit_of_measurement() { return nullptr; }

void PH4502PHSensor::setup() { ads_.begin(); }

void PH4502PHSensor::loop(unsigned long tick)
{
    if (tick - last_update_ >= 3000) {
        last_update_     = tick;
        int measurements = 0;

        for (auto i = 0; i < samples; i++) {
            measurements += ads_.readADC_SingleEnded(0);
            delay(10);
        }
        float adc0    = measurements / samples;
        float voltage = ads_.computeVolts(adc0);
        Serial.print("Voltage: ");
        Serial.println(voltage);
        float ph = -5.64 * voltage + 21.55;
        // float ph = 7 + ((2.5 - voltage) / 0.1875);

        Serial.println(ph);

        current_val_ = ph;
    }
}

bool PH4502PHSensor::get_state() const { return true; }

float PH4502PHSensor::get_measurement() const { return current_val_; }

} // namespace aquareo
