#include "controller.h"
#include <Adafruit_ADS1X15.h>

namespace aquareo {

class PH4502PHSensor : public Sensor {
  private:
    Adafruit_ADS1115& ads;
    float             currentVal{0};
    unsigned long     lastUpdate{0};

  public:
    PH4502PHSensor(Adafruit_ADS1115& ads);
    void    setup() override;
    void    loop(unsigned long tick) override;
    uint8_t getDeviceCount() const override;
    float   getCurrentValueByIndex(uint8_t idx) const override;
};

} // namespace aquareo
