package solutions

import (
	"bufio"
	"fmt"
	"os"
)

func Day12() {
	fmt.Println("Day 12")
	file, err := os.Open(fmt.Sprintf("%s/day_12.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var grid [][]rune

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		chars := []rune(line)
		grid = append(grid, chars)
	}

	grid = [][]rune{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'X', 'X', 'X', 'X'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'X', 'X', 'X', 'X'},
		{'E', 'E', 'E', 'E', 'E'},
	}

	day12part1(grid)
	day12part2(grid)
}

type Point struct {
	x, y int
}

var directions = []Point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func day12part1(grid [][]rune) {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	inBounds := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	var dfs func(x, y int, char rune) (int, int)
	dfs = func(x, y int, char rune) (int, int) {
		stack := []Point{{x, y}}
		visited[x][y] = true
		area := 0
		perimeter := 0

		for len(stack) > 0 {
			cell := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area++

			// Check all 4 neighbors
			for _, dir := range directions {
				nx, ny := cell.x+dir.x, cell.y+dir.y

				if !inBounds(nx, ny) || grid[nx][ny] != char {
					perimeter++
				} else if !visited[nx][ny] {
					visited[nx][ny] = true
					stack = append(stack, Point{nx, ny})
				}
			}
		}

		return area, perimeter
	}

	cost := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				char := grid[i][j]
				area, perimeter := dfs(i, j, char)
				cost += area * perimeter
			}
		}
	}
	fmt.Println("Total cost:", cost)
}

var allDirections = []Point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

var cardinalDirs = []Point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

var diagonalDirs = []Point{
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

func day12part2(grid [][]rune) {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	inBounds := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	var dfs func(x, y int, char rune) (int, int)
	dfs = func(x, y int, char rune) (int, int) {
		stack := []Point{{x, y}}
		visited[x][y] = true
		area := 0

		// Use a map to track unique corners
		corners := make(map[Point]bool)

		for len(stack) > 0 {
			cell := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area++

			// Check diagonal neighbors for corners
			for _, diag := range diagonalDirs {
				nx, ny := cell.x+diag.x, cell.y+diag.y

				fmt.Printf("Cell (%x, %x): %c\n", cell.x, cell.y, grid[cell.x][cell.y])

				// Check if the diagonal neighbor is outside the region
				if inBounds(nx, ny) && grid[nx][ny] == char {
					continue // Diagonal neighbor is part of the same region
				}

				if inBounds(nx, ny) {
					fmt.Printf("Diag %c\n", grid[nx][ny])
				}

				// Check the two cardinal neighbors adjacent to the diagonal
				ax, ay := cell.x+diag.x, cell.y // Adjacent cell 1 (cardinal)
				bx, by := cell.x, cell.y+diag.y // Adjacent cell 2 (cardinal)

				if inBounds(ax, ay) {
					fmt.Printf("Adjacent A: %c\n", grid[ax][ay])
				}

				if inBounds(bx, by) {
					fmt.Printf("Adjacent B: %c\n", grid[bx][by])
				}

				adjacentA := inBounds(ax, ay) && grid[ax][ay] == char
				adjacentB := inBounds(bx, by) && grid[bx][by] == char

				if (adjacentA && adjacentB) || (!adjacentA && !adjacentB) {
					corners[Point{cell.x, cell.y}] = true
				}
			}

			// Explore cardinal neighbors
			for _, dir := range cardinalDirs {
				nx, ny := cell.x+dir.x, cell.y+dir.y
				if inBounds(nx, ny) && grid[nx][ny] == char && !visited[nx][ny] {
					visited[nx][ny] = true
					stack = append(stack, Point{nx, ny})
				}
			}
		}

		fmt.Println("Area", area)
		fmt.Println("Corners", len(corners))

		return area, len(corners)
	}

	cost := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				char := grid[i][j]
				area, numCorners := dfs(i, j, char)
				cost += area * numCorners
			}
		}
	}
	fmt.Println("Total cost (with bulk discount):", cost)
}
