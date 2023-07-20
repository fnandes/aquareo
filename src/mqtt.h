#pragma once

#include "types.h"
#include <PubSubClient.h>
#include <WiFi.h>

void init_mqtt(WiFiClient* client);

void publish_mqtt_sensor_data(const sensorData_t data);