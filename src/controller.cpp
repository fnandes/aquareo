#include "controller.h"
#include "configuration.h"

namespace aquareo {

Controller::Controller(MQTTClient& mqtt, Display& display, Sensor& temperature, Sensor& ph)
    : mqtt{mqtt}, display{display}, temperature{temperature}, ph{ph}
{
}

void Controller::setup()
{
    temperature.setup();
    ph.setup();
    display.setup();
    mqtt.setup();
}

void Controller::loop(unsigned long tick)
{
    if (tick - lastLoop >= AQ_MAIN_LOOP_TIME) {
        lastLoop = tick;

        temperature.loop(tick);
        ph.loop(tick);
        mqtt.loop(tick);

        const float temp_val_1 = temperature.getCurrentValueByIndex(0);
        const float temp_val_2 = temperature.getCurrentValueByIndex(1);
        const float ph_val     = ph.getCurrentValueByIndex(0);

        if (tick - lastDisplayUpdate >= AQ_DISPLAY_UPDATE_TIME) {
            lastDisplayUpdate = tick;

            displayData_t data;
            data.temperature1 = temp_val_1;
            data.temperature2 = temp_val_2;
            data.ph           = ph_val;

            display.print(data);
        }

        if (tick - lastPublish >= AQ_MQTT_PUBLISH_TIME) {
            lastPublish = tick;

            mqtt.sendSensorData("aquareo/sensor/temperature_1", temp_val_1);
            mqtt.sendSensorData("aquareo/sensor/temperature_2", temp_val_2);
            mqtt.sendSensorData("aquareo/sensor/ph", ph_val);
        }
    }
}

Controller::~Controller() {}

} // namespace aquareo
