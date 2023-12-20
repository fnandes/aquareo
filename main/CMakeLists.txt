set(SOURCES main.c
			wifi.c
			mqtt.c
			ds18b20.c
			hd44780.c
			ph.c
			http_server.c)

idf_component_register(SRCS ${SOURCES}
					INCLUDE_DIRS "../include"
					REQUIRES esp_wifi nvs_flash esp_event esp_netif lwip driver esp_http_server)