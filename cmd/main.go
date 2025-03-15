package main

import (
	"muse/internal/telegram/handlers"
)

func main() {
	a := handlers.New()
	// TODO: add webhook, make config only through docker and volumes
	a.Start()
}
