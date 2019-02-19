package entity

import (
	"errors"
	"fmt"
)

// EventはGreeterを持つ
type Event struct {
	Greeter Greeter
}

// Eventは開始メソッドを持つ
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// Greeterを受け取ってEventを生成する
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}
