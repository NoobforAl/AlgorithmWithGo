package main

import "fmt"

func show(g *[][]uint, par *[]int) {
	if par == nil {
		for _, v := range *g {
			for _, va := range v {
				if va == INF {
					fmt.Print("INF", " ")
				} else {
					fmt.Print(va, " ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println()
		fmt.Println("Edge \tWeight")
		for i := 1; i < len(*g); i++ {
			fmt.Printf("%d - %d \t%d \n",
				(*par)[i], i,
				(*g)[i][(*par)[i]])
		}
	}
	fmt.Println()
}
