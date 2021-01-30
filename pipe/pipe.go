package pipe

import (
	"log"

	"github.com/x-research-team/bus"
)

// Init Инициализация обработчиков шин
func Init() {
	go Sys()
	go Error()
	go Info()
	go Debug()
}

// Sys Обработчик ошибок системного уровня
func Sys() {
	for true {
		log.Printf("[SYS] %v\n", <-bus.Sys)
	}
}

// Error Обработчик шины ошибок
func Error() {
	for true {
		if bus.Trace.Error {
			log.Printf("[ERR] %v\n", <-bus.Error)
		}
	}
}

// Info Обработчик шины ошибок
func Info() {
	for true {
		if bus.Trace.Info {
			log.Printf("[INF] %v\n", <-bus.Info)
		}
	}
}

// Debug Обработчик шины ошибок
func Debug() {
	for true {
		if bus.Trace.Debug {
			log.Printf("[DBG] %v\n", <-bus.Debug)
		}
	}
}
