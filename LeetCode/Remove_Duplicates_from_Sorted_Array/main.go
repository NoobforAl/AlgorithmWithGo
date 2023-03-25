package main

import (
	"fmt"
	"reflect"
)
import "C"

/*
//* this answer with CPP
//* for golang we have a problem
//* in leetCode
#include <iostream>
#include <map>
#include <vector>

using namespace std;

int removeDuplicates(vector<int> &nums)

	{
	    int sum{0};
	    map<int, bool> checkList;

	    for (int i = 0; i < nums.size(); i++)
	    {
	        if (checkList.find(nums[i]) == checkList.end())
	        {
	            sum++;
	            checkList[nums[i]] = 0;
	        }
	    }

	    nums.clear();
	    for (auto x : checkList)
	        nums.push_back(x.first);

	    return sum;
	}
*/
func removeDuplicates(nums *[]int) int {
	// set a counter sum
	var sum int = 0

	// make a check list
	checkList := make(map[int]bool)

	// check all nums in check list
	// if notfound in check list
	// added
	for _, n := range *nums {
		if _, ok := checkList[n]; !ok {
			sum++
			checkList[n] = false
		}
	}

	// clear nums and
	// append all key in
	// check list!
	*nums = []int{}
	for k := range checkList {
		*nums = append(*nums, k)
	}

	return sum
}

func main() {
	// very simple test
	nums := []int{1, 1, 1}
	fmt.Println(removeDuplicates(&nums) == 1)
	fmt.Println(reflect.DeepEqual(nums, []int{1}))
}
