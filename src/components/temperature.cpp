#include "components/temperature.h"
#include <string>

namespace aquareo {

const char* TemperatureSensor::get_device_class() { return "temperature"; }

const char* TemperatureSensor::get_name() { return "temperature"; }

const char* TemperatureSensor::get_unique_id() { return unique_id_; }

const char* TemperatureSensor::get_unit_of_measurement() { return "°C"; }

void TemperatureSensor::setup() {}

void TemperatureSensor::loop(unsigned long tick)
{
    if (tick - last_update_ >= AQ_TP_SENSOR_TIME) {
        sensors_.requestTemperatures();
        current_val_ = sensors_.getTempCByIndex(device_index_);
    }
}

float TemperatureSensor::get_measurement() const { return current_val_; }

bool TemperatureSensor::get_state() const { return sensors_.isConnected(&device_index_); }

} // namespace aquareo
