#include "sensor.h"
#include <Adafruit_ADS1X15.h>

namespace aquareo {

class PH4502PHSensor : public Sensor {
  public:
    PH4502PHSensor(Adafruit_ADS1115& ads);

    const char* get_device_class() override;
    const char* get_name() override;
    const char* get_unique_id() override;
    const char* get_unit_of_measurement() override;

    void  setup() override;
    void  loop(unsigned long tick) override;
    bool  get_state() const override;
    float get_measurement() const override;

  private:
    Adafruit_ADS1115& ads_;
    float             current_val_{0};
    unsigned long     last_update_{0};
};

} // namespace aquareo
