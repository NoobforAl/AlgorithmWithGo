package main

import (
	"bytes"
	"log"
	"math/rand"

	"github.com/kr/pretty"
)

type Num interface {
	int64 | int32 | int16 | int8 | int |
		float64 | float32
}

func QuickSort[T Num](nums []T) []T {
	if len(nums) < 2 {
		return nums
	}

	p := rand.Intn(len(nums))
	less, greater := []T{}, []T{}

	for i := range nums {
		if i == p {
			continue
		}
		if nums[p] >= nums[i] {
			less = append(less, nums[i])
		}
	}

	for i := range nums {
		if i == p {
			continue
		}
		if nums[p] < nums[i] {
			greater = append(greater, nums[i])
		}
	}

	tmp := append(QuickSort(less), nums[p])
	return append(tmp, QuickSort(greater)...)
}

func main() {
	nums := []float32{1.8, 1.2, 1.3, 1.0, 1.4}

	var buf bytes.Buffer

	pretty.Fdiff(&buf, QuickSort(nums), []float32{1.0, 1.2, 1.3, 1.4, 1.8})
	if buf.String() != "" {
		log.Fatalf("Wrong Answer!, %s", buf.String())
	}
	log.Println("All OK!")
}
