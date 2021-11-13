package bus

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/x-research-team/contract"
)

// Init Инициализация шин
func Init(config json.RawMessage) {
	if err := json.Unmarshal(config, Trace); err != nil {
		log.Printf("[SYS] Logger are not initialized; Error: %v", err)
	}
	Sys = make(TSys)
	Error = make(TError)
	Info = make(TInfo)
	Debug = make(TDebug)
	Signals = make(contract.ISignalsBus, 0, 0)
}

// New Новый пул сигалов
func Add(x ...contract.ISignalBus) {
	Signals = append(Signals, x...)
}

// Bus Шина генерации магистралей
var Signals contract.ISignalsBus

// TSignal Сигнал для магистрали биллинга
type TSignal struct {
	pid     uuid.UUID
	message contract.IMessage
	channel contract.ISignalBus
}

// TMessage Сообщение для сигнала
type TMessage struct {
	id      uuid.UUID
	route   string
	command string
	data    string
}

// ID Correlation ID for kernel message
func (message TMessage) ID() uuid.UUID {
	return message.id
}

// Route Путь сообщения
func (message TMessage) Route() string {
	return message.route
}

// Command Команда которую нужно выполнить с данными в сообщении
func (message TMessage) Command() string {
	return message.command
}

// Data Данные в сообщении сигнала
func (message TMessage) Data() string {
	return message.data
}

// Message Сообщение в сигнале
func Message(route, command, data string) *TMessage {
	return &TMessage{uuid.New(), route, command, data}
}

// Signal Новый сигнал для магистрали
func Signal(message contract.IMessage) contract.ISignal {
	return &TSignal{
		pid:     uuid.New(),
		message: message,
	}
}

// Message Сообщение сигнала
func (signal TSignal) Message() contract.IMessage {
	return signal.message
}

func (signal TSignal) Read() string {
	return signal.Message().Data()
}

// Channel c
func (signal *TSignal) Channel() <-chan contract.ISignal {
	return nil
}

// Send c
func (signal *TSignal) Send() error {
	return nil
}

// Pid c
func (signal TSignal) Pid() string {
	return signal.pid.String()
}

// Name c
func (signal TSignal) Name() string {
	return fmt.Sprintf("Signal.PID[%s]", signal.Pid())
}

// Up c
func (signal *TSignal) Up(graceful bool) error {
	return nil
}

// Down c
func (signal *TSignal) Down(graceful bool) error {
	return nil
}

// Sleep c
func (signal *TSignal) Sleep(time.Duration) error {
	return nil
}

// Restart c
func (signal *TSignal) Restart(graceful bool) error {
	return nil
}

// Pause c
func (signal *TSignal) Pause() error {
	return nil
}

// Cron c
func (signal *TSignal) Cron(rule string) error {
	return nil
}

// Stop c
func (signal *TSignal) Stop() error {
	return nil
}

// Kill c
func (signal *TSignal) Kill() error {
	return nil
}

// Sync c
func (signal *TSignal) Sync(with string) error {
	return nil
}

// Backup c
func (signal *TSignal) Backup(to string) error {
	return nil
}
