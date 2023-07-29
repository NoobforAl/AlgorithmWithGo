package main

import (
	"log"
	"reflect"
)

/*
*
*
* Plus One
*
* args :
*   nums1 []int slice
*   m 		int
*   nums2 []int slice
*   n 		int
*
*
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Merge Sort :)
	var res []int
	var i, j int

	for i < m+n && j < n {
		if nums1[i] == 0 && i >= m {
			i++
			continue
		}

		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	for ; i < m+n; i++ {
		res = append(res, nums1[i])
	}

	for ; j < n; j++ {
		res = append(res, nums2[j])
	}

	for i := 0; i < n+m; i++ {
		nums1[i] = res[i]
	}
}

func main() {
	// simple test
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)

	if !reflect.DeepEqual(nums1, []int{1, 2, 2, 3, 5, 6}) {
		log.Fatalf("Worn Answer! answer equal `[1,2,2,3,5,6]` not equal : %v", nums1)
	}
	log.Println("Answer is True!")
}
