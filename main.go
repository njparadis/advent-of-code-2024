package main

import (
	"flag"

	"github.com/njparadis/advent-of-code-2024/solutions"
)

func main() {
	day := flag.String("day", "", "Day to run")
	flag.Parse()
	switch string(*day) {
	case "1":
		solutions.Day1()
	case "2":
		solutions.Day2()
	case "3":
		solutions.Day3()
	case "4":
		solutions.Day4()
	default:
		solutions.Day1()
		solutions.Day2()
		solutions.Day3()
		solutions.Day4()
	}
}
