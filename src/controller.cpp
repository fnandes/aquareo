#include "controller.h"

namespace aquareo {

const unsigned int DISPLAY_UPDATE_INTERVAL = 3000;
const unsigned int MAIN_LOOP_INTERVAL = 1000;

Controller::Controller(MQTTClient &mqtt, Display &display, Sensor &temperature,
                       Sensor &ph)
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
    if (tick - lastLoop >= MAIN_LOOP_INTERVAL) {

        temperature.loop(tick);
        ph.loop(tick);
        mqtt.loop(tick);

        const float temp_val_1 = temperature.getCurrentValueByIndex(0);
        const float temp_val_2 = temperature.getCurrentValueByIndex(1);
        const float ph_val = ph.getCurrentValueByIndex(0);

        if (tick - lastDisplayUpdate >= DISPLAY_UPDATE_INTERVAL) {
            lastDisplayUpdate = tick;

            displayData_t data;
            data.temperature1 = temp_val_1;
            data.temperature2 = temp_val_2;
            data.ph = ph_val;

            display.print(data);
        }
    }
}

Controller::~Controller() {}

} // namespace aquareo
