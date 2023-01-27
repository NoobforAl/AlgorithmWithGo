package main

import "fmt"

func show(m *[][]int8) {
	keys := [2]string{"*", "Q"}
	for _, v := range *m {
		for _, va := range v {
			fmt.Print(keys[va], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
