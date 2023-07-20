#pragma once
#include <LiquidCrystal_I2C.h>
#include "types.h"

void init_display(LiquidCrystal_I2C *lcd);

void update_display_data(LiquidCrystal_I2C *lcd, sensorData_t data);