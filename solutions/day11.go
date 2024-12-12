package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day11() {
	fmt.Println("Day 11")
	file, err := os.Open(fmt.Sprintf("%s/day_11.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var stones []int

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, " ")

		for _, c := range chars {
			num, err := strconv.Atoi(c)
			if err != nil {
				fmt.Println(err)
				continue
			}
			stones = append(stones, num)
		}
	}

	fmt.Println("Number of stones after 25 blinks:", blink(stones, 25))

	fmt.Println("Number of stones after 75 blinks:", blink(stones, 75))
}

func blink(stones []int, blinks int) int {
	memo := make(map[string]int)

	var count func(int, int) int
	count = func(stone, remainingBlinks int) int {
		if remainingBlinks == 0 {
			return 1
		}

		// Memoize results based on stone and remaining blinks
		key := fmt.Sprintf("%d-%d", stone, remainingBlinks)
		if val, ok := memo[key]; ok {
			return val
		}

		var result int
		if stone == 0 {
			result = count(1, remainingBlinks-1)
		} else {
			numDigits := len(strconv.Itoa(stone))
			if numDigits%2 == 0 && stone >= 10 {
				half := numDigits / 2
				divisor := int(math.Pow10(half))
				i1, i2 := stone/divisor, stone%divisor
				result = count(i1, remainingBlinks-1) + count(i2, remainingBlinks-1)
			} else {
				result = count(stone*2024, remainingBlinks-1)
			}
		}

		memo[key] = result
		return result
	}

	totalCount := 0
	for _, stone := range stones {
		totalCount += count(stone, blinks)
	}
	return totalCount
}
