package main

import (
	"fmt"
	"strconv"
)

type Drive struct {
	twist string
}

func NewDrive() *Drive {
	return &Drive{twist: "initial position"}
}

func (d *Drive) TurnLeft() {
	d.twist = "Turned left"
	fmt.Println("Turned left")
}

func (d *Drive) TurnRight() {
	d.twist = "Turned right"
	fmt.Println("Turned right")
}

func (d *Drive) Stop() {
	d.twist = "Stopped"
	fmt.Println("Stopped")
}

func (d *Drive) String() string {
	return fmt.Sprintf("Drive: %s", d.twist)
}

type Power struct {
	microwavePower int
}

func NewPower() *Power {
	fmt.Println("Power now is 0")
	return &Power{microwavePower: 0}
}

func (p *Power) SetPower(power int) {
	fmt.Println("Power now is " + strconv.Itoa(power))
	p.microwavePower = power
}

func Itoa(power int) {
	panic("unimplemented")
}

func (p *Power) String() string {
	return fmt.Sprintf("Microwave power set to %dw", p.microwavePower)
}

type Notification struct {
	message string
}

func NewNotification() *Notification {
	return &Notification{message: ""}
}

func (n *Notification) StartNotification() {
	fmt.Println("Operation started")
	n.message = "Operation started"
}

func (n *Notification) StopNotification() {
	fmt.Println("Operation finished")
	n.message = "Operation finished"
}

func (n *Notification) String() string {
	return fmt.Sprintf("Information: %s", n.message)
}

type Microwave struct {
	drive        *Drive
	power        *Power
	notification *Notification
}

func NewMicrowave() *Microwave {
	return &Microwave{
		drive:        NewDrive(),
		power:        NewPower(),
		notification: NewNotification(),
	}
}

func (m *Microwave) Defrost() {
	m.notification.StartNotification()

	m.power.SetPower(1000)
	m.drive.TurnRight()
	m.drive.TurnRight()
	m.power.SetPower(500)
	m.drive.Stop()
	m.drive.TurnLeft()
	m.drive.TurnLeft()
	m.power.SetPower(200)
	m.drive.Stop()
	m.drive.TurnRight()
	m.drive.TurnRight()
	m.drive.Stop()
	m.power.SetPower(0)

	m.notification.StopNotification()
}

func main() {
	microwave := NewMicrowave()

	fmt.Println("Microwave Operation:")
	microwave.Defrost()
	fmt.Println(microwave.power)
	fmt.Println(microwave.drive)
	fmt.Println(microwave.notification)
}
