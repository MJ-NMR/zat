package backend

import "math/rand"

type Status [][]Cell

type Cell interface {
	index() (y, x int)
}

type cell struct {
	location struct{ x, y int }
}

func (c cell) index() (y, x int) {
	return c.location.y, c.location.x
}

type Warrier struct {
	Name  string
	Range uint
	speed int
	cell
}

type Monster struct {
	Name  string
	Hp    int
	speed int
	cell
}

type Ground struct {
	Name    string
	inRange []Warrier
	cell
}

func RandomStatus(rows, cols int) Status {
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
			default:
				st[y][x] = &Ground{}
			}
		}
	}
	return st
}
