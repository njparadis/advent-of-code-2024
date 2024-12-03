package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	fmt.Println("Day 2")
	file, err := os.Open(fmt.Sprintf("%s/day_2.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var reports [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var levels []int

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			levels = append(levels, num)
		}

		reports = append(reports, levels)
	}
	day2part1(reports)
	day2part2(reports)
}

func isReportSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	prev := report[0]
	direction := 0

	for _, level := range report[1:] {
		diff := level - prev

		if diff < 1 && diff > -1 || diff > 3 || diff < -3 {
			return false
		}

		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else if diff < 0 {
				direction = -1
			}
		}

		if (direction == 1 && diff < 0) || (direction == -1 && diff > 0) {
			return false
		}

		prev = level
	}
	return true
}

func day2part1(reports [][]int) {
	numSafe := 0
	for _, report := range reports {
		if isReportSafe(report) {
			numSafe++
		}
	}
	fmt.Println("Number of safe reports: ", numSafe)
}

func day2part2(reports [][]int) {
	numSafe := 0
	for _, report := range reports {
		anySafe := false
		for i := 0; i < len(report); i++ {
			newSlice := make([]int, len(report)-1)
			copy(newSlice, report[:i])
			copy(newSlice[i:], report[i+1:])
			anySafe = anySafe || isReportSafe(newSlice)
		}
		if anySafe {
			numSafe++
		}
	}
	fmt.Println("Number of safe reports (Problem Dampener): ", numSafe)
}
