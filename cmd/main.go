package main

import (
	"go_to_2048/internal/game"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func main() {
	console := game.InitGame()
	console.RenderScreen()

	sChan := make(chan os.Signal, 1)
	signal.Notify(sChan, os.Interrupt)
	defer signal.Stop(sChan)

	terminalState, _ := term.GetState(int(syscall.Stdin))

	go game.StayCansel(sChan, terminalState)

	console.Play()

	term.Restore(int(syscall.Stdin), terminalState)
}
