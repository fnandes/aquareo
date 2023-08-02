#include "config.h"
#include "display.h"
#include "sdkconfig.h"
#include <driver/gpio.h>
#include <driver/i2c.h>
#include <freertos/FreeRTOS.h>
#include <freertos/task.h>
#include <stdio.h>

static const char* TAG = "DP";

// LCD module defines
#define LCD_LINEONE 0x00 // start of line 1
#define LCD_LINETWO 0x40 // start of line 2

#define LCD_BACKLIGHT 0x08
#define LCD_ENABLE 0x04
#define LCD_COMMAND 0x00
#define LCD_WRITE 0x01

#define LCD_SET_DDRAM_ADDR 0x80
#define LCD_READ_BF 0x40

// LCD instructions
#define LCD_CLEAR 0x01             // replace all characters with ASCII 'space'
#define LCD_HOME 0x02              // return cursor to first position on first line
#define LCD_ENTRY_MODE 0x06        // shift cursor from left to right on read/write
#define LCD_DISPLAY_OFF 0x08       // turn display off
#define LCD_DISPLAY_ON 0x0C        // display on, cursor off, don't blink character
#define LCD_FUNCTION_RESET 0x30    // reset the LCD
#define LCD_FUNCTION_SET_4BIT 0x28 // 4-bit data, 2-line display, 5 x 7 font
#define LCD_SET_CURSOR 0x80        // set cursor position

void display_init()
{
    i2c_config_t i2c_cfg = {
        .mode             = I2C_MODE_MASTER,
        .sda_io_num       = AQ_GPIO_SDA_PIN,
        .scl_io_num       = AQ_GPIO_SLC_PIN,
        .sda_pullup_en    = GPIO_PULLUP_ENABLE,
        .scl_pullup_en    = GPIO_PULLUP_ENABLE,
        .master.clk_speed = 100000,
    };
    i2c_param_config(I2C_NUM_0, &i2c_cfg);
    i2c_driver_install(I2C_NUM_0, I2C_MODE_MASTER, 0, 0, 0);

    // initial delay
    vTaskDelay(100 / portTICK_PERIOD_MS);

    hdd44780_write_nibble(LCD_FUNCTION_RESET, LCD_COMMAND);  // First part of reset sequence
    vTaskDelay(10 / portTICK_PERIOD_MS);                     // 4.1 mS delay (min)
    hdd44780_write_nibble(LCD_FUNCTION_RESET, LCD_COMMAND);  // second part of reset sequence
    ets_delay_us(200);                                       // 100 uS delay (min)
    hdd44780_write_byte(LCD_FUNCTION_RESET, LCD_COMMAND);    // Third time's a charm
    hdd44780_write_byte(LCD_FUNCTION_SET_4BIT, LCD_COMMAND); // Activate 4-bit mode
    ets_delay_us(80);                                        // 40 uS delay (min)

    // --- Busy flag now available ---
    // Function Set instruction
    hdd44780_write_byte(LCD_FUNCTION_SET_4BIT, LCD_COMMAND); // Set mode, lines, and font
    ets_delay_us(80);

    // Clear Display instruction
    hdd44780_write_byte(LCD_CLEAR, LCD_COMMAND); // clear display RAM
    vTaskDelay(2 / portTICK_PERIOD_MS);          // Clearing memory takes a bit longer

    // Entry Mode Set instruction
    hdd44780_write_byte(LCD_ENTRY_MODE, LCD_COMMAND); // Set desired shift characteristics
    ets_delay_us(80);

    hdd44780_write_byte(LCD_DISPLAY_ON, LCD_COMMAND); // Ensure LCD is set to on
}

void display_write_state(const aquareo_state_t* state) {}

static void hdd44780_write_nibble(uint8_t nibble, uint8_t mode)
{
    uint8_t          data = (nibble & 0xF0) | mode | LCD_BACKLIGHT;
    i2c_cmd_handle_t cmd  = i2c_cmd_link_create();

    ESP_ERROR_CHECK(i2c_master_start(cmd));
    ESP_ERROR_CHECK(i2c_master_write_byte(cmd, (AQ_DISPLAY_I2C_ADDR << 1) | I2C_MASTER_WRITE, 1));
    ESP_ERROR_CHECK(i2c_master_write_byte(cmd, data, 1));
    ESP_ERROR_CHECK(i2c_master_stop(cmd));
    ESP_ERROR_CHECK(i2c_master_cmd_begin(I2C_NUM_0, cmd, 1000 / portTICK_PERIOD_MS));

    i2c_cmd_link_delete(cmd);

    hdd44780_pulse_enable(data);
}

static void hdd44780_pulse_enable(uint8_t data)
{
    i2c_cmd_handle_t cmd = i2c_cmd_link_create();
    ESP_ERROR_CHECK(i2c_master_start(cmd));
    ESP_ERROR_CHECK(i2c_master_write_byte(cmd, (AQ_DISPLAY_I2C_ADDR << 1) | I2C_MASTER_WRITE, 1));
    ESP_ERROR_CHECK(i2c_master_write_byte(cmd, data | LCD_ENABLE, 1));
    ESP_ERROR_CHECK(i2c_master_stop(cmd));
    ESP_ERROR_CHECK(i2c_master_cmd_begin(I2C_NUM_0, cmd, 1000 / portTICK_PERIOD_MS));
    i2c_cmd_link_delete(cmd);
    ets_delay_us(1);

    i2c_cmd_handle_t cmd2 = i2c_cmd_link_create();
    ESP_ERROR_CHECK(i2c_master_start(cmd));
    ESP_ERROR_CHECK(i2c_master_write_byte(cmd, (AQ_DISPLAY_I2C_ADDR << 1) | I2C_MASTER_WRITE, 1));
    ESP_ERROR_CHECK(i2c_master_write_byte(cmd, (data & ~LCD_ENABLE), 1));
    ESP_ERROR_CHECK(i2c_master_stop(cmd));
    ESP_ERROR_CHECK(i2c_master_cmd_begin(I2C_NUM_0, cmd2, 1000 / portTICK_PERIOD_MS));
    i2c_cmd_link_delete(cmd2);
    ets_delay_us(500);
}

static void hdd44780_write_byte(uint8_t data, uint8_t mode)
{
    hdd44780_write_nibble(data & 0xF0, mode);
    hdd44780_write_nibble((data << 4) & 0xF0, mode);
}

void hdd44780_set_cursor(uint8_t col, uint8_t row)
{
    if (row > LCD_rows - 1) {
        ESP_LOGE(tag, "Cannot write to row %d. Please select a row in the range (0, %d)", row, LCD_rows - 1);
        row = LCD_rows - 1;
    }
    uint8_t row_offsets[] = {LCD_LINEONE, LCD_LINETWO};
    hdd44780_write_byte(LCD_SET_DDRAM_ADDR | (col + row_offsets[row]), LCD_COMMAND);
}