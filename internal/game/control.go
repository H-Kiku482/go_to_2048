package game

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func (g *Game) Control(r rune, move *speculative) {
	before := g.Field
	act := false
	switch r {
	case 'w':
		g.Up(move)
		act = true
	case 'a':
		g.Left(move)
		act = true
	case 's':
		g.Down(move)
		act = true
	case 'd':
		g.Right(move)
		act = true
	case ' ':
		g.Space(move)
	case 27:
		g.Esc()
	default:
	}
	if act && (before != g.Field) {
		g.newBlock()
		g.Move += 1
	}

	if before == move.upRes && before == move.downRes && before == move.leftRes && before == move.rightRes && g.Bomb == 0 {
		g.GameOver()
	} else if g.checkState() {
		g.GameClear()
	} else {
		g.RenderScreen()
	}
}

func (g *Game) Up(s *speculative) {
	fmt.Print("\033[1A↑     \n\r")
	g.Field = s.upRes
	g.Score += s.upSc
}

func (g *Game) Down(s *speculative) {
	fmt.Print("\033[1A↓     \n\r")
	g.Field = s.downRes
	g.Score += s.downSc
}

func (g *Game) Left(s *speculative) {
	fmt.Print("\033[1A←     \n\r")
	g.Field = s.leftRes
	g.Score += s.leftSc
}

func (g *Game) Right(s *speculative) {
	fmt.Print("\033[1A→     \n\r")
	g.Field = s.rightRes
	g.Score += s.rightSc
}

func (g *Game) Space(s *speculative) {
	fmt.Print("\033[1Abomb   \n\r")
	if g.Bomb > 0 {
		before := g.Field
		g.Field = s.bombRes
		if before != g.Field {
			g.Bomb -= 1
		}
	}
}

func (g *Game) Esc() {
	g.InitGame()
}

func StayCansel(sc chan os.Signal, terminalState *term.State) {
	<-sc
	term.Restore(int(syscall.Stdin), terminalState)
	os.Exit(1)
}
