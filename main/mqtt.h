#pragma once
#include "config.h"
#include <stdbool.h>

void mqtt_init();

bool mqtt_is_connected();

void mqtt_publish_state(const aquareo_state_t* state);