package bus

// TSys Канал шины системного уровня
type TSys chan interface{}

// TError Канал шины ошибок
type TError chan interface{}

// TInfo Канал шины информационных сообщений
type TInfo chan string

// TDebug Канал шины отдадочной информации
type TDebug chan interface{}

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
var Trace *tracerConfig

// LogInfo Логгировать информационное сообщение
func LogInfo(s string) {
	Info <- s
}

func LogSys(s string) {
	Sys <- s
}

func LogError(s string) {
	Error <- s
}

func LogDebug(s string) {
	Debug <- s
}
