package main

import (
	"log"
	"reflect"
)

type TypeSort interface {
	int | int64 | int32 | int16 | int8 |
		float64 | float32 | byte
}

func merge[T TypeSort](first, second []T) []T {
	var res []T
	var i, j int

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			res = append(res, first[i])
			i++
		} else {
			res = append(res, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		res = append(res, first[i])
	}

	for ; j < len(second); j++ {
		res = append(res, second[j])
	}

	return res
}

func mergeSort[T TypeSort](items []T) []T {
	if len(items) < 2 {
		return items
	}

	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func main() {
	sl := []int{2, 5, 4, 6, 1, 8, 7, 3, 9}
	ans := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ansF := mergeSort(sl)

	if !reflect.DeepEqual(ans, ansF) {
		log.Fatalf("This answer not equal :\n%v\n%v", ans, ansF)
	}

	log.Println("True Answer! :", ansF)
}
