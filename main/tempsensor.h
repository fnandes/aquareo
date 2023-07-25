#pragma once

void tempsensor_init();

int tempsensor_device_count();

float tempsensor_read_temperature(int index);