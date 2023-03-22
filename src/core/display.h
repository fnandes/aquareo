#pragma once

namespace aquareo {

struct displayData_t {
    float temperature1{0.0f};
    float temperature2{0.0f};
    float ph{0.0f};
};

class Display {
  public:
    Display() = default;
    virtual void setup() = 0;
    virtual void print(displayData_t data) = 0;
};

} // namespace aquareo
