#include "Display.h"
#include "SensorMonitor.h"

class Controller {
  public:
    Controller(SensorMonitor* sensors, Display* display) : m_sensors{sensors}, m_display{display} {}
    void setup();
    void loop(const unsigned long ticks);

  private:
    unsigned long  m_lastUpdate{0};
    SensorMonitor* m_sensors{nullptr};
    Display*       m_display{nullptr};
};