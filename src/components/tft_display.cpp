#include "components/tft_display.h"

namespace aquareo {

TFTDisplay::TFTDisplay(U8G2& u8g2) : u8g2_{u8g2} {}

void TFTDisplay::print(displayData_t data)
{
    u8g2_.setFont(u8g2_font_lubB08_tf);
    u8g2_.firstPage();

    do {
        u8g2_.setCursor(0, 20);
        u8g2_.print("T1:  ");
        u8g2_.print(data.temperature1, 1);
        u8g2_.setCursor(0, 30);
        u8g2_.print("T2:  ");
        u8g2_.print(data.temperature2, 1);
        u8g2_.setCursor(0, 40);
        u8g2_.print("PH:  ");
        u8g2_.print(data.ph, 1);
    } while (u8g2_.nextPage());
}

void TFTDisplay::setup() { u8g2_.begin(); }

} // namespace aquareo
