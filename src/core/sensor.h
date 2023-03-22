#pragma once

#include <stdint.h>

namespace aquareo {

class Sensor {
  public:
    Sensor() = default;
    virtual void setup() = 0;
    virtual void update() = 0;
    virtual uint8_t getDeviceCount() const = 0;
    virtual float getCurrentValueByIndex(uint8_t idx) const = 0;
};

} // namespace aquareo
