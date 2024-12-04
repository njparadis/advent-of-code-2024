package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day4() {
	fmt.Println("Day 4")
	file, err := os.Open(fmt.Sprintf("%s/day_4.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var grid [][]string

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

	day4part1(grid)
	day4part2(grid)
}

func checkDirectionForChar(grid [][]string, i, j int, char string) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return false
	}

	if grid[i][j] == char {
		return true
	}

	return false
}

func printGridCandidates(grid [][]string, x, y int) {
	if x < 3 || y < 3 || x >= len(grid[0])-3 || y >= len(grid)-3 {
		return
	}

	center := grid[y][x]
	fmt.Println("XMAS candidate at", x, y)
	fmt.Println(grid[y-3][x-3], grid[y-3][x-2], grid[y-3][x-1], grid[y-3][x], grid[y-3][x+1], grid[y-3][x+2], grid[y-3][x+3])
	fmt.Println(grid[y-2][x-3], grid[y-2][x-2], grid[y-2][x-1], grid[y-2][x], grid[y-2][x+1], grid[y-2][x+2], grid[y-2][x+3])
	fmt.Println(grid[y-1][x-3], grid[y-1][x-2], grid[y-1][x-1], grid[y-1][x], grid[y-1][x+1], grid[y-1][x+2], grid[y-1][x+3])
	fmt.Println(grid[y][x-3], grid[y][x-2], grid[y][x-1], center, grid[y][x+1], grid[y][x+2], grid[y][x+3])
	fmt.Println(grid[y+1][x-3], grid[y+1][x-2], grid[y+1][x-1], grid[y+1][x], grid[y+1][x+1], grid[y+1][x+2], grid[y+1][x+3])
	fmt.Println(grid[y+2][x-3], grid[y+2][x-2], grid[y+2][x-1], grid[y+2][x], grid[y+2][x+1], grid[y+2][x+2], grid[y+2][x+3])
	fmt.Println(grid[y+3][x-3], grid[y+3][x-2], grid[y+3][x-1], grid[y+3][x], grid[y+3][x+1], grid[y+3][x+2], grid[y+3][x+3])
}

func day4part1(grid [][]string) {
	xmasCount := 0

	for i, row := range grid {
		for j, letter := range row {
			if letter == "X" {
				xmasfound := 0
				// check surrounding letters to see if we can make XMAS
				// up
				if checkDirectionForChar(grid, i-1, j, "M") &&
					checkDirectionForChar(grid, i-2, j, "A") &&
					checkDirectionForChar(grid, i-3, j, "S") {
					xmasfound++
				}
				// left
				if checkDirectionForChar(grid, i, j-1, "M") &&
					checkDirectionForChar(grid, i, j-2, "A") &&
					checkDirectionForChar(grid, i, j-3, "S") {
					xmasfound++
				}
				// down
				if checkDirectionForChar(grid, i+1, j, "M") &&
					checkDirectionForChar(grid, i+2, j, "A") &&
					checkDirectionForChar(grid, i+3, j, "S") {
					xmasfound++
				}
				// right
				if checkDirectionForChar(grid, i, j+1, "M") &&
					checkDirectionForChar(grid, i, j+2, "A") &&
					checkDirectionForChar(grid, i, j+3, "S") {
					xmasfound++
				}
				// diagonal right down
				if checkDirectionForChar(grid, i-1, j-1, "M") &&
					checkDirectionForChar(grid, i-2, j-2, "A") &&
					checkDirectionForChar(grid, i-3, j-3, "S") {
					xmasfound++
				}
				// diagonal left down
				if checkDirectionForChar(grid, i+1, j-1, "M") &&
					checkDirectionForChar(grid, i+2, j-2, "A") &&
					checkDirectionForChar(grid, i+3, j-3, "S") {
					xmasfound++
				}
				// diagonal right up
				if checkDirectionForChar(grid, i+1, j+1, "M") &&
					checkDirectionForChar(grid, i+2, j+2, "A") &&
					checkDirectionForChar(grid, i+3, j+3, "S") {
					xmasfound++
				}
				// diagonal left up
				if checkDirectionForChar(grid, i-1, j+1, "M") &&
					checkDirectionForChar(grid, i-2, j+2, "A") &&
					checkDirectionForChar(grid, i-3, j+3, "S") {
					xmasfound++
				}
				xmasCount += xmasfound

			}
		}
	}

	fmt.Println("XMAS count:", xmasCount)
}

func day4part2(grid [][]string) {
	xmasCount := 0

	for i, row := range grid {
		for j, letter := range row {
			if letter == "A" {
				xmasfound := 0
				// check surrounding letters to see if we can make XMAS
				// up left to down right
				if checkDirectionForChar(grid, i-1, j-1, "M") &&
					checkDirectionForChar(grid, i+1, j+1, "S") {
					xmasfound++
				}
				// down left to up right
				if checkDirectionForChar(grid, i+1, j+1, "M") &&
					checkDirectionForChar(grid, i-1, j-1, "S") {
					xmasfound++
				}
				// down right to up left
				if checkDirectionForChar(grid, i+1, j-1, "M") &&
					checkDirectionForChar(grid, i-1, j+1, "S") {
					xmasfound++
				}
				// up right to down left
				if checkDirectionForChar(grid, i-1, j+1, "M") &&
					checkDirectionForChar(grid, i+1, j-1, "S") {
					xmasfound++
				}
				if xmasfound == 2 {
					xmasCount++
				}

			}
		}
	}

	fmt.Println("X-MAS count:", xmasCount)
}
