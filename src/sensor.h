#pragma once

namespace aquareo {

class Sensor {
  public:
    explicit Sensor() = default;

    virtual void  setup()                  = 0;
    virtual void  loop(unsigned long tick) = 0;
    virtual bool  get_state() const        = 0;
    virtual float get_measurement() const  = 0;

    virtual const char* get_device_class()        = 0;
    virtual const char* get_name()                = 0;
    virtual const char* get_unique_id()           = 0;
    virtual const char* get_unit_of_measurement() = 0;
};

} // namespace aquareo