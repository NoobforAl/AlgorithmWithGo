package main

import (
	"fmt"
)

func show(m *[][]uint) {
	for _, v := range *m {
		for _, va := range v {
			if va == INF {
				fmt.Print("INF", " ")
			} else {
				fmt.Print(va, " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
