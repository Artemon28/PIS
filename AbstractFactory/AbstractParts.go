package main

type AbstractCar interface {
	MaxSpeed(Engine) int
	ToString() string
	GetType() string
}

type Car struct {
	Name string
	Type string
}

func (car *Car) ToString() string {
	return "Auto " + car.Name
}

func (car *Car) GetType() string {
	return "Auto type is " + car.Type
}

func (car *Car) MaxSpeed(eng Engine) int {
	return eng.max_speed
}

type AbstractEngine interface {
	getMaxSpeed() int
}

type Engine struct {
	max_speed int
}

func (e Engine) getMaxSpeed() int {
	return e.max_speed
}
