package main

import (
	"machine"
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	// Включаем BLE-адаптер
	must("enable BLE stack", adapter.Enable())

	// Создаем BLE-устройство
	adv := adapter.DefaultAdvertisement()
	must("configure advertisement", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName:    "NiceNano",
		ServiceUUIDs: []bluetooth.UUID{bluetooth.ServiceUUIDNordicUART},
	}))

	// Запускаем рекламу BLE
	must("start advertising", adv.Start())

	// Включаем светодиод для индикации работы BLE
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.High()
		time.Sleep(time.Millisecond * 500)
		led.Low()
		time.Sleep(time.Millisecond * 500)
	}
}

// Функция для обработки ошибок
func must(action string, err error) {
	if err != nil {
		println("Failed to", action, ":", err.Error())
	}
}
