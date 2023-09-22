package main

import (
	"fmt"
)

// Command interface
type Command interface {
	Execute() float64
	UnExecute() float64
}

// ArithmeticUnit performs arithmetic operations
type ArithmeticUnit struct {
	Register float64
}

// AddCommand represents the addition operation
type AddCommand struct {
	unit    *ArithmeticUnit
	operand float64
}

func NewAddCommand(unit *ArithmeticUnit, operand float64) *AddCommand {
	return &AddCommand{
		unit:    unit,
		operand: operand,
	}
}

func (ac *AddCommand) Execute() float64 {
	ac.unit.Register += ac.operand
	return ac.unit.Register
}

func (ac *AddCommand) UnExecute() float64 {
	ac.unit.Register -= ac.operand
	return ac.unit.Register
}

// ControlUnit manages commands
type ControlUnit struct {
	commands []Command
	current  int
}

func (cu *ControlUnit) StoreCommand(command Command) {
	// Clear the redo stack when a new command is stored
	cu.commands = cu.commands[:cu.current]
	cu.commands = append(cu.commands, command)
	cu.current++
}

func (cu *ControlUnit) ExecuteCommand() float64 {
	if cu.current > 0 {
		cu.current--
		return cu.commands[cu.current].Execute()
	}
	return 0
}

func (cu *ControlUnit) Undo() float64 {
	if cu.current > 0 {
		cu.current--
		return cu.commands[cu.current].UnExecute()
	}
	return 0
}

func (cu *ControlUnit) Redo() float64 {
	if cu.current < len(cu.commands) {
		result := cu.commands[cu.current].Execute()
		cu.current++
		return result
	}
	return 0
}

// Calculator is the client
type Calculator struct {
	arithmeticUnit *ArithmeticUnit
	controlUnit    *ControlUnit
}

func NewCalculator() *Calculator {
	return &Calculator{
		arithmeticUnit: &ArithmeticUnit{},
		controlUnit:    &ControlUnit{},
	}
}

func (c *Calculator) Add(operand float64) float64 {
	command := NewAddCommand(c.arithmeticUnit, operand)
	c.controlUnit.StoreCommand(command)
	return c.controlUnit.ExecuteCommand()
}

func (c *Calculator) Undo() float64 {
	return c.controlUnit.Undo()
}

func (c *Calculator) Redo() float64 {
	return c.controlUnit.Redo()
}

// SubCommand represents the subtraction operation
type SubCommand struct {
	unit    *ArithmeticUnit
	operand float64
}

func NewSubCommand(unit *ArithmeticUnit, operand float64) *SubCommand {
	return &SubCommand{
		unit:    unit,
		operand: operand,
	}
}

func (sc *SubCommand) Execute() float64 {
	sc.unit.Register -= sc.operand
	return sc.unit.Register
}

func (sc *SubCommand) UnExecute() float64 {
	sc.unit.Register += sc.operand
	return sc.unit.Register
}

// MulCommand represents the multiplication operation
type MulCommand struct {
	unit    *ArithmeticUnit
	operand float64
}

func NewMulCommand(unit *ArithmeticUnit, operand float64) *MulCommand {
	return &MulCommand{
		unit:    unit,
		operand: operand,
	}
}

func (mc *MulCommand) Execute() float64 {
	mc.unit.Register *= mc.operand
	return mc.unit.Register
}

func (mc *MulCommand) UnExecute() float64 {
	mc.unit.Register /= mc.operand
	return mc.unit.Register
}

// DivCommand represents the division operation
type DivCommand struct {
	unit    *ArithmeticUnit
	operand float64
}

func NewDivCommand(unit *ArithmeticUnit, operand float64) *DivCommand {
	return &DivCommand{
		unit:    unit,
		operand: operand,
	}
}

func (dc *DivCommand) Execute() float64 {
	if dc.operand != 0 {
		dc.unit.Register /= dc.operand
		return dc.unit.Register
	}
	return 0
}

func (dc *DivCommand) UnExecute() float64 {
	if dc.operand != 0 {
		dc.unit.Register *= dc.operand
		return dc.unit.Register
	}
	return 0
}

func (c *Calculator) Subtract(operand float64) float64 {
	command := NewSubCommand(c.arithmeticUnit, operand)
	c.controlUnit.StoreCommand(command)
	return c.controlUnit.ExecuteCommand()
}

func (c *Calculator) Multiply(operand float64) float64 {
	command := NewMulCommand(c.arithmeticUnit, operand)
	c.controlUnit.StoreCommand(command)
	return c.controlUnit.ExecuteCommand()
}

func (c *Calculator) Divide(operand float64) float64 {
	if operand != 0 {
		command := NewDivCommand(c.arithmeticUnit, operand)
		c.controlUnit.StoreCommand(command)
		return c.controlUnit.ExecuteCommand()
	}
	return 0
}
func main() {
	calculator := NewCalculator()

	result := 0.0
	result = calculator.Add(6)
	fmt.Println(result)
	result = calculator.Divide(2)
	fmt.Println(result)
	result = calculator.Multiply(7)
	fmt.Println(result)

	result = calculator.Redo()
	fmt.Println(result)

	result = calculator.Undo()
	fmt.Println(result)
}

//work strange, redo can repeat prev action
