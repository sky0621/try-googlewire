package model

type Notice struct {
	ID    string
	Title string
	Text  string
}

type NoticeCondition struct {
	IDs []string
}
