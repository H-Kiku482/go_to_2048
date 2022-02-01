package game

import (
	"fmt"
	"os"

	"github.com/mattn/go-tty"
)

type Game struct {
	Score   int
	Move    int
	Bomb    int
	EndGame bool
	Movable bool
	Field   Field
}

func (g *Game) String() string {
	screenStr := ""
	screenStr += fmt.Sprintf("SCORE:%7d MOVE:%8d\n", g.Score, g.Move)
	screenStr += fmt.Sprint(g.Field, "BOMB: ")
	for i := 0; i < g.Bomb; i++ {
		screenStr += "*"
	}
	return screenStr + "\n\n"
}

func (g *Game) RenderScreen() {
	spc := "                            "
	fmt.Print("\033[7A" + spc + "\n" + spc + "\n" + spc + "\n" + spc + "\n" + spc + "\n" + spc + "\n" + spc + "\n")
	fmt.Print("\033[7A", g, "\r")
}

func (g *Game) Play() {
	tty, err := tty.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tty.Close()

	for !(g.EndGame) {
		if g.Movable {
			r, _ := tty.ReadRune()
			g.Control(r, startSpeculativeExecution(g.Field))
		} else {
			g.printResult()
			g.EndGame = true
		}
	}
}

func InitGame() *Game {
	g := new(Game)
	g.Score = 0
	g.Move = 0
	g.Bomb = 3
	g.EndGame = false
	g.Movable = true
	g.newBlock()
	g.newBlock()
	fmt.Print("\n\n\n\n\n\n\n\r")
	return g
}

func (g *Game) InitGame() {
	g.Score = 0
	g.Move = 0
	g.Bomb = 3
	g.EndGame = false
	g.Movable = true
	g.Field[0][0] = 0
	g.Field[0][1] = 0
	g.Field[0][2] = 0
	g.Field[0][3] = 0
	g.Field[1][0] = 0
	g.Field[1][1] = 0
	g.Field[1][2] = 0
	g.Field[1][3] = 0
	g.Field[2][0] = 0
	g.Field[2][1] = 0
	g.Field[2][2] = 0
	g.Field[2][3] = 0
	g.Field[3][0] = 0
	g.Field[3][1] = 0
	g.Field[3][2] = 0
	g.Field[3][3] = 0
	g.newBlock()
	g.newBlock()
	fmt.Print("\n\n\n\n\n\n\n\n")
	g.RenderScreen()
}
