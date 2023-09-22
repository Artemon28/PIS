// package main

// import (
// 	"fmt"
// 	"os"
// 	"sync"
// 	"time"
// )

// // Log представляет собой структуру для журнала.
// type Log struct {
// 	filename string
// 	file     *os.File
// 	mu       sync.Mutex
// }

// // SingletonLog возвращает единственный экземпляр Log.
// var SingletonLog *Log

// // GetLogInstance возвращает экземпляр Log, создавая его при необходимости.
// func GetLogInstance() *Log {
// 	if SingletonLog == nil {
// 		SingletonLog = &Log{filename: "log.txt"}
// 		SingletonLog.openLogFile()
// 	}
// 	return SingletonLog
// }

// // LogExecution записывает информацию в журнал.
// func (log *Log) LogExecution(message string) {
// 	log.mu.Lock()
// 	defer log.mu.Unlock()
// 	if log.file != nil {
// 		log.logMessage(message)
// 	}
// }

// func (log *Log) openLogFile() {
// 	file, err := os.OpenFile(log.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println("Ошибка при открытии файла журнала:", err)
// 	} else {
// 		log.file = file
// 		log.logMessage("Начало работы")
// 	}
// }

// func (log *Log) logMessage(message string) {
// 	if log.file != nil {
// 		currentTime := time.Now()
// 		log.file.WriteString(fmt.Sprintf("\nLog Entry : %s %s", currentTime.Format("15:04:05"), currentTime.Format("02.01.2006")))
// 		log.file.WriteString("\nДействие: " + message)
// 		log.file.WriteString("\n-------------------------------")
// 	}
// }

// // Operation представляет арифметические операции.
// type Operation struct{}

// // Run выполняет арифметическую операцию и записывает информацию в журнал.
// func (op *Operation) Run(operationCode rune, operand int) float64 {
// 	log := GetLogInstance()
// 	result := 0.0

// 	switch operationCode {
// 	case '+':
// 		result += float64(operand)
// 		log.LogExecution(fmt.Sprintf("Сложение %d", operand))
// 	case '-':
// 		result -= float64(operand)
// 		log.LogExecution(fmt.Sprintf("Вычитание %d", operand))
// 	case '*':
// 		result *= float64(operand)
// 	case '/':
// 	case ':':
// 		result /= float64(operand)
// 	}
// 	return result
// }

// func main() {
// 	log := GetLogInstance()
// 	log.LogExecution("Метод Main()")

// 	op := Operation{}
// 	result1 := op.Run('-', 35)
// 	result2 := op.Run('+', 30)

// 	fmt.Printf("Результат 1: %.2f\n", result1)
// 	fmt.Printf("Результат 2: %.2f\n", result2)
// }

// =======================
// version 2
// ======================

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Log представляет собой структуру для журнала.
type Log struct {
	filename string
	file     *os.File
	mu       sync.Mutex
}

// SingletonLog представляет собой экземпляр Log.
var SingletonLog *Log

// GetLogInstance возвращает экземпляр Log, создавая его при необходимости.
func GetLogInstance() *Log {
	if SingletonLog == nil {
		SingletonLog = &Log{filename: "log.txt"}
		SingletonLog.openLogFile()
	}
	return SingletonLog
}

// Приватный конструктор, чтобы запретить создание объектов извне.
func NewLog() *Log {
	return &Log{}
}

// LogExecution записывает информацию в журнал.
func (log *Log) LogExecution(message string) {
	log.mu.Lock()
	defer log.mu.Unlock()
	if log.file != nil {
		log.logMessage(message)
	}
}

func (log *Log) openLogFile() {
	file, err := os.OpenFile(log.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла журнала:", err)
	} else {
		log.file = file
		log.logMessage("Начало работы")
	}
}

func (log *Log) logMessage(message string) {
	if log.file != nil {
		currentTime := time.Now()
		log.file.WriteString(fmt.Sprintf("\nLog Entry : %s %s", currentTime.Format("15:04:05"), currentTime.Format("02.01.2006")))
		log.file.WriteString("\nДействие: " + message)
		log.file.WriteString("\n-------------------------------")
	}
}

// Operation представляет арифметические операции.
type Operation struct{}

// Run выполняет арифметическую операцию и записывает информацию в журнал.
func (op *Operation) Run(operationCode rune, operand int) float64 {
	log := GetLogInstance()
	result := 0.0

	switch operationCode {
	case '+':
		result = float64(operand)
		log.LogExecution(fmt.Sprintf("Сложение %d", operand))
	case '-':
		result = -float64(operand)
		log.LogExecution(fmt.Sprintf("Вычитание %d", operand))
	case '*':
		// Здесь можно добавить код для умножения, если необходимо.
	case '/':
		// Здесь можно добавить код для деления, если необходимо.
	}

	return result
}

func main() {
	log := GetLogInstance()
	log.LogExecution("Метод Main()")

	op := Operation{}
	result1 := op.Run('-', 35)
	result2 := op.Run('+', 30)

	fmt.Printf("Результат 1: %.2f\n", result1)
	fmt.Printf("Результат 2: %.2f\n", result2)
}

//Golang doesn't have Lazy thing
