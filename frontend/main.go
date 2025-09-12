package frontend

import "zat/backend"


func RenderStatus(st backend.Status) (s string) {
	for _, row := range st {
		for _, cell := range row {
			switch cell.(type) {
			case *backend.Monster:
				s += "*"
			case *backend.Warrier:
				s += "%"
			default:
				s +=  "?"
			}
		}
		s += "\n"
	}
	return s
}
