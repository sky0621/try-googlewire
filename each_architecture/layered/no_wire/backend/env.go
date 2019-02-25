package main

type Env interface {
}

type env struct {
}

func ReadEnv() Env {
	return &env{}
}
