package main

import (
	"fmt"
	"math/rand"
)

// Game
type IGame interface {
	Brosok() int
}

// Kost'
type Kost struct {
}

func NewKost() *Kost {
	kost := &Kost{}
	return kost
}

func (k Kost) Brosok() int {
	res := rand.Intn(6) + 1
	return res
}

// Gamer
type Gamer struct {
	Name string
}

func NewGamer(name string) *Gamer {
	return &Gamer{
		Name: name,
	}
}

func (g Gamer) ToString() string {
	return g.Name
}

func (g Gamer) SeansGame(ig IGame) int {
	return ig.Brosok()
}

// Moneta
type Monet struct {
}

func NewMonet() *Monet {
	m := &Monet{}
	return m
}

func (m Monet) BrosokM() int {
	res := rand.Intn(2) + 1
	return res
}

// Adapter
type AdapterGame struct {
	mon Monet
}

func NewAdapter(m Monet) *AdapterGame {
	a := &AdapterGame{
		mon: m,
	}
	return a
}

func (a AdapterGame) Brosok() int {
	return a.mon.BrosokM()
}

func main() {
	kost := NewKost()
	g1 := NewGamer("Ivan")
	fmt.Printf("Выпало очков %d для игрока %s \n", g1.SeansGame(kost), g1.ToString())

	mon := NewMonet()
	bmon := NewAdapter(*mon)
	fmt.Printf("Монета показала %d для игрока %s \n", g1.SeansGame(bmon), g1.ToString())

}
