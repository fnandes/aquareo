#pragma once
#include "configuration.h"
#include <DallasTemperature.h>

void init_temperature_sensors();

float get_temperature();