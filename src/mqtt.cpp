#include "types.h"
#include "mqtt.h"
#include <PubSubClient.h>
#include <Arduino.h>

void init_mqtt(PubSubClient *client)
{
    client->connect();
}

void publish_mqtt_sensor_data(PubSubClient *client, sensorData_t data)
{
    if (!client->connected())
    {
        Serial.println("MQTT Broker not connected! trying to connect again ...");
        client->connect();
        return;
    }

    client->publish("aquareo/sensor/temperature", data.temperature);
    client->publish("aquareo/sensor/ph", data.ph)
}