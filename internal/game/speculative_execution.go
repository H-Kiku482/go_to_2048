package game

type speculative struct {
	upRes    Field
	downRes  Field
	leftRes  Field
	rightRes Field
	bombRes  Field

	upSc    int
	downSc  int
	leftSc  int
	rightSc int
}

func startSpeculativeExecution(current Field) *speculative {
	s := new(speculative)

	upC := make(chan Field, 1)
	downC := make(chan Field, 1)
	leftC := make(chan Field, 1)
	rightC := make(chan Field, 1)
	bombC := make(chan Field, 1)

	upScC := make(chan int, 1)
	downScC := make(chan int, 1)
	leftScC := make(chan int, 1)
	rightScC := make(chan int, 1)

	go s.up(current, upC, upScC)
	go s.down(current, downC, downScC)
	go s.left(current, leftC, leftScC)
	go s.right(current, rightC, rightScC)
	go s.bomb(current, bombC)

	s.upRes, s.downRes, s.leftRes, s.rightRes = <-upC, <-downC, <-leftC, <-rightC

	s.upSc, s.downSc, s.leftSc, s.rightSc = <-upScC, <-downScC, <-leftScC, <-rightScC

	s.bombRes = <-bombC

	return s
}

func (s *speculative) up(f Field, c chan Field, sc chan int) {
	f.moveUp()
	sc <- f.addUp()
	c <- f
}

func (s *speculative) down(f Field, c chan Field, sc chan int) {
	f.moveDown()
	sc <- f.addDown()
	c <- f
}

func (s *speculative) left(f Field, c chan Field, sc chan int) {
	f.moveLeft()
	sc <- f.addLeft()
	c <- f
}

func (s *speculative) right(f Field, c chan Field, sc chan int) {
	f.moveRight()
	sc <- f.addRight()
	c <- f
}

func (s *speculative) bomb(f Field, c chan Field) {
	f.useBomb()
	c <- f
}
