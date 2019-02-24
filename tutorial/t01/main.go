package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// --------------------------------------------------------
// 依存関係を持つ構造体群を定義
//
// [Event]
//     +-- [Greeter]
//             +-- [Message]
// --------------------------------------------------------

// EventはGreeterを持つ
type Event struct {
	Greeter Greeter
}

// GreeterはMessageを持つ
type Greeter struct {
	Grumpy  bool
	Message Message
}

// Message型
type Message string

// --------------------------------------------------------
// 依存関係の生成関数
// --------------------------------------------------------

// Greeterを受け取ってEventを生成する
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

// Messageを受け取ってGreeterを生成する
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

// 文字列を受け取ってMessageを生成する
func NewMessage(phrase string) Message {
	return Message(phrase)
}

// --------------------------------------------------------
// 構造体が持つメソッド
// --------------------------------------------------------

// Eventは開始メソッドを持つ
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// Greeterは挨拶メソッドを持つ
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// --------------------------------------------------------
// アプリケーション起動関数
// --------------------------------------------------------

func main() {
	e, err := InitializeEvent("hi there!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
