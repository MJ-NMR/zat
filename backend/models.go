package backend

type Frame [][]Cell

type Cell any


type Charactor struct {
	Name string
	HP    uint
	Speed uint
	}

type Monster struct {
	Charactor
}

func (m *Monster) damage(d uint) {
	if d > 0 {
		if d > m.HP {
			m.HP = 0
			return
		}
		m.HP -= d
	}
}

type Warrier struct {
	Charactor
}
