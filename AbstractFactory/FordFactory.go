package main

type Ford struct {
}

type FordCar struct {
	Car
}

type FordEngine struct {
	Engine
}

func (fe FordEngine) getMaxSpeed() int {
	return fe.max_speed
}

func (a *Ford) createCar() AbstractCar {
	return &FordCar{
		Car: Car{
			Name: "Ford",
			Type: "Sedan",
		},
	}
}

func (fc *FordCar) ToString() string {
	return fc.Car.ToString()
}

func (fc *FordCar) GetType() string {
	return fc.Car.GetType()
}

func (fc *FordCar) MaxSpeed(eng Engine) int {
	return fc.Car.MaxSpeed(eng)
}

func (a *Ford) createEngine() AbstractEngine {
	return &FordEngine{
		Engine: Engine{
			max_speed: 220,
		},
	}
}
