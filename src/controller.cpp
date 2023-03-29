#include "controller.h"
#include "configuration.h"
#include <ArduinoJson.h>
#include <PubSubClient.h>
#include <StreamUtils.h>

namespace aquareo {

Controller::Controller(PubSubClient& mqtt, Display& display, Sensor* sensors[]) : mqtt_{mqtt}, display_{display}
{
    for (auto i = 0; i < AQ_SENSOR_COUNT; i++) {
        this->sensors_[i] = sensors[i];
    }
}

Controller::~Controller() {}

void Controller::setup()
{
    for (auto i = 0; i < AQ_SENSOR_COUNT; i++) {
        sensors_[i]->setup();
    }
    display_.setup();

    this->reconnect_mqtt();
}

void Controller::reconnect_mqtt()
{
    mqtt_.setServer(AQ_MQTT_CONN_HOST, AQ_MQTT_CONN_PORT);
    mqtt_.connect(AQ_MQTT_CONN_ID, "aquareo", "1234");

    this->publish_discovery_data();
}

void Controller::loop(unsigned long tick)
{
    if (tick - last_loop_ >= AQ_MAIN_LOOP_TIME) {
        last_loop_ = tick;

        if (!mqtt_.connected()) {
            Serial.println("MQTT not connected, trying to reconnect...");
            this->reconnect_mqtt();
        }

        for (auto i = 0; i < AQ_SENSOR_COUNT; i++) {
            sensors_[i]->loop(tick);
        }

        if (tick - last_display_update_ >= AQ_DISPLAY_UPDATE_TIME) {
            last_display_update_ = tick;

            displayData_t data;
            data.temperature1 = sensors_[0]->get_measurement();
            data.temperature2 = sensors_[1]->get_measurement();
            data.ph           = sensors_[2]->get_measurement();

            display_.print(data);
        }

        if (tick - last_publish_ >= AQ_MQTT_PUBLISH_TIME) {
            last_publish_ = tick;

            this->publish_measurements();
        }
    }
}

const char* stateTopic = "homeassistant/sensor/aquareo/state";

void Controller::publish_measurements()
{
    DynamicJsonDocument doc(1024);
    char                buff[256];

    for (auto i = 0; i < AQ_SENSOR_COUNT; i++) {
        doc[this->sensors_[i]->get_unique_id()] = this->sensors_[i]->get_measurement();
    }

    size_t n = serializeJson(doc, buff);

    if (!mqtt_.publish(stateTopic, buff, n)) {
        Serial.println("Unable to publish state");
    }
}

void Controller::publish_discovery_data()
{
    DynamicJsonDocument payload(1024);
    char                buff[256];

    for (auto i = 0; i < AQ_SENSOR_COUNT; i++) {
        Sensor*      sensor         = this->sensors_[i];
        const auto   uid            = String(sensor->get_unique_id());
        const String discoveryTopic = "homeassistant/sensor/aquareo/" + uid + "/config";

        Serial.print("Sending discovery data: ");
        Serial.println(discoveryTopic);

        payload["uniq_id"]  = sensor->get_unique_id();
        payload["stat_cla"] = "total";
        payload["stat_t"]   = stateTopic;
        payload["frc_upd"]  = true;
        if (sensor->get_unit_of_measurement() != nullptr) {
            payload["unit_of_meas"] = sensor->get_unit_of_measurement();
        }
        if (sensor->get_device_class() != nullptr) {
            payload["dev_cla"] = sensor->get_device_class();
        }
        payload["val_tpl"] = "{{ value_json." + uid + "|default(0) }}";

        this->mqtt_.beginPublish(discoveryTopic.c_str(), measureJson(payload), false);
        {
            BufferingPrint buffered_msg(this->mqtt_, 64);
            serializeJson(payload, buffered_msg);
            buffered_msg.flush();
        }
        this->mqtt_.endPublish();

        payload.clear();
    }
}

} // namespace aquareo

// amarelo 3
// preto 4
// verm 2