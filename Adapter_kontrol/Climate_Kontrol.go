package main

import (
	"fmt"
	"math/rand"
)

type ThirdPartyInterface interface {
	GetTemperatureInFahrenheit() float64
}

type ThirdPartyTemperatureSensor struct {
	readingInFahrenheit float64
}

func NewThirdPartyTemperatureSensor() *ThirdPartyTemperatureSensor {
	return &ThirdPartyTemperatureSensor{}
}

func (sensor *ThirdPartyTemperatureSensor) GetTemperatureInFahrenheit() float64 {
	return sensor.readingInFahrenheit
}

// interface that used to to
type CommonInterface interface {
	GetTemperatureInCelsius() float64
}

type CelsiusTemperatureAdapter struct {
	thirdPartySensor *ThirdPartyTemperatureSensor
}

func NewCelsiusTemperatureAdapter(sensor *ThirdPartyTemperatureSensor) *CelsiusTemperatureAdapter {
	return &CelsiusTemperatureAdapter{thirdPartySensor: sensor}
}

func (adapter *CelsiusTemperatureAdapter) GetTemperatureInCelsius() float64 {
	fahrenheitTemp := adapter.thirdPartySensor.GetTemperatureInFahrenheit()
	celsiusTemp := (fahrenheitTemp - 32) * 5 / 9
	return celsiusTemp
}

type SystemClimateControl struct {
}

func (s *SystemClimateControl) HeatUp() {
	fmt.Println("Turn on heating")
}

func (s *SystemClimateControl) HeatOff() {
	fmt.Println("Turn off heating")
}

func (c *SystemClimateControl) ChangeClimate(com CommonInterface) {
	if com.GetTemperatureInCelsius() < 20.0 {
		c.HeatUp()
	} else {
		c.HeatOff()
	}
}

func main() {
	thirdPartySensor := NewThirdPartyTemperatureSensor()
	thirdPartySensor.readingInFahrenheit = 60.0 + rand.Float64()*20.0
	celsiusAdapter := NewCelsiusTemperatureAdapter(thirdPartySensor)
	celsiusTemp := celsiusAdapter.GetTemperatureInCelsius()
	fmt.Printf("Temperature in Celsius: %.2f°C\n", celsiusTemp)

	system := SystemClimateControl{}
	system.ChangeClimate(celsiusAdapter)
}
