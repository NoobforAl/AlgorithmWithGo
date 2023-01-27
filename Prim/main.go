package main

import (
	"math"
)

const INF uint = math.MaxUint

func min(ln int, key *[]uint, mst *[]bool) uint {
	var m, mIndex uint = INF, 0

	for i := 0; i < ln; i++ {
		if !(*mst)[i] && (*key)[i] < m {
			m, mIndex = (*key)[i], uint(i)
		}
	}

	return mIndex
}

func prim(g *[][]uint) *[]int {

	var (
		par []int  = make([]int, len(*g))
		key []uint = make([]uint, len(*g))
		mst []bool = make([]bool, len(*g))
	)

	for i := range key {
		key[i] = INF
	}

	key[0] = 0
	par[0] = -1

	for i := 0; i < len(*g); i++ {
		mst[min(len(*g), &key, &mst)] = true

		for j := 0; j < len(*g); j++ {
			if (*g)[i][j] != 0 &&
				!mst[j] && (*g)[i][j] < key[j] {
				par[j], key[j] = i, (*g)[i][j]
			}
		}
	}

	return &par
}

func main() {
	example := [][]uint{
		{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0},
	}

	show(&example, nil)
	ans := prim(&example)
	show(&example, ans)
}
