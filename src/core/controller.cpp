#include "core/controller.h"
#include "core/sensor.h"

namespace aquareo {

const unsigned long DISPLAY_UPDATE_INTERVAL = 3000;
const unsigned long SENSOR_UPDATE_INTERVAL = 5000;

Controller::Controller(Display &display, Sensor &temperature, Sensor &ph)
    : display{display}, temperature{temperature}, ph{ph}
{
}

void Controller::setup()
{
    temperature.setup();
    ph.setup();
    display.setup();
}

void Controller::update(unsigned long tick)
{
    if (tick - lastSensorUpdate >= SENSOR_UPDATE_INTERVAL) {
        lastSensorUpdate = tick;

        temperature.update();
        ph.update();
    }

    float temp_val_1 = temperature.getCurrentValueByIndex(0);
    float temp_val_2 = temperature.getCurrentValueByIndex(1);
    float ph_val = ph.getCurrentValueByIndex(0);

    if (tick - lastDisplayUpdate >= DISPLAY_UPDATE_INTERVAL) {
        lastDisplayUpdate = tick;

        displayData_t data;
        data.temperature1 = temp_val_1;
        data.temperature2 = temp_val_2;
        data.ph = ph_val;

        display.print(data);
    }
}

Controller::~Controller() {}

} // namespace aquareo
