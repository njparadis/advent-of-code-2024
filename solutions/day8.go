package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day8() {
	fmt.Println("Day 8")
	file, err := os.Open(fmt.Sprintf("%s/day_8.txt", INPUT_FOLDER))
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

	antennas := make(map[string][][2]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != "." {
				antennas[grid[i][j]] = append(antennas[grid[i][j]], [2]int{i, j})
			}
		}
	}

	day8part1(antennas)
	day8part2(antennas)
}

func day8part1(antennas map[string][][2]int) {
	antinodeLocations := make([][]bool, 50)
	for i := 0; i < 50; i++ {
		antinodeLocations[i] = make([]bool, 50)
	}
	for _, locations := range antennas {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				p1 := locations[i]
				p2 := locations[j]

				dx := p2[0] - p1[0]
				dy := p2[1] - p1[1]

				// Calculate the antinodes
				antinode1X := p1[0] - dx
				antinode1Y := p1[1] - dy
				antinode2X := p2[0] + dx
				antinode2Y := p2[1] + dy

				if antinode1X >= 0 && antinode1X < 50 && antinode1Y >= 0 && antinode1Y < 50 {
					antinodeLocations[antinode1X][antinode1Y] = true
				}

				if antinode2X >= 0 && antinode2X < 50 && antinode2Y >= 0 && antinode2Y < 50 {
					antinodeLocations[antinode2X][antinode2Y] = true
				}
			}
		}
	}

	totalUnique := 0
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			if antinodeLocations[i][j] {
				totalUnique++
			}
		}
	}

	fmt.Println("Total unique antinode locations:", totalUnique)
}

func day8part2(antennas map[string][][2]int) {
	antinodeLocations := make([][]bool, 50)
	for i := 0; i < 50; i++ {
		antinodeLocations[i] = make([]bool, 50)
	}
	for _, locations := range antennas {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				p1 := locations[i]
				p2 := locations[j]

				dx := p2[0] - p1[0]
				dy := p2[1] - p1[1]

				// Calculate the antinodes
				x, y := p1[0], p1[1]
				for {
					// Calculate the next antinode in the forward direction
					x += dx
					y += dy

					// Check if the point is outside the grid
					if x < 0 || x >= 50 || y < 0 || y >= 50 {
						break
					}
					antinodeLocations[x][y] = true
				}

				// Backward direction
				x, y = p2[0], p2[1]
				for {
					// Calculate the next antinode in the backward direction
					x -= dx
					y -= dy

					// Check if the point is outside the grid
					if x < 0 || x >= 50 || y < 0 || y >= 50 {
						break
					}
					antinodeLocations[x][y] = true
				}
			}
		}
	}

	totalUnique := 0
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			if antinodeLocations[i][j] {
				totalUnique++
			}
		}
	}

	fmt.Println("Total unique antinode locations (with harmonics):", totalUnique)
}
