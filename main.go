package main

import (
	"songguessr/game"
	"time"
)

func main() {
	go game.InitServer()

	for {
		time.Sleep(1000)
	}
}
