#pragma once

#include "configuration.h"
#include "sensor.h"
#include <DallasTemperature.h>

namespace aquareo {

class TemperatureSensor : public Sensor {
  public:
    TemperatureSensor(DallasTemperature& sensors, const char* unique_id, uint8_t index)
        : device_index_(index), sensors_{sensors}, unique_id_(unique_id)
    {
    }

    const char* get_device_class() override;
    const char* get_name() override;
    const char* get_unique_id() override;
    const char* get_unit_of_measurement() override;

    void  setup() override;
    void  loop(unsigned long tick) override;
    bool  get_state() const override;
    float get_measurement() const override;

  private:
    DallasTemperature& sensors_;
    float              current_val_ = {0.0f};
    uint8_t            device_index_{0};
    // TODO: find a way to easily get the unique_id using the device_index_
    const char*   unique_id_;
    unsigned long last_update_{0};
};

} // namespace aquareo