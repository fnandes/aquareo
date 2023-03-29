#pragma once

#include "configuration.h"
#include "sensor.h"
#include <PubSubClient.h>
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

class Controller {
  public:
    Controller(PubSubClient& mqtt, Display& display, Sensor* sensors[]);
    ~Controller();

    void setup();
    void reconnect_mqtt();
    void loop(unsigned long tick);

  private:
    PubSubClient& mqtt_;
    Display&      display_;
    Sensor*       sensors_[AQ_SENSOR_COUNT];
    unsigned long last_display_update_{0};
    unsigned long last_loop_{0};
    unsigned long last_publish_{0};

    void publish_measurements();
    void publish_discovery_data();
};

} // namespace aquareo