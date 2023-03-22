#pragma once

#include "core/display.h"
#include "core/sensor.h"

namespace aquareo {

class Controller {
  private:
    Sensor &ph;
    Sensor &temperature;
    Display &display;
    unsigned long lastSensorUpdate{0};
    unsigned long lastDisplayUpdate{0};

  public:
    Controller(Display &display, Sensor &temperature, Sensor &ph);
    ~Controller();

    void setup();
    void update(unsigned long tick);
};

} // namespace aquareo