package frontend

import "zat/backend"

func AssciCell(c backend.Cell) string {
	switch c.(type) {
	case *backend.Monster:
		return "*"
	case *backend.Warrier:
		return "%"
	default:
		return "?"
	}
}

