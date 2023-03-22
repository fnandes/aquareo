#include "components/wifi_mqtt_client.h"
#include <PubSubClient.h>

namespace aquareo {

WiFiMQTTClient::WiFiMQTTClient(PubSubClient &client) : client{client} {}

void WiFiMQTTClient::loop(unsigned long tick) {}

void WiFiMQTTClient::setup() {}

void WiFiMQTTClient::sendSensorData(char *ns, float val) {}

} // namespace aquareo