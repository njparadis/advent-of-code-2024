package solutions

import (
	"bufio"
	"fmt"
	"os"
)

type Robot struct {
	x, y   int
	vX, vY int
}

func Day14() {
	fmt.Println("Day 14")
	file, err := os.Open(fmt.Sprintf("%s/day_14.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	const (
		gridWidth  = 101
		gridHeight = 103
	)

	scanner := bufio.NewScanner(file)

	robots := make([]Robot, 0)

	for scanner.Scan() {
		var pX, pY, vX, vY int
		line := scanner.Text()

		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pX, &pY, &vX, &vY)
		if err != nil {
			fmt.Println("Failed to parse line:", line)
			continue
		}
		robots = append(robots, Robot{pX, pY, vX, vY})
	}

	day14part1(gridWidth, gridHeight, robots)
	day14part2(gridWidth, gridHeight, robots)
}

func findRobotPosition(gridWidth, gridHeight, pX, pY, vX, vY, seconds int) (int, int) {
	nX := ((pX+vX*seconds)%gridWidth + gridWidth) % gridWidth
	nY := ((pY+vY*seconds)%gridHeight + gridHeight) % gridHeight
	return nX, nY
}

func day14part1(gridWidth, gridHeight int, robots []Robot) {
	grid := make([][]int, gridHeight)
	for i := range grid {
		grid[i] = make([]int, gridWidth)
	}

	for _, r := range robots {
		nX, nY := findRobotPosition(gridWidth, gridHeight, r.x, r.y, r.vX, r.vY, 100)
		grid[nY][nX]++
	}

	midX, midY := gridWidth/2, gridHeight/2
	q1, q2, q3, q4 := 0, 0, 0, 0

	for row := 0; row < gridHeight; row++ {
		for col := 0; col < gridWidth; col++ {
			if row == midY || col == midX {
				continue
			}

			if row < midY && col < midX {
				q1 += grid[row][col]
			} else if row < midY && col > midX {
				q2 += grid[row][col]
			} else if row > midY && col < midX {
				q3 += grid[row][col]
			} else if row > midY && col > midX {
				q4 += grid[row][col]
			}
		}
	}
	fmt.Println("Safety factor:", q1*q2*q3*q4)
}

func day14part2(gridWidth, gridHeight int, robots []Robot) {
	seconds := 0
	for seconds < 10000 {
		grid := make([][]bool, gridHeight)
		for i := range grid {
			grid[i] = make([]bool, gridWidth)
		}
		for _, r := range robots {
			nX, nY := findRobotPosition(gridWidth, gridHeight, r.x, r.y, r.vX, r.vY, seconds)
			grid[nY][nX] = true
		}

		maxConsecutive := 0
		for row := 0; row < gridHeight; row++ {
			numConsecutive := 0
			for col := 0; col < gridWidth; col++ {
				if grid[row][col] {
					numConsecutive++
				} else {
					if numConsecutive > maxConsecutive {
						maxConsecutive = numConsecutive
					}
					numConsecutive = 0
				}

			}
		}
		// assuming we've found the line denoting the tree's bottom
		if maxConsecutive >= 10 {
			fmt.Println("Seconds to tree:", seconds)
			for row := 0; row < gridHeight; row++ {
				for col := 0; col < gridWidth; col++ {
					if grid[row][col] {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
			break
		}
		seconds++
	}
}
