package solutions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day3() {
	fmt.Println("Day 3")
	day3part1()
	day3part2()
}

func day3part1() {
	file, err := os.Open(fmt.Sprintf("%s/day_3.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	re := regexp.MustCompile(`mul\((-?\d{1,3}|-?9[0-8][0-9]|-?99[0-9]),(-?\d{1,3}|-?9[0-8][0-9]|-?99[0-9])\)`)

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			xStr := match[1]
			yStr := match[2]

			// Convert X and Y to integers
			x, err1 := strconv.Atoi(xStr)
			y, err2 := strconv.Atoi(yStr)
			if err1 != nil || err2 != nil {
				fmt.Printf("Error converting numbers: %v, %v\n", err1, err2)
				continue
			}

			// Compute the product and add to the total
			total += x * y
		}
	}
	fmt.Println("Total: ", total)
}

func day3part2() {
	file, err := os.Open(fmt.Sprintf("%s/day_3.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	mulRe := regexp.MustCompile(`mul\((-?\d{1,3}|-?9[0-8][0-9]|-?99[0-9]),(-?\d{1,3}|-?9[0-8][0-9]|-?99[0-9])\)`)
	doRe := regexp.MustCompile(`\bdo\(\)`)
	dontRe := regexp.MustCompile(`\bdon't\(\)`)

	tokenRe := regexp.MustCompile(`mul\((-?\d{1,3}|-?9[0-8][0-9]|-?99[0-9]),(-?\d{1,3}|-?9[0-8][0-9]|-?99[0-9])\)|\bdo\(\)|\bdon't\(\)`)

	total := 0     // Running total of products
	modeOn := true // Start in "on" mode

	// Scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Find all tokens in the line
		tokens := tokenRe.FindAllString(line, -1)

		// Process tokens sequentially
		for _, token := range tokens {
			// Handle `don't()`
			if dontRe.MatchString(token) {
				modeOn = false
				continue
			}

			// Handle `do()`
			if doRe.MatchString(token) {
				modeOn = true
				continue
			}

			// Handle `mul(X,Y)` if in "on" mode
			if modeOn && mulRe.MatchString(token) {
				matches := mulRe.FindStringSubmatch(token)
				xStr := matches[1]
				yStr := matches[2]

				// Convert X and Y to integers
				x, err1 := strconv.Atoi(xStr)
				y, err2 := strconv.Atoi(yStr)
				if err1 != nil || err2 != nil {
					fmt.Printf("Error converting numbers: %v, %v\n", err1, err2)
					continue
				}

				// Compute the product and add to the total
				total += x * y
			}
		}
	}

	fmt.Println("Total (with conditional statements): ", total)
}
