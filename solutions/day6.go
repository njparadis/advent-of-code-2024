package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func Day6() {
	fmt.Println("Day 6")
	file, err := os.Open(fmt.Sprintf("%s/day_6.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var grid [][]string
	const STARTING_Y = 86
	const STARTING_X = 45
	const (
		UP    = 0
		RIGHT = 1
		DOWN  = 2
		LEFT  = 3
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var letters []string

		for _, field := range fields {
			chars := strings.Split(field, "")
			letters = append(letters, chars...)
		}

		grid = append(grid, letters)
	}

	visited := findVisitedCells(grid, STARTING_X, STARTING_Y, UP)

	total := 0
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] {
				total++
			}
		}
	}

	loopCount := 0
	var loopCountMu sync.Mutex
	var wg sync.WaitGroup

	for y := range grid {
		for x := range grid[y] {
			if visited[y][x] && grid[y][x] == "." {
				wg.Add(1)

				go func(y, x int) {
					defer wg.Done()

					// Create a copy of the grid for this goroutine
					gridCopy := make([][]string, len(grid))
					for i := range grid {
						gridCopy[i] = append([]string(nil), grid[i]...)
					}

					// Modify the copy of the grid
					gridCopy[y][x] = "#" // Place a `#`
					if causesLoop(gridCopy, STARTING_X, STARTING_Y, UP) {
						loopCountMu.Lock()
						loopCount++
						loopCountMu.Unlock()
					}
				}(y, x)
			}
		}
	}

	wg.Wait()

	fmt.Println("Unique places visited: ", total)
	fmt.Println("Number of cells that cause a loop: ", loopCount)
}

func findVisitedCells(grid [][]string, startX, startY, startDir int) [][]bool {
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}

	x, y, direction := startX, startY, startDir
	dx := []int{0, 1, 0, -1} // UP, RIGHT, DOWN, LEFT
	dy := []int{-1, 0, 1, 0} // UP, RIGHT, DOWN, LEFT

	// Traverse the grid
	for {
		visited[y][x] = true

		// Compute the next position
		xNext, yNext := x+dx[direction], y+dy[direction]

		// Check for boundaries
		if xNext < 0 || xNext >= len(grid[0]) || yNext < 0 || yNext >= len(grid) {
			break
		}

		// Check for `#` and adjust direction if needed
		if grid[yNext][xNext] == "#" {
			direction = (direction + 1) % 4 // Turn 90 degrees right
			continue                        // Re-evaluate the next position
		}

		// Move to the next position
		x, y = xNext, yNext
	}

	return visited
}

func causesLoop(grid [][]string, startX, startY, startDir int) bool {
	visited := make(map[[3]int]bool)
	x, y, direction := startX, startY, startDir
	dx := []int{0, 1, 0, -1} // UP, RIGHT, DOWN, LEFT
	dy := []int{-1, 0, 1, 0} // UP, RIGHT, DOWN, LEFT

	for {
		state := [3]int{x, y, direction}
		if visited[state] {
			return true // Loop detected
		}
		visited[state] = true

		// Compute next position
		xNext, yNext := x+dx[direction], y+dy[direction]

		// Check for boundaries
		if xNext < 0 || xNext >= len(grid[0]) || yNext < 0 || yNext >= len(grid) {
			return false // Traversal terminated
		}

		// Redirect on `#`
		if grid[yNext][xNext] == "#" {
			direction = (direction + 1) % 4 // Turn 90 degrees right
		} else {
			x, y = xNext, yNext // Move forward
		}
	}
}
