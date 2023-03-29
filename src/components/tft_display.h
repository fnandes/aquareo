#pragma once

#include "controller.h"
#include <U8g2lib.h>

namespace aquareo {

class TFTDisplay : public Display {
  public:
    TFTDisplay(U8G2& u8g2);
    void print(displayData_t data) override;
    void setup() override;

  private:
    U8G2& u8g2_;
};

} // namespace aquareo
