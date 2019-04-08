package main

import "fmt"

var NUM_PEGS = 3

func towerOfHanoi(numRings int) {
	//Initialize pegs and fill first peg with rings
	pegs := [][]int{}
	for i := 0; i < NUM_PEGS; i++ {
		pegs = append(pegs, []int{})
	}
	for i := 0; i < numRings; i++ {
		pegs[0] = append(pegs[0], i+1)
	}
	fmt.Println("Initial State =", pegs)

	helper(numRings, pegs, 0, 1, 2)

}

func helper(numRingsToMove int, pegs [][]int, fromPeg, toPeg, usePeg int) {
	if numRingsToMove > 0 {

		// Move top n-1 peg from "fromPeg" to "usePeg" using "toPeg" as helper
		helper(numRingsToMove-1, pegs, fromPeg, usePeg, toPeg)

		// Move the last remaining peg from "fromPeg" to "toPeg"
		ring := pegs[fromPeg][0]
		pegs[fromPeg] = pegs[fromPeg][1:]
		pegs[toPeg] = append([]int{ring}, pegs[toPeg]...)
		fmt.Printf("Move ring %d from peg %d to peg %d\n", ring, fromPeg, toPeg)
		fmt.Println("State =", pegs)

		// Move the previously moved n-1 pegs from "usePeg" to "toPeg" using "fromPeg" as helper
		helper(numRingsToMove-1, pegs, usePeg, toPeg, fromPeg)
	}
}

func main() {
	towerOfHanoi(3)
}
