#include "SensorMonitor.h"
#include "config.h"
#include <Arduino.h>

void SensorMonitor::setup()
{
    if (m_tempSensors == nullptr)
        return;

    m_tempSensors->begin();
    if (m_tempSensors->getAddress(m_deviceAddress, 0)) {
        m_enabled = true;
    }
}

void SensorMonitor::loop(unsigned long ticks)
{
    if (ticks - m_lastUpdate < AQ_SENSOR_UPDATE_INTERVAL)
        return;

    m_currTemp   = m_tempSensors->getTempCByIndex(0);
    m_lastUpdate = ticks;
}

float SensorMonitor::readTemp() { return m_currTemp; }

float SensorMonitor::readPh() { return m_currPh; }