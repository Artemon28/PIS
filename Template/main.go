// package main

// import (
// 	"fmt"
// )

// type ProgressionAbstr interface {
// 	TemplateMethod()
// 	Print()
// 	InitializeProgression(a, b, h int)
// 	Progress()
// }

// type Progression struct {
// 	First    int
// 	Last     int
// 	H        int
// 	ProgList []int
// }

// func NewProgression(first, last, h int) *Progression {
// 	return &Progression{
// 		First:    first,
// 		Last:     last,
// 		H:        h,
// 		ProgList: []int{},
// 	}
// }

// func (p *Progression) TemplateMethod() {
// 	p.InitializeProgression(p.First, p.Last, p.H)
// 	p.Progress()
// 	p.Print()
// }

// func (p *Progression) Print() {
// 	fmt.Println("Последовательность:")
// 	for _, item := range p.ProgList {
// 		fmt.Print(" ", item)
// 	}
// 	fmt.Println()
// }

// func (p *Progression) InitializeProgression(a, b, h int) {
// 	p.First = a
// 	p.Last = b
// 	p.H = h
// }

// func (p *Progression) Progress() {
// }

// type ArithmeticProgression struct {
// 	*Progression
// }

// func NewArithmeticProgression(first, last, h int) *ArithmeticProgression {
// 	return &ArithmeticProgression{
// 		Progression: NewProgression(first, last, h),
// 	}
// }

// func (p *ArithmeticProgression) TemplateMethod() {
// 	p.InitializeProgression(p.First, p.Last, p.H)
// 	p.Progress()
// 	p.Print()
// }

// func (ap *ArithmeticProgression) Progress() {
// 	fF := ap.First
// 	for fF < ap.Last {
// 		ap.ProgList = append(ap.ProgList, fF)
// 		fF += ap.H
// 	}
// }

// type GeometricProgression struct {
// 	*Progression
// }

// func NewGeometricProgression(first, last, h int) *GeometricProgression {
// 	return &GeometricProgression{
// 		Progression: NewProgression(first, last, h),
// 	}
// }

// func (p *GeometricProgression) TemplateMethod() {
// 	p.InitializeProgression(p.First, p.Last, p.H)
// 	p.Progress()
// 	p.Print()
// }

// func (ap *GeometricProgression) Progress() {
// 	fF := ap.First
// 	for fF < ap.Last {
// 		ap.ProgList = append(ap.ProgList, fF)
// 		fF *= ap.H
// 	}
// }

// func main() {
// 	var f, l, h int

// 	fmt.Print("Введите начало прогрессии: ")
// 	fmt.Scan(&f)

// 	fmt.Print("Введите конец прогрессии: ")
// 	fmt.Scan(&l)

// 	fmt.Print("Введите шаг прогрессии: ")
// 	fmt.Scan(&h)

// 	arithmeticProgression := NewGeometricProgression(f, l, h)
// 	arithmeticProgression.TemplateMethod()
// }

//Control task

package main

import (
	"fmt"
)

type ProgressionAbstr interface {
	CutHair()
}

type BarberingProcess struct{}

func (b *BarberingProcess) TemplateMethod() {
	fmt.Println("1. Мойка волос")
	fmt.Println("2. Подготовка волос")
	b.CutHair()
	fmt.Println("4. Укладка (при необходимости)")
	fmt.Println("5. Завершение и оценка")
}

func (b *BarberingProcess) CutHair() {
}

type ShortHaircutProcess struct {
	BarberingProcess
}

func (b *ShortHaircutProcess) TemplateMethod() {
	fmt.Println("1. Мойка волос")
	fmt.Println("2. Подготовка волос")
	b.CutHair()
	fmt.Println("4. Укладка (при необходимости)")
	fmt.Println("5. Завершение и оценка")
}

func (s *ShortHaircutProcess) CutHair() {
	fmt.Println("3. Стрижка коротких волос")
}

type LongHaircutProcess struct {
	BarberingProcess
}

func (b *LongHaircutProcess) TemplateMethod() {
	fmt.Println("1. Мойка волос")
	fmt.Println("2. Подготовка волос")
	b.CutHair()
	fmt.Println("4. Укладка (при необходимости)")
	fmt.Println("5. Завершение и оценка")
}

func (l *LongHaircutProcess) CutHair() {
	fmt.Println("3. Стрижка длинных волос")
}

type BeardTrimmingProcess struct {
	BarberingProcess
}

func (b *BeardTrimmingProcess) TemplateMethod() {
	fmt.Println("1. Мойка волос")
	fmt.Println("2. Подготовка волос")
	b.CutHair()
	fmt.Println("4. Укладка (при необходимости)")
	fmt.Println("5. Завершение и оценка")
}

func (b *BeardTrimmingProcess) CutHair() {
	fmt.Println("3. Стрижка бороды")
}

func main() {
	fmt.Println("Процесс короткой стрижки:")
	shortHaircut := &ShortHaircutProcess{}
	shortHaircut.TemplateMethod()

	fmt.Println("\nПроцесс длинной стрижки:")
	longHaircut := &LongHaircutProcess{}
	longHaircut.TemplateMethod()

	fmt.Println("\nПроцесс стрижки бороды:")
	beardTrimming := &BeardTrimmingProcess{}
	beardTrimming.TemplateMethod()
}
