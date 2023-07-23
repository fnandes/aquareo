#include "Controller.h"
#include "config.h"
#include <Wifi.h>

void Controller::setup()
{
    m_sensors->setup();

    m_sensors->setup();
    m_display->setup();
}

void Controller::loop(const unsigned long ticks)
{
    if (ticks - m_lastUpdate < AQ_MAIN_LOOP_INTERVAL)
        return;

    m_sensors->loop(ticks);

    SensorReading_t status = {
        .currentTemp = m_sensors->readTemp(),
        .currentPh   = m_sensors->readPh(),
    };

    m_display->setData(&status);

    m_lastUpdate = ticks;
}