package main

import (
	"AOC2023/utils"
	"fmt"
	"strings"
	"time"
)

/*
Converts the line of instructions into an int array of instructions, where
L is 0 and R is 1, so that the traversal is easier to write.
*/
func instructions(firstLine string) []int {
	var instructions = make([]int, 0)
	for _, r := range firstLine {
		if r == 'L' {
			instructions = append(instructions, 0)
		} else {
			instructions = append(instructions, 1)
		}
	}
	return instructions
}

/*
Converts the network lines to a network map where
the keys are the string representation of the nodes
and the values are string arrays where the next left and right nodes are stored,
in this order.
*/
func getNetwork(mapLines []string) map[string][]string {
	var res = make(map[string][]string)
	for _, line := range mapLines {
		var components = strings.Split(line, " = ")
		var directions = strings.Split(components[1], ", ")
		res[strings.TrimSpace(components[0])] = []string{directions[0][1:], directions[1][:3]}
	}
	return res
}

/*
Operates a simple traversal of the network as defined by part one.
Returns the number of steps required to reach the end node.
*/
func simpleTraversal(instructions []int, network map[string][]string) int {
	//Instead of keeping track of all the steps, the number of full traversals are tracked.
	// i from the for loop will keep track of the number of steps within a single traversal.
	var currentNode = "AAA"
	var fullTraversals = 0
	for currentNode != "ZZZ" {
		for i, direction := range instructions {
			if currentNode == "ZZZ" {
				//End node reached within a traversal
				return i + len(instructions)*fullTraversals
			} else {
				currentNode = network[currentNode][direction]
			}
		}
		//A full traversal was done
		fullTraversals += 1
	}
	//End node reached at the end of a traversal
	return len(instructions) * fullTraversals
}

func partOne(input string) int {
	var fullLines = utils.LineByLine(input)
	return simpleTraversal(instructions(fullLines[0]), getNetwork(fullLines[2:]))
}

/*
From a network map, gets the starting nodes, as defined by part two.
*/
func getStartingNodes(network map[string][]string) []string {
	var startingNodes = make([]string, 0)
	for node := range network {
		if node[2] == 'A' {
			startingNodes = append(startingNodes, node)
		}
	}
	return startingNodes
}

/*
Traversal of the network as defined by part two, from a given starting node.
Same process as a simple traversal, where the definition on the end node has changed.
*/
func traversal(startingNode string, instructions []int, network map[string][]string) int {
	var currentNode = startingNode
	var fullTraversals = 0
	for currentNode[2] != 'Z' {
		for i, direction := range instructions {
			if currentNode[2] == 'Z' {
				return i + len(instructions)*fullTraversals
			} else {
				currentNode = network[currentNode][direction]
			}
		}
		fullTraversals += 1
	}
	return len(instructions) * fullTraversals
}

/*
Computes the answer for part two.
Computes the number of steps for each starting node, then uses LCM to determine the answer.
*/
func partTwo(input string) int {
	var fullLines = utils.LineByLine(input)
	var instructions = instructions(fullLines[0])
	var network = getNetwork(fullLines[2:])
	var startingNodes = getStartingNodes(network)
	var traversalResults = make([]int, 0)
	for _, node := range startingNodes {
		traversalResults = append(traversalResults, traversal(node, instructions, network))
	}
	return utils.PPCM(traversalResults)
}

func main() {
	var start = time.Now()
	fmt.Println(partOne("08/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("08/input.txt"))
	fmt.Println(time.Since(start))
}
