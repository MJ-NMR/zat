package backend

import "math/rand"

type Status [][]Cell

type Cell interface {
	index() (y, x int)
}


type Warrier struct {
	Name     string
	location struct{ x, y int }
	Range    uint
	speed    int
}

func (w Warrier) index() (y, x int) {
	return w.location.y, w.location.x
}

type Monster struct {
	Name     string
	location struct{ x, y int }
	Hp       int
	speed    int
}

func (m Monster) index() (y, x int) {
	return m.location.y, m.location.x
}

type blank struct {
	Name    string
	location struct{ y, x int }
}

func (b blank) index() (y, x int) {
	return b.location.y, b.location.x
}

func InitStatus(rows, cols int) Status {
	st := make(Status, rows)
	for y := range rows {
		st[y] = make([]Cell, cols)
		for x := range cols {
			n := rand.Intn(3)
			switch n {
			case 0:
				st[y][x] = &Warrier{}
			case 1:
				st[y][x] = &Monster{}
			case 2:

			}
			st[y][x] = &Warrier{}
		}
	}
	return st
}
