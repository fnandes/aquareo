#include "types.h"
#include "display.h"
#include "configuration.h"

void init_display(LiquidCrystal_I2C *lcd)
{
    lcd->init();
    lcd->backlight();
}

void update_display_data(LiquidCrystal_I2C *lcd, sensorData_t data)
{
    lcd->setCursor(0, 0);
    lcd->print("T: ");
    lcd->setCursor(3, 0);
    lcd->print(data.temperature);
    lcd->setCursor(0, 1);
    lcd->print("PH: ");
    lcd->setCursor(4, 1);
    lcd->print(data.ph);
}