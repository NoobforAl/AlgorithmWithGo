package main

import (
	"math"
)

const INF uint = math.MaxUint

/*func init() {
	rand.Seed(time.Now().Unix())
	example = make(ex, 5)

	listNum := make([]uint, 40)
	listNum[39] = INF
	listNum[0] = 40
	for i := 1; i < 39; i++ {
		listNum[i] = uint(i)
	}

	example[1] = [][]uint{
		{0, 8, 10, 15, 20},
		{15, 0, INF, 10, 8},
		{4, INF, 0, 8, INF},
		{INF, 6, 6, 0, INF},
		{15, 3, INF, INF, 0},
	}

	for k := 2; k < 5; k++ {
		size := k * k
		maps := make([][]uint, size)

		for i := 0; i < size; i++ {
			maps[i] = make([]uint, size)
			for j := 0; j < size; j++ {
				if i == j {
					maps[i][j] = 0
				} else {
					maps[i][j] = listNum[rand.Intn(40)]
				}
			}
		}
		example[k] = maps
	}
}
*/

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
