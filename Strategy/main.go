// package main

// import "fmt"

// // StrategySort - интерфейс стратегии сортировки
// type StrategySort interface {
// 	Sort(arr []int)
// }

// // InsertionSort - конкретная стратегия сортировки вставками
// type InsertionSort struct{}

// func (is InsertionSort) Sort(arr []int) {
// 	for i := 1; i < len(arr); i++ {
// 		j := 0
// 		buffer := arr[i]
// 		for j = i - 1; j >= 0; j-- {
// 			if arr[j] < buffer {
// 				break
// 			}
// 			arr[j+1] = arr[j]
// 		}
// 		arr[j+1] = buffer
// 	}
// }

// // SelectionSort - конкретная стратегия сортировки выбором
// type SelectionSort struct{}

// func (ss SelectionSort) Sort(arr []int) {
// 	for i := 0; i < len(arr)-1; i++ {
// 		k := i
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[k] > arr[j] {
// 				k = j
// 			}
// 		}
// 		if k != i {
// 			temp := arr[k]
// 			arr[k] = arr[i]
// 			arr[i] = temp
// 		}
// 	}
// }

// // BubbleSort - конкретная стратегия сортировки пузырьком
// type BubbleSort struct{}

// func (bs BubbleSort) Sort(arr []int) {
// 	n := len(arr)
// 	for i := 0; i < n-1; i++ {
// 		swapped := false
// 		for j := 0; j < n-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// }

// // Context - контекст для работы с выбранной стратегией сортировки
// type Context struct {
// 	strategy StrategySort
// 	array    []int
// }

// func (c *Context) Sort() {
// 	c.strategy.Sort(c.array)
// }

// func (c *Context) PrintArray() {
// 	fmt.Println(c.strategy)
// 	fmt.Println(c.array)
// }

// func main() {
// 	arr1 := []int{31, 15, 10, 2, 4, 2, 14, 23, 12, 66}
// 	sort := SelectionSort{}
// 	context := Context{strategy: sort, array: arr1}
// 	context.Sort()
// 	context.PrintArray()

// 	arr2 := []int{1, 5, 10, 2, 4, 12, 14, 23, 12, 66}
// 	sort2 := InsertionSort{}
// 	context = Context{strategy: sort2, array: arr2}
// 	context.Sort()
// 	context.PrintArray()

// 	arr3 := []int{44, 12, 94, 8, 45, 32, 75, 89, 5, 24}
// 	sort3 := BubbleSort{}
// 	context = Context{strategy: sort3, array: arr3}
// 	context.Sort()
// 	context.PrintArray()
// }

package main

import (
	"fmt"
)

type Landmark struct {
	Name string
}

type Route struct {
	Landmarks []Landmark
}

type Map struct {
}

type TravelModeStrategy interface {
	DisplayRoute(route Route, mapInstance *Map)
}

type CarTravel struct {
}

func (ct *CarTravel) DisplayRoute(route Route, mapInstance *Map) {
	fmt.Println("Displaying car route on the map")
}

type BikeTravel struct {
}

func (bt *BikeTravel) DisplayRoute(route Route, mapInstance *Map) {
	fmt.Println("Displaying bike route on the map")
}

type WalkingTravel struct {
}

func (wt *WalkingTravel) DisplayRoute(route Route, mapInstance *Map) {
	fmt.Println("Displaying walking route on the map")
}

type SightseeingTravel struct {
}

func (st *SightseeingTravel) DisplayRoute(route Route, mapInstance *Map) {
	fmt.Println("Displaying sightseeing route on the map")
}

type Navigator struct {
	currentStrategy TravelModeStrategy
}

func (n *Navigator) SetTravelModeStrategy(strategy TravelModeStrategy) {
	n.currentStrategy = strategy
}

func (n *Navigator) DisplayRoute(route Route, mapInstance *Map) {
	n.currentStrategy.DisplayRoute(route, mapInstance)
}

func main() {
	mapInstance := &Map{}
	carTravel := &CarTravel{}
	//bikeTravel := &BikeTravel{}
	walkingTravel := &WalkingTravel{}
	//sightseeingTravel := &SightseeingTravel{}
	navigator := &Navigator{}
	navigator.SetTravelModeStrategy(carTravel)
	route := Route{
		Landmarks: []Landmark{
			{Name: "Start Point"},
			{Name: "Landmark 1"},
			{Name: "Landmark 2"},
			{Name: "Destination Point"},
		},
	}
	navigator.DisplayRoute(route, mapInstance)
	navigator.SetTravelModeStrategy(walkingTravel)
	navigator.DisplayRoute(route, mapInstance)
}
