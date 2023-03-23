#pragma once

#include <PubSubClient.h>
#include <WiFi.h>
#include <controller.h>

namespace aquareo {

class WiFiMQTTClient : public MQTTClient {
  private:
    bool          isBrokerConnected{false};
    PubSubClient& client;

  public:
    WiFiMQTTClient(PubSubClient& client);

    void setup() override;
    void loop(unsigned long tick) override;
    void sendSensorData(char* ns, float val) override;
};

} // namespace aquareo
