package main

import "fmt"

func main() {

	grid := make([][]string, 5)
	for x := 0; x < len(grid); x++ {
		grid[x] = make([]string, 1)
	}

	fmt.Println(grid)

}
