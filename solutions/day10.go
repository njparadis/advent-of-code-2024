package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day10() {
	fmt.Println("Day 10")
	file, err := os.Open(fmt.Sprintf("%s/day_10.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var grid [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var numbers []int

		chars := strings.Split(line, "")
		for _, c := range chars {
			num, err := strconv.Atoi(c)
			if err != nil {
				fmt.Println(err)
				continue
			}
			numbers = append(numbers, num)
		}

		grid = append(grid, numbers)
	}

	day10part1(grid)
	day10part2(grid)
}

func day10part1(grid [][]int) {
	score := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == 0 {
				visited := make([][]bool, len(grid))
				for i := range visited {
					visited[i] = make([]bool, len(grid[0]))
				}
				score += scoreTrail(grid, x, y, visited)
			}
		}
	}

	fmt.Println("Total score:", score)
}

func scoreTrail(grid [][]int, x, y int, visited [][]bool) int {
	if visited[y][x] {
		return 0
	}

	visited[y][x] = true

	if grid[y][x] == 9 {
		return 1
	}

	score := 0
	if x-1 >= 0 && grid[y][x-1] == grid[y][x]+1 {
		score += scoreTrail(grid, x-1, y, visited)
	}
	if x+1 < len(grid[y]) && grid[y][x+1] == grid[y][x]+1 {
		score += scoreTrail(grid, x+1, y, visited)
	}
	if y-1 >= 0 && grid[y-1][x] == grid[y][x]+1 {
		score += scoreTrail(grid, x, y-1, visited)
	}
	if y+1 < len(grid) && grid[y+1][x] == grid[y][x]+1 {
		score += scoreTrail(grid, x, y+1, visited)
	}

	visited[y][x] = false

	return score
}

func day10part2(grid [][]int) {
	totalTrails := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == 0 {
				totalTrails += countTrails(grid, x, y)
			}
		}
	}

	fmt.Println("Total distinct trails:", totalTrails)
}

func countTrails(grid [][]int, x, y int) int {
	if grid[y][x] == 9 {
		return 1
	}

	count := 0
	if x-1 >= 0 && grid[y][x-1] == grid[y][x]+1 {
		count += countTrails(grid, x-1, y)
	}
	if x+1 < len(grid[y]) && grid[y][x+1] == grid[y][x]+1 {
		count += countTrails(grid, x+1, y)
	}
	if y-1 >= 0 && grid[y-1][x] == grid[y][x]+1 {
		count += countTrails(grid, x, y-1)
	}
	if y+1 < len(grid) && grid[y+1][x] == grid[y][x]+1 {
		count += countTrails(grid, x, y+1)
	}

	return count
}
