package bus

// TSys Канал шины системного уровня
type TSys chan interface{}
func (s TSys) Send(v interface{}) {
	select {
		case s <- v:
	}
}

// TError Канал шины ошибок
type TError chan interface{}
func (s TError) Send(v interface{}) {
	select {
		case s <- v:
	}
}

// TInfo Канал шины информационных сообщений
type TInfo chan string
func (s TInfo) Send(v string) {
	select {
		case s <- v:
	}
}

// TDebug Канал шины отдадочной информации
type TDebug chan interface{}
func (s TDebug) Send(v interface{}) {
	select {
		case s <- v:
	}
}

// Sys Шина системного уровня
var Sys TSys

// Error Шина ошибок
var Error TError

// Info Шина информационных сообщений
var Info TInfo

// Debug Шина отдадочной информации
var Debug TDebug

// tracerConfig Конфигурация трассера
type tracerConfig struct {
	Error bool `json:"error"`
	Info  bool `json:"info"`
	Debug bool `json:"debug"`
}

// Trace Конфигурация трассера
var Trace = new(tracerConfig)