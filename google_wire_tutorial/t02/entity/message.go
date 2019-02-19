package entity

// Message型
type Message string

// 文字列を受け取ってMessageを生成する
func NewMessage(phrase string) Message {
	return Message(phrase)
}
