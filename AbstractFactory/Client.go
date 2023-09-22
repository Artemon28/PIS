package main

import "fmt"

type Client struct {
	ac AbstractCar
	ae AbstractEngine
}

func NewClient(cf CarFactory) *Client {
	return &Client{
		ac: cf.createCar(),
		ae: cf.createEngine(),
	}
}

func (c Client) RunMaxSpeed() int {
	return c.ae.getMaxSpeed()
}

func main() {
	ford_car := Ford{}
	c1 := NewClient(&ford_car)

	fmt.Println(c1.RunMaxSpeed())

	fmt.Println(c1.ac.ToString())

	fmt.Println(c1.ac.GetType())
}

/*
Golang is not the best programming language for abstract alsses as it doesn't have abstract class
So it can't store value in interface and the all point of this pattern breaking down

Adding new parameter to Pattern really easy, as I don't need to change a lot

Need to add logic only for Abstrac class and call functions in client's classes
*/
