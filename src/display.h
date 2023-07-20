#pragma once
#include "types.h"
#include <LiquidCrystal_I2C.h>

void init_display();

void update_display_data(sensorData_t data);