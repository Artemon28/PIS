package main

import (
	"fmt"
)

type CarFactory interface {
	createCar() AbstractCar
	createEngine() AbstractEngine
}

func GetSportsFactory(brand string) (CarFactory, error) {

	if brand == "Ford" {
		return &Ford{}, nil
	}

	if brand == "Audi" {
		return &Audi{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
