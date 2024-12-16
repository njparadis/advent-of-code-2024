package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/njparadis/advent-of-code-2024/solutions"
)

func main() {
	day := flag.String("day", "", "Day to run")
	flag.Parse()

	switch string(*day) {
	case "1":
		start := time.Now()
		solutions.Day1()
		duration := time.Since(start)
		fmt.Printf("Day 1 runtime: %v\n\n", duration)
	case "2":
		start := time.Now()
		solutions.Day2()
		duration := time.Since(start)
		fmt.Printf("Day 2 runtime: %v\n\n", duration)
	case "3":
		start := time.Now()
		solutions.Day3()
		duration := time.Since(start)
		fmt.Printf("Day 3 runtime: %v\n\n", duration)
	case "4":
		start := time.Now()
		solutions.Day4()
		duration := time.Since(start)
		fmt.Printf("Day 4 runtime: %v\n\n", duration)
	case "5":
		start := time.Now()
		solutions.Day5()
		duration := time.Since(start)
		fmt.Printf("Day 5 runtime: %v\n\n", duration)
	case "6":
		start := time.Now()
		solutions.Day6()
		duration := time.Since(start)
		fmt.Printf("Day 6 runtime: %v\n\n", duration)
	case "7":
		start := time.Now()
		solutions.Day7()
		duration := time.Since(start)
		fmt.Printf("Day 7 runtime: %v\n\n", duration)
	case "8":
		start := time.Now()
		solutions.Day8()
		duration := time.Since(start)
		fmt.Printf("Day 8 runtime: %v\n\n", duration)
	case "9":
		start := time.Now()
		solutions.Day9()
		duration := time.Since(start)
		fmt.Printf("Day 9 runtime: %v\n\n", duration)
	case "10":
		start := time.Now()
		solutions.Day10()
		duration := time.Since(start)
		fmt.Printf("Day 10 runtime: %v\n\n", duration)
	case "11":
		start := time.Now()
		solutions.Day11()
		duration := time.Since(start)
		fmt.Printf("Day 11 runtime: %v\n\n", duration)
	default:
		start := time.Now()

		start1 := time.Now()
		solutions.Day1()
		duration1 := time.Since(start1)
		fmt.Printf("Day 1 runtime: %v\n\n", duration1)

		start2 := time.Now()
		solutions.Day2()
		duration2 := time.Since(start2)
		fmt.Printf("Day 2 runtime: %v\n\n", duration2)

		start3 := time.Now()
		solutions.Day3()
		duration3 := time.Since(start3)
		fmt.Printf("Day 3 runtime: %v\n\n", duration3)

		start4 := time.Now()
		solutions.Day4()
		duration4 := time.Since(start4)
		fmt.Printf("Day 4 runtime: %v\n\n", duration4)

		start5 := time.Now()
		solutions.Day5()
		duration5 := time.Since(start5)
		fmt.Printf("Day 5 runtime: %v\n\n", duration5)

		start6 := time.Now()
		solutions.Day6()
		duration6 := time.Since(start6)
		fmt.Printf("Day 6 runtime: %v\n\n", duration6)

		start7 := time.Now()
		solutions.Day7()
		duration7 := time.Since(start7)
		fmt.Printf("Day 7 runtime: %v\n\n", duration7)

		start8 := time.Now()
		solutions.Day8()
		duration8 := time.Since(start8)
		fmt.Printf("Day 8 runtime: %v\n\n", duration8)

		start9 := time.Now()
		solutions.Day9()
		duration9 := time.Since(start9)
		fmt.Printf("Day 9 runtime: %v\n\n", duration9)

		start10 := time.Now()
		solutions.Day10()
		duration10 := time.Since(start10)
		fmt.Printf("Day 10 runtime: %v\n\n", duration10)

		start11 := time.Now()
		solutions.Day11()
		duration11 := time.Since(start11)
		fmt.Printf("Day 11 runtime: %v\n\n", duration11)

		duration := time.Since(start)
		fmt.Printf("Total runtime: %v\n", duration)
	}
}
