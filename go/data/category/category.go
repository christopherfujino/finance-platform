package category

import (
	"fmt"
)

type T int

const (
	Null      T = iota
	EatingOut
)

func (c T) ToString() string {
	switch c {
	case Null:
		return "Uncategorized" // TODO?
	case EatingOut:
		return "Eating out"
	}
	panic(fmt.Sprintf("Unreachable: %d", c))
}
