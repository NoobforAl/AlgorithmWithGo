package main

import (
	"math"
)

const INF uint = math.MaxUint

func min(x, y uint) uint {
	if x > y {
		return y
	}
	return x
}

func floyd(rolmap *[][]uint) {
	for k := 0; k < len(*rolmap); k++ {
		for i := 0; i < len(*rolmap); i++ {
			for j := 0; j < len(*rolmap); j++ {
				(*rolmap)[i][j] = min((*rolmap)[i][j],
					(*rolmap)[i][k]+(*rolmap)[k][j])
			}
		}
	}
}

func main() {
	example := [][]uint{
		{0, 5, INF, 10},
		{INF, 0, 3, INF},
		{INF, INF, 0, 1},
		{INF, INF, INF, 0},
	}

	show(&example)
	floyd(&example)
	show(&example)
}
