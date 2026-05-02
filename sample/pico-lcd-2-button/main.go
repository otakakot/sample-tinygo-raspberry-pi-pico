package main

import (
	"machine"
	"time"
)

func main() {
	// Pico LCD 2のボタンピン設定
	// KEY0: GP15, KEY1: GP17, KEY2: GP2, KEY3: GP3
	key0 := machine.GP15
	key1 := machine.GP17
	key2 := machine.GP2
	key3 := machine.GP3

	key0.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	key1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	key2.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	key3.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	prevKey0 := true
	prevKey1 := true
	prevKey2 := true
	prevKey3 := true

	println("Pico LCD 2 ボタンテスト開始")
	println("ボタンを押してください...")

	for {
		currentKey0 := key0.Get()
		currentKey1 := key1.Get()
		currentKey2 := key2.Get()
		currentKey3 := key3.Get()

		if !currentKey0 && prevKey0 {
			println("ボタン 0 が押されました")
		}

		if !currentKey1 && prevKey1 {
			println("ボタン 1 が押されました")
		}

		if !currentKey2 && prevKey2 {
			println("ボタン 2 が押されました")
		}

		if !currentKey3 && prevKey3 {
			println("ボタン 3 が押されました")
		}

		prevKey0 = currentKey0
		prevKey1 = currentKey1
		prevKey2 = currentKey2
		prevKey3 = currentKey3

		time.Sleep(time.Millisecond * 50)
	}
}
