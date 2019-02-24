package main

import (
	"fmt"
	"os"
)

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
