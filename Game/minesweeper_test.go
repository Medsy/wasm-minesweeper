package Game

import (
	"fmt"
	"testing"
)

func TestPosListSort(t *testing.T) {
	l := PosList{
		{2, 2},
		{1, 1},
		{2, 1},
		{1, 2},
		{3, 4},
		{1, 9},
	}

	p := PosList{
		{2, 2},
		{1, 1},
		{3, 4},
		{1, 2},
		{2, 1},
		{1, 9},
	}

	l.sort()
	p.sort()

	fmt.Println(p)
	fmt.Println(l)
}
