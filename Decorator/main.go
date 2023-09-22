package main

import (
	"fmt"
)

// AutoBase is the abstract class that defines the base functionality
type AutoBase interface {
	GetCost() float64
	ToString() string
}

// Renault represents the base car
type Renault struct {
	Name        string
	Description string
	CostBase    float64
}

func (r Renault) GetCost() float64 {
	return r.CostBase * 1.18
}

func (r Renault) ToString() string {
	return fmt.Sprintf("Your car:\n%s\nDescription: %s\nCost: %.2f\n", r.Name, r.Description, r.GetCost())
}

// DecoratorOptions is the abstract decorator class
type DecoratorOptions struct {
	AutoProperty AutoBase
	Title        string
}

func (d *DecoratorOptions) GetCost() float64 {
	return d.AutoProperty.GetCost()
}

func (d *DecoratorOptions) ToString() string {
	return d.AutoProperty.ToString()
}

// MediaNAV represents the multimedia navigation system package decorator
type MediaNAV struct {
	*DecoratorOptions
}

func NewMediaNAV(base AutoBase, title string) *MediaNAV {
	return &MediaNAV{
		DecoratorOptions: &DecoratorOptions{
			AutoProperty: base,
			Title:        title,
		},
	}
}

func (m *MediaNAV) GetCost() float64 {
	return m.AutoProperty.GetCost() + 15.99
}

func (m *MediaNAV) ToString() string {
	return fmt.Sprintf("%s. Modern\n%s. Updated multimedia navigation system\nCost: %.2f\n", m.AutoProperty.ToString(), m.Title, m.GetCost())
}

// SystemSecurity represents the security system package decorator
type SystemSecurity struct {
	*DecoratorOptions
}

func NewSystemSecurity(base AutoBase, title string) *SystemSecurity {
	return &SystemSecurity{
		DecoratorOptions: &DecoratorOptions{
			AutoProperty: base,
			Title:        title,
		},
	}
}

func (s *SystemSecurity) GetCost() float64 {
	return s.AutoProperty.GetCost() + 20.99
}

func (s *SystemSecurity) ToString() string {
	return fmt.Sprintf("%s. Enhanced security\n%s. Front side safety, ESP - electronic stability control system\nCost: %.2f\n", s.AutoProperty.ToString(), s.Title, s.GetCost())
}

func main() {
	reno := &Renault{
		Name:        "Renault",
		Description: "Renault LOGAN Active",
		CostBase:    499.0,
	}

	Print(reno)

	myRenoWithMediaNAV := NewMediaNAV(reno, "Navigation")
	Print(myRenoWithMediaNAV)

	myRenoWithSecurity := NewSystemSecurity(myRenoWithMediaNAV, "Safety")
	Print(myRenoWithSecurity)
}

func Print(av AutoBase) {
	fmt.Println(av.ToString())
}
