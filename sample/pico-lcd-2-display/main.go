package main

import (
	"machine"
	"time"

	"image/color"

	"tinygo.org/x/drivers/st7789"
)

func main() {
	machine.Serial.Configure(machine.UARTConfig{
		BaudRate: 115200,
	})
	time.Sleep(2 * time.Second)

	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 32_000_000, // 32MHz
		SCK:       machine.GP10,
		SDO:       machine.GP11,
		SDI:       machine.GP12, // ダミー - ST7789は受信しないが、SPIドライバが要求
		Mode:      0,
	})

	display := st7789.New(
		machine.SPI1,
		machine.GP12, // RST
		machine.GP8,  // DC
		machine.GP9,  // CS
		machine.GP13, // BL
	)

	display.Configure(st7789.Config{
		Width:        240,
		Height:       320,
		Rotation:     st7789.ROTATION_90, // 90度回転で横向き表示
		RowOffset:    0,
		ColumnOffset: 0,
	})

	white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 255, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	colors := []color.RGBA{white, red, green, blue}

	for {
		for _, c := range colors {
			display.FillScreen(c)
			time.Sleep(3 * time.Second)
		}
	}
}
