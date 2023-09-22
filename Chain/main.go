package main

import "fmt"

// Receiver represents the possible payment methods.
type Receiver struct {
	BankTransfer   bool
	MoneyTransfer  bool
	PayPalTransfer bool
}

// PaymentHandler is the abstract handler with a successor.
type PaymentHandler interface {
	SetSuccessor(handler PaymentHandler)
	Handle(receiver Receiver)
}

// BasePaymentHandler is a base handler implementing the successor chaining.
type BasePaymentHandler struct {
	successor PaymentHandler
}

func (h *BasePaymentHandler) SetSuccessor(handler PaymentHandler) {
	h.successor = handler
}

// BankPaymentHandler handles bank transfers.
type BankPaymentHandler struct {
	BasePaymentHandler
}

func (h *BankPaymentHandler) Handle(receiver Receiver) {
	if receiver.BankTransfer {
		fmt.Println("Performing bank transfer")
	} else if h.successor != nil {
		h.successor.Handle(receiver)
	}
}

// MoneyPaymentHandler handles money transfers.
type MoneyPaymentHandler struct {
	BasePaymentHandler
}

func (h *MoneyPaymentHandler) Handle(receiver Receiver) {
	if receiver.MoneyTransfer {
		fmt.Println("Performing money transfer")
	} else if h.successor != nil {
		h.successor.Handle(receiver)
	}
}

// PayPalPaymentHandler handles PayPal payments.
type PayPalPaymentHandler struct {
	BasePaymentHandler
}

func (h *PayPalPaymentHandler) Handle(receiver Receiver) {
	if receiver.PayPalTransfer {
		fmt.Println("Performing PayPal transfer")
	} else if h.successor != nil {
		h.successor.Handle(receiver)
	}
}

func main() {
	receiver := Receiver{
		BankTransfer:   false,
		MoneyTransfer:  true,
		PayPalTransfer: true,
	}

	bankPaymentHandler := &BankPaymentHandler{}
	moneyPaymentHandler := &MoneyPaymentHandler{}
	payPalPaymentHandler := &PayPalPaymentHandler{}

	bankPaymentHandler.SetSuccessor(moneyPaymentHandler)
	moneyPaymentHandler.SetSuccessor(payPalPaymentHandler)
	payPalPaymentHandler.SetSuccessor(bankPaymentHandler)

	fmt.Println("Payment process:")
	bankPaymentHandler.Handle(receiver)
}

// execute inly first one, no matter what we add
