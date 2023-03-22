#include "components/wifi_mqtt_client.h"
#include "configuration.h"
#include <PubSubClient.h>

namespace aquareo {

WiFiMQTTClient::WiFiMQTTClient(PubSubClient &client) : client{client} {}

void WiFiMQTTClient::loop(unsigned long tick)
{
    isBrokerConnected = client.connected();

    if (!isBrokerConnected) {
        client.connect(AQ_MQTT_CONN_ID);
    }
}

void WiFiMQTTClient::setup()
{
    client.setServer(AQ_MQTT_CONN_HOST, AQ_MQTT_CONN_PORT);
    client.connect(AQ_MQTT_CONN_ID);
}

void WiFiMQTTClient::sendSensorData(char *ns, float val)
{
    if (isBrokerConnected) {
        char tempStr[8];
        dtostrf(val, 1, 2, tempStr);
        client.publish(ns, tempStr);
    }
}

} // namespace aquareo