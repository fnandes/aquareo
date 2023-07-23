#pragma once
#include <DallasTemperature.h>

typedef struct {
    float currentTemp;
    float currentPh;
} SensorReading_t;

class SensorMonitor {
  public:
    SensorMonitor(DallasTemperature* ds18b20) {}

    void  setup();
    void  loop(unsigned long ticks);
    float readTemp();
    float readPh();

  private:
    DallasTemperature* m_tempSensors{nullptr};
    uint8_t            m_deviceAddress[8];

    bool          m_enabled;
    float         m_currTemp{0};
    float         m_currPh{0};
    unsigned long m_lastUpdate{0};
};