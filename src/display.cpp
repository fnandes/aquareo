#include "display.h"
#include "configuration.h"
#include "types.h"
#include <LiquidCrystal_I2C.h>
#include <esp_log.h>

static const char* TAG = "DSPL";

static LiquidCrystal_I2C lcd(0x27, 16, 2);

void init_display()
{
    ESP_LOGI(TAG, "init");

    lcd.init();
    lcd.backlight();
}

void update_display_data(sensorData_t data)
{
    if (lcd.availableForWrite()) {
        lcd.setCursor(0, 0);
        lcd.print("T: ");
        lcd.setCursor(3, 0);
        lcd.print(data.temperature);
        lcd.setCursor(0, 1);
        lcd.print("PH: ");
        lcd.setCursor(4, 1);
        lcd.print(data.ph);
    }
}