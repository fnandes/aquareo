#pragma once
#include "SensorMonitor.h"

class Display {
  public:
    Display() {}
    void setup();
    void loop(const unsigned long ticks);
    void setData(SensorReading_t* data);
};