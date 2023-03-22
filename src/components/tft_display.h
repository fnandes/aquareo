#pragma once

#include "core/display.h"
#include <U8g2lib.h>

namespace aquareo {

class TFTDisplay : public Display {
  private:
    U8G2 &u8g2;

  public:
    TFTDisplay(U8G2 &u8g2);
    void print(displayData_t data) override;
    void setup() override;
};

} // namespace aquareo
