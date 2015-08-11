package play

import (
)

func RunSwitch(n int) string {
	switch n {
	case 0: return "Zero"
	case 1: return "One"
	}
	return "Many"
}

func RunSwitch2() string {
	var t interface{}
	t = RunSwitch(0)
	switch t.(type) {
	default: return "unknown"
	case int: return "int"
	case string: return "string"
	}
}
