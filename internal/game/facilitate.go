package game

import (
	"fmt"
	"math/rand"
	"time"
)

func (g *Game) newBlock() {
	var zeroField [][2]int // index of g.Body
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.Field[i][j] == 0 {
				var p [2]int
				p[0] = i
				p[1] = j
				zeroField = append(zeroField, p)
			}
		}
	}

	rand.Seed(time.Now().UnixMicro())
	r := rand.Intn(len(zeroField))
	if time.Now().Nanosecond()%2 == 0 {
		g.Field[zeroField[r][0]][zeroField[r][1]] = 2
	} else {
		g.Field[zeroField[r][0]][zeroField[r][1]] = 4
	}
}

func (g *Game) checkState() bool {
	goal := false
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.Field[i][j] == 2048 {
				goal = true
			}
		}
	}
	return goal
}

func (g *Game) GameClear() {
	fmt.Println("You arrived 2048!")
	g.Movable = false
}

func (g *Game) GameOver() {
	fmt.Println("Failed to go to 2048!")
	g.Movable = false
}

func (g *Game) printResult() {
	totalScore := g.Score + (g.Bomb * 256) - g.Move
	fmt.Println("=== RESULT ===")
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("SCORE : %d\n", g.Score)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("BOMB %d * 512 : %d\n", g.Bomb, g.Bomb*512)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("MOVE %d : -%d\n", g.Move, g.Move)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("TOTAL SCORE :", totalScore)
}
