#pragma once

#include "types.h"
#include <PubSubClient.h>

void init_mqtt(PubSubClient *mqtt_client);

void publish_mqtt_sensor_data(PubSubClient *mqtt_client, sensorData_t data);