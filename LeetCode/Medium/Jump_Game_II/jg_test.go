package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert.Equal(t, jump([]int{2, 3, 1, 1, 4}), 2, "Worn Answer! answer equal '2'")
}

func TestCase2(t *testing.T) {
	assert.Equal(t, jump([]int{2, 3, 0, 1, 4}), 2, "Worn Answer! answer equal '2'")
}
