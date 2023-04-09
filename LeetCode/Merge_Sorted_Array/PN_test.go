package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{}
	merge(nums1, 1, nums2, 0)

	if !reflect.DeepEqual(nums1, []int{1}) {
		t.Fatalf("Worn Answer! answer equal '[1]' not equal : %v", nums1)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	merge(nums1, 1, nums2, 1)

	if !reflect.DeepEqual(nums1, []int{1, 2}) {
		t.Fatalf("Worn Answer! answer equal '[ 1, 2 ] not equal : %v", nums1)
	}
	t.Log("Answer is True!")
}
