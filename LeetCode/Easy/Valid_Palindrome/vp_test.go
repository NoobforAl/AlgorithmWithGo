package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	testStr := "A man, a plan, a canal: Panama"
	assert.Equal(t, isPalindrome(testStr), true)
}

func TestCase2(t *testing.T) {
	testStr := "race a car"
	assert.Equal(t, isPalindrome(testStr), false)
}
