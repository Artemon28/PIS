// package main

// type Audi struct {
// }

// type AudiCar struct {
// 	Car
// }

// type AudiEngine struct {
// 	Engine
// }

// func (ae AudiEngine) getMaxSpeed() int {
// 	return ae.max_speed
// }

// func (a *Audi) createCar() AbstractCar {
// 	return &AudiCar{
// 		Car: Car{
// 			Name: "Audi",
// 			Type: "Crossover",
// 		},
// 	}
// }

// func (fc *AudiCar) ToString() string {
// 	return fc.Car.ToString()
// }

// func (fc *AudiCar) GetType() string {
// 	return fc.Car.GetType()
// }

// func (fc *AudiCar) MaxSpeed(eng Engine) int {
// 	return fc.Car.MaxSpeed(eng)
// }

// func (a *Audi) createEngine() AbstractEngine {
// 	return &AudiEngine{
// 		Engine: Engine{
// 			max_speed: 300,
// 		},
// 	}
// }

//Singleton version

package main

import "sync"

type AudiCar struct {
	Car
}

type AudiEngine struct {
	Engine
}

// AudiFactory is a singleton factory for creating Audi cars and engines.
type Audi struct {
	once sync.Once
}

var audiFactory *Audi

// GetAudiFactoryInstance returns the singleton instance of the AudiFactory.
func GetAudiFactoryInstance() *Audi {
	if audiFactory == nil {
		audiFactory = &Audi{}
	}
	return audiFactory
}

func (ae AudiEngine) getMaxSpeed() int {
	return ae.max_speed
}

func (a *Audi) createCar() AbstractCar {
	return &AudiCar{
		Car: Car{
			Name: "Audi",
			Type: "Crossover",
		},
	}
}

func (fc *AudiCar) ToString() string {
	return fc.Car.ToString()
}

func (fc *AudiCar) GetType() string {
	return fc.Car.GetType()
}

func (fc *AudiCar) MaxSpeed(eng Engine) int {
	return fc.Car.MaxSpeed(eng)
}

func (a *Audi) createEngine() AbstractEngine {
	return &AudiEngine{
		Engine: Engine{
			max_speed: 300,
		},
	}
}
