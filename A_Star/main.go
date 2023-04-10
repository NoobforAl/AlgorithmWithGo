package main

import (
	"log"
	"reflect"

	pr "github.com/kr/pretty"
)

type Map [][]int
type Point [2]int

type Node struct {
	Parent    *Node
	Location  Point
	Reachable bool
	H, G, F   int
}

var newPosition = [8]Point{
	{0, -1}, {0, 1},
	{-1, 0}, {1, 0},
	{-1, -1}, {-1, 1},
	{1, -1}, {1, 1},
}

func (n *Node) equalTwoNode(node *Node) bool {
	return reflect.DeepEqual(n.Location, node.Location)
}

func aStar(m Map, start, end *Node) []Point {
	var OpenNode []Node = []Node{*start}
	var CloseNode []Node

	for len(OpenNode) > 0 {
		currentNode := OpenNode[0]
		currentIndex := 0

		for i, v := range OpenNode {
			if v.F > currentNode.F {
				currentNode = v
				currentIndex = i
			}
		}

		if currentNode.equalTwoNode(end) {
			var path []Point
			var current *Node = &currentNode
			for current != nil {
				path = append(path, current.Location)
				current = current.Parent
			}
			return path
		}

		// remove open node with Index and add to close Node
		OpenNode = append(OpenNode[:currentIndex], OpenNode[currentIndex+1:]...)
		CloseNode = append(CloseNode, currentNode)

		var children []Node
		for _, v := range newPosition {
			nodePosition := Point{currentNode.Location[0] + v[0], currentNode.Location[1] + v[1]}

			if nodePosition[0] > len(m)-1 ||
				nodePosition[0] < 0 ||
				nodePosition[1] > len(m[len(m)-1])-1 ||
				nodePosition[1] < 0 {
				continue
			}

			if m[nodePosition[0]][nodePosition[1]] != 0 {
				continue
			}

			newNode := Node{Parent: &currentNode, Location: nodePosition}
			children = append(children, newNode)
		}

		for _, child := range children {
			var is_close bool
			for _, c := range CloseNode {
				if child.equalTwoNode(&c) {
					is_close = true
					break
				}
			}

			if is_close {
				continue
			}

			child.G = currentNode.G + 1
			child.H = (child.Location[0] - end.Location[0]*end.Location[0]) +
				(child.Location[1] - end.Location[1]*end.Location[1])
			child.F = child.G + child.H

			OpenNode = append(OpenNode, child)
		}

	}

	return []Point{}
}

func main() {
	m := Map{
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	start := Node{Location: Point{0, 0}}
	end := Node{Location: Point{7, 6}}

	out := aStar(m, &start, &end)
	ans := []Point{
		{7, 6}, {6, 5},
		{5, 4}, {4, 3},
		{3, 3}, {2, 2},
		{1, 1}, {0, 0}}

	if !reflect.DeepEqual(ans, out) {
		log.Fatalf("This Answer Not Equal:\n %# v \nYou Answer:\n %# v",
			pr.Formatter(out),
			pr.Formatter(ans))
	}
	log.Println("True Answer!")
}
