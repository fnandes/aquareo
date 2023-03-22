#include "configuration.h"
#include "core/sensor.h"
#include <DallasTemperature.h>

namespace aquareo {

class TemperatureSensor : public Sensor {

  private:
    DallasTemperature &sensors;
    uint8_t deviceCount{0};
    std::array<float, 2> currentTemperatures = {0.0f, 0.0f};

  public:
    TemperatureSensor(DallasTemperature &sensors) : sensors{sensors} {}

    void setup() override;
    void update() override;
    uint8_t getDeviceCount() const override;
    float getCurrentValueByIndex(uint8_t idx) const override;
};

} // namespace aquareo