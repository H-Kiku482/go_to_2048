package game

import "fmt"

type Field [4][4]int

func (f Field) String() string {
	var str string
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			str += fmt.Sprintf("[%s] ", color(f[i][j]))
		}
		str += fmt.Sprintf("[%s]\n", color(f[i][3]))
	}
	return str
}

func (f *Field) moveUp() {
	var next Field
	for i := 1; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if f[i][j] == 0 {
				continue
			} else {
				if f[i-1][j] == 0 {
					next = *f
					next[i-1][j] = next[i][j]
					next[i][j] = 0
					*f = next
					f.moveUp()
				}
			}
		}
	}
}

func (f *Field) addUp() int {
	var (
		next  Field
		score int
	)
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if f[i][j] == 0 {
				continue
			} else {
				if f[i][j] == f[i+1][j] {
					next = *f
					next[i][j] = next[i][j] * 2
					next[i+1][j] = 0
					score += next[i][j]
					*f = next
					f.moveUp()
				}
			}
		}
	}
	return score
}

func (f *Field) moveDown() {
	var next Field
	for i := 2; i >= 0; i-- {
		for j := 0; j < 4; j++ {
			if f[i][j] == 0 {
				continue
			} else {
				if f[i+1][j] == 0 {
					next = *f
					next[i+1][j] = next[i][j]
					next[i][j] = 0
					*f = next
					f.moveDown()
				}
			}
		}
	}
}

func (f *Field) addDown() int {
	var (
		next  Field
		score int
	)
	for i := 3; i > 0; i-- {
		for j := 0; j < 4; j++ {
			if f[i][j] == 0 {
				continue
			} else {
				if f[i][j] == f[i-1][j] {
					next = *f
					next[i][j] = next[i][j] * 2
					next[i-1][j] = 0
					score += next[i][j]
					*f = next
					f.moveDown()
				}
			}
		}
	}
	return score
}

func (f *Field) moveLeft() {
	var next Field
	for i := 0; i < 4; i++ {
		for j := 1; j < 4; j++ {
			if f[i][j] == 0 {
				continue
			} else {
				if f[i][j-1] == 0 {
					next = *f
					next[i][j-1] = next[i][j]
					next[i][j] = 0
					*f = next
					f.moveLeft()
				}
			}
		}
	}
}

func (f *Field) addLeft() int {
	var (
		next  Field
		score int
	)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if f[i][j] == 0 {
				continue
			} else {
				if f[i][j] == f[i][j+1] {
					next = *f
					next[i][j] = next[i][j] * 2
					next[i][j+1] = 0
					score += next[i][j]
					*f = next
					f.moveLeft()
				}
			}
		}
	}
	return score
}

func (f *Field) moveRight() {
	var next Field
	for i := 0; i < 4; i++ {
		for j := 2; j >= 0; j-- {
			if f[i][j] == 0 {
				continue
			} else {
				if f[i][j+1] == 0 {
					next = *f
					next[i][j+1] = next[i][j]
					next[i][j] = 0
					*f = next
					f.moveRight()
				}
			}
		}
	}
}

func (f *Field) addRight() int {
	var (
		next  Field
		score int
	)
	for i := 0; i < 4; i++ {
		for j := 3; j > 0; j-- {
			if f[i][j] == 0 {
				continue
			} else {
				if f[i][j] == f[i][j-1] {
					next = *f
					next[i][j] = next[i][j] * 2
					next[i][j-1] = 0
					score += next[i][j]
					*f = next
					f.moveRight()
				}
			}
		}
	}
	return score
}

func (f *Field) useBomb() {
	min := 2048
	max := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if f[i][j] != 0 {
				if f[i][j] < min {
					min = f[i][j]
				}
				if f[i][j] > max {
					max = f[i][j]
				}
			}
		}
	}

	if min != max {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if f[i][j] == min {
					f[i][j] = 0
				}
			}
		}
	}
}

func color(r int) string {
	switch r {
	case 0:
		return fmt.Sprintf("%4d", r)
	default:
		return fmt.Sprintf("\x1b[31m%4d\x1b[0m", r)
	}
}
