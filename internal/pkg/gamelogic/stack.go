package gamelogic

type xy struct {
	X int
	Y int
}

type xyStack []xy

func (s xyStack) peek() xy {
	cur := len(s) - 1
	return s[cur]
}

func (s *xyStack) push(e xy) {
	n := append(*s, e)
	*s = n
}

func (s *xyStack) pop() xy {
	a := *s
	cur := len(a) - 1
	e := a[cur]
	*s = a[:cur]
	return e
}

func (s xyStack) empty() bool {
	return len(s) == 0
}
