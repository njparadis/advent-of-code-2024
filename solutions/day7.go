package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day7() {
	fmt.Println("Day 7")
	file, err := os.Open(fmt.Sprintf("%s/day_7.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	totalValid := 0
	totalValidWithConcat := 0

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var numbers []int

		parts := strings.SplitN(fields[0], ":", 2)
		goal, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Couldn't parse goal:", err)
			continue
		}

		for _, field := range fields[1:] {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println(err)
				continue
			}
			numbers = append(numbers, num)
		}

		if calibrateInputs(goal, numbers[0], numbers[1:]) {
			totalValid += goal
		}

		if calibrateInputsWithConcat(goal, numbers[0], numbers[1:]) {
			totalValidWithConcat += goal
		}

	}

	fmt.Println("Total valid calibrations:", totalValid)
	fmt.Println("Total valid calibrations (with concatenation):", totalValidWithConcat)
}

func calibrateInputs(target, total int, input []int) bool {
	if len(input) == 0 {
		if total == target {
			return true
		}
		return false
	}
	return calibrateInputs(
		target, total+input[0], input[1:],
	) || calibrateInputs(
		target, total*input[0], input[1:],
	)
}

func calibrateInputsWithConcat(target, total int, input []int) bool {
	if len(input) == 0 {
		if total == target {
			return true
		}
		return false
	}
	add := calibrateInputsWithConcat(target, total+input[0], input[1:])
	mul := calibrateInputsWithConcat(target, total*input[0], input[1:])

	concatNum, _ := strconv.Atoi(fmt.Sprintf("%d%d", total, input[0]))
	concat := calibrateInputsWithConcat(target, concatNum, input[1:])

	return add || mul || concat
}
