#pragma once
#include <esp_err.h>

typedef enum {
    WF_NOT_INITIALIZED,
    WF_CONNECTING,
    WF_CONNECTED,
    WF_ERROR,
} c_wifi_stat_e;

typedef enum {
    MQ_NOT_INITIALIZED,
    MQ_CONNECTING,
    MQ_CONNECTED,
    MQ_ERROR,
} c_mqtt_stat_e;

typedef struct c_controller_t {
    c_wifi_stat_e wifi_stat;
    c_mqtt_stat_e mqtt_stat;
    float         temperature_current_value;
    float         ph_current_value;
    float         mqtt_last_publish;
} c_controller_t;

esp_err_t c_controller_init();
esp_err_t c_controller_run(c_controller_t* state);

esp_err_t     c_wifi_init(const char* ssid, const char* pass);
c_wifi_stat_e c_wifi_get_stat();

esp_err_t     c_mqtt_init();
c_mqtt_stat_e c_mqtt_get_stat();
esp_err_t     c_mqtt_publish_state(c_controller_t* state);

esp_err_t c_temperature_sensor_init();
float     c_temperature_sensor_read();

esp_err_t c_ph_sensor_init();
float     c_ph_sensor_read();

esp_err_t c_display_init();
esp_err_t c_display_update(c_controller_t* state);

esp_err_t c_http_server_init();