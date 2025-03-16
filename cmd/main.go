package main

import (
	"muse/internal/telegram/handlers"
)

func main() {
	a := handlers.New()
	a.Start()
}
