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
		select {
		case v := <-bus.Sys:
			log.Printf("[SYS] %v\n", v)
		}
	}
}

// Error Обработчик шины ошибок
func Error() {
	for true {
		select {
		case v := <-bus.Error:
			if bus.Trace.Error {
				log.Printf("[ERR] %v\n", v)
			}
		}
	}
}

// Info Обработчик шины ошибок
func Info() {
	for true {
		select {
		case v := <-bus.Info:
			if bus.Trace.Info {
				log.Printf("[INF] %v\n", v)
			}
		}
	}
}

// Debug Обработчик шины ошибок
func Debug() {
	for true {
		select {
		case v := <-bus.Debug:
			if bus.Trace.Debug {
				log.Printf("[DBG] %v\n", v)
			}
		}
	}
}
