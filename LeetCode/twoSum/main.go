package main

import "fmt"

/*
*
*
* this func for find 2 value in nums for
* sum of tow value equal target
*
* args :
*   nums slices []int
*   target int
*
* return value:
*   slices []int ( always return with len 2 )
*
*
 */
func twoSum(nums []int, target int) []int {

	// make one map < check > is a check list
	// key is num of nums
	// value is one index of nums
	check := make(map[int]int)

	for i, v := range nums {

		// check <target - v> in map or not
		/*
			/ target = v + x => x = target - v
			/ different = target - v
			/ check different in check map or not
		*/
		if vv, ok := check[target-v]; ok {
			// if different in check return ( past index, now index )
			return []int{vv, i}
		}

		// if not found with different key ( target - v )
		// add new key value : num , index
		check[v] = i
	}

	// if not found tow int, return nil value
	return nil
}

func main() {
	// very simple test
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
