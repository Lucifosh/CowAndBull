package main

import (
	"time"
)

type Game struct {
	number string
	state  string
	time   time.Time
	count  int
}
type Board struct {
	ID    int
	score int
}

func main() {
	TelegramBot()
}
