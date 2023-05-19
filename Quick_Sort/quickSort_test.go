package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert.Equal(t, QuickSort([]int{4, 3, 7, 2, 6, 4, 7, 9, 1, 2, 3}),
		[]int{1, 2, 2, 3, 3, 4, 4, 6, 7, 7, 9})
}

func TestCase2(t *testing.T) {
	assert.Equal(t, QuickSort([]rune{'1', 'a', 'c', 'z', 't', 'q', 'z', 'x', 'y', 'd'}),
		[]rune{'1', 'a', 'c', 'd', 'q', 't', 'x', 'y', 'z', 'z'})
}
