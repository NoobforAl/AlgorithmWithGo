package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	if ans := plusOne([]int{4, 3, 2, 1}); !reflect.DeepEqual(ans, []int{4, 3, 2, 2}) {
		t.Fatalf("Worn Answer! answer equal '[4,3,2,2]' not equal : %v", ans)
	}
	t.Log("Answer is True!")
}

func TestCase2(t *testing.T) {
	if ans := plusOne([]int{9, 9}); !reflect.DeepEqual(ans, []int{1, 0, 0}) {
		t.Fatalf("Worn Answer! answer equal '[ 1, 0, 0 ] not equal : %v", ans)
	}
	t.Log("Answer is True!")
}
