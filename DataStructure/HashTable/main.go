package main

import (
	"fmt"
)

const SIZE = 15

type Node struct {
	Val  int
	Next *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func HashFunction(i, size int) int {
	return (i % size)
}

func Insert(hash *HashTable, val int) int {
	index := HashFunction(val, hash.Size)
	element := Node{Val: val, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func Traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			fmt.Print(k, ": ")
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Val)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func Lookup(hash *HashTable, val int) bool {
	index := HashFunction(val, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Val == val {
				return true
			}
			t = t.Next
		}
	}

	return false
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}

	fmt.Println("Number of space: ", hash.Size)
	for i := 0; i < 120; i++ {
		Insert(hash, i)
	}

	Traverse(hash)
	fmt.Println(Lookup(hash, 1000))
}
