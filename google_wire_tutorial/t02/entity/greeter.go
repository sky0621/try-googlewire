package entity

import "time"

// GreeterはMessageを持つ
type Greeter struct {
	Grumpy  bool
	Message Message
}

// Greeterは挨拶メソッドを持つ
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// Messageを受け取ってGreeterを生成する
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}
