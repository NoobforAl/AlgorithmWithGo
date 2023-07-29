package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	testArr := []int{3, 2, 3}
	assert.Equal(t, majorityElement(testArr), 3)
}

func TestCase2(t *testing.T) {
	testArr := []int{2, 2, 1, 1, 1, 2, 2}
	assert.Equal(t, majorityElement(testArr), 2)
}
