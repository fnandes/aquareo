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
    void sendSensorData(const char* ns, float val) override;
};

} // namespace aquareo
