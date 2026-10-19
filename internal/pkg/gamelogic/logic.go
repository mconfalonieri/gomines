package gamelogic

import (
	"strconv"

	"github.com/mconfalonieri/gomines/internal/pkg/minefield"
)

const (
	CellUnknown = "?"
	CellMarked  = "F"
	CellMine    = "*"
	CellEmpty   = " "
	CellXMarked = "X"
	CellMineOK  = "W"
)

type PlayField struct {
	width         int
	height        int
	unmarkedMines int
	hiddenCells   int
	VisibleField  [][]string
	realField     *minefield.MineField
	statusMsg     string
}

func New(w, h, m int) (PlayField, error) {
	f, err := minefield.New(w, h, m)
	if err != nil {
		return PlayField{}, err
	}
	freeCells := w*h - m // cells without mines
	return PlayField{
		width:         w,
		height:        h,
		unmarkedMines: m,
		hiddenCells:   freeCells,
		VisibleField:  initialRender(w, h),
		realField:     f,
	}, nil
}

func initialRender(w, h int) [][]string {
	f := make([][]string, w)
	for x := range f {
		f[x] = make([]string, h)
		for y := range f[x] {
			f[x][y] = CellUnknown
		}
	}
	return f
}

func (p *PlayField) Play(x, y int, mark bool) bool {
	col := p.VisibleField[x]
	c := col[y]
	if discovered(c) {
		return false
	}
	return p.play(x, y, mark)
}

func (p PlayField) addCellsToStack(point xy, s *xyStack) {
	for _, xa := range []int{-1, 0, 1} {
		for _, ya := range []int{-1, 0, 1} {
			x := point.X + xa
			y := point.Y + ya
			if x < 0 || x >= p.width || y < 0 || y >= p.height || discovered(p.VisibleField[x][y]) {
				continue
			}
			v := xy{X: x, Y: y}
			s.push(v)
		}
	}
}

func (p *PlayField) play(x, y int, mark bool) bool {
	if mark {
		p.markMine(x, y)
		p.updateStats()
		return false
	} else {
		v := p.openCell(x, y)
		if v == true {
			return true
		}
		s := xyStack{}
		s.push(xy{X: x, Y: y})
		for !s.empty() {
			point := s.pop()
			if discovered(p.VisibleField[point.X][point.Y]) {
				continue
			}
			mines := p.scanNearCells(point)
			if mines > 0 {
				p.VisibleField[point.X][point.Y] = strconv.Itoa(mines)
			} else {
				p.VisibleField[point.X][point.Y] = CellEmpty
				p.addCellsToStack(point, &s)
			}
		}
		p.updateStats()
	}
	return false
}

func (p *PlayField) markMine(x, y int) {
	if p.VisibleField[x][y] == CellMarked {
		p.VisibleField[x][y] = CellUnknown
	} else if p.unmarkedMines > 0 {
		p.VisibleField[x][y] = CellMarked
	}
}

func (p *PlayField) openCell(x, y int) bool {
	if p.realField.Field[x][y] {
		p.VisibleField[x][y] = CellMine
		return true
	}
	return false
}

func discovered(s string) bool {
	return s != CellMarked && s != CellUnknown
}

func (p PlayField) scanNearCells(point xy) int {
	mines := 0
	for _, xa := range []int{-1, 0, 1} {
		for _, ya := range []int{-1, 0, 1} {
			if xa == 0 && ya == 0 {
				continue
			}
			x := point.X + xa
			y := point.Y + ya
			if x < 0 || x >= p.width || y < 0 || y >= p.height || discovered(p.VisibleField[x][y]) {
				continue
			}
			if p.realField.Field[x][y] {
				mines++
			}
		}
	}
	return mines
}

func (p PlayField) UnmarkedMines() int {
	return p.unmarkedMines
}

func (p PlayField) HiddenCells() int {
	return p.hiddenCells
}

func (p PlayField) revealCell(x, y int) {
	marked := (p.VisibleField[x][y] == CellMarked)
	mined := p.realField.Field[x][y]
	if marked && !mined {
		p.VisibleField[x][y] = CellXMarked
	} else if !marked && mined {
		p.VisibleField[x][y] = CellMine
	} else if marked && mined {
		p.VisibleField[x][y] = CellMineOK
	} else {
		p.VisibleField[x][y] = CellEmpty
	}
}

func (p *PlayField) RevealCells() {
	for x := range p.realField.Field {
		for y := range p.realField.Field[x] {
			p.revealCell(x, y)
		}
	}
	p.hiddenCells = 0
	p.unmarkedMines = 0
}

func (p *PlayField) updateStats() {
	um := p.realField.Mines()
	hc := p.width*p.height - um
	for x := range p.VisibleField {
		for y := range p.VisibleField[x] {
			if discovered(p.VisibleField[x][y]) {
				hc--
			} else if p.VisibleField[x][y] == CellMarked {
				um--
			}
		}
	}
	p.hiddenCells = hc
	p.unmarkedMines = um
}
