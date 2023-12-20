#include "aquareo.h"

esp_err_t c_mqtt_init() { return ESP_OK; }

inline c_mqtt_stat_e c_mqtt_get_stat() { return MQ_CONNECTING; }

esp_err_t c_mqtt_publish_state(c_controller_t* state) { return ESP_OK; }