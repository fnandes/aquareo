#pragma once

#include <stdint.h>

namespace aquareo {

struct displayData_t {
    float temperature1{0.0f};
    float temperature2{0.0f};
    float ph{0.0f};
};

class Display {
  public:
    Display()                              = default;
    virtual void setup()                   = 0;
    virtual void print(displayData_t data) = 0;
};

class Sensor {
  public:
    Sensor()                                                  = default;
    virtual void    setup()                                   = 0;
    virtual void    loop(unsigned long tick)                  = 0;
    virtual uint8_t getDeviceCount() const                    = 0;
    virtual float   getCurrentValueByIndex(uint8_t idx) const = 0;
};

class MQTTClient {
  public:
    MQTTClient()                                     = default;
    virtual void setup()                             = 0;
    virtual void loop(unsigned long tick)            = 0;
    virtual void sendSensorData(char* ns, float val) = 0;
};

class Controller {
  private:
    MQTTClient&   mqtt;
    Display&      display;
    Sensor&       temperature;
    Sensor&       ph;
    unsigned long lastDisplayUpdate{0};
    unsigned long lastLoop{0};
    unsigned long lastPublish{0};

  public:
    Controller(MQTTClient& mqtt, Display& display, Sensor& temperature, Sensor& ph);
    ~Controller();

    void setup();
    void loop(unsigned long tick);
};

} // namespace aquareo