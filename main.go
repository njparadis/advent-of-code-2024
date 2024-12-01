package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INPUT_FOLDER = "./inputs"

func day1() {
	file, err := os.Open(fmt.Sprintf("%s/day_1.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var leftList []int = make([]int, 0)
	var rightList []int = make([]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var number1, number2 int
		fmt.Sscanf(scanner.Text(), "%d %d", &number1, &number2)
		leftList = append(leftList, number1)
		rightList = append(rightList, number2)
	}
	day1part1(leftList, rightList)
	day1part2(leftList, rightList)
}

func day1part1(leftList []int, rightList []int) {
	sort.Ints(leftList)
	sort.Ints(rightList)

	var totalDistance int = 0

	for i := 0; i < len(leftList); i++ {
		distance := leftList[i] - rightList[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}

	fmt.Println("Total distance: ", totalDistance)
}

func day1part2(leftList []int, rightList []int) {
	var rightListOccurences map[int]int = make(map[int]int)
	for i := 0; i < len(rightList); i++ {
		rightListOccurences[rightList[i]]++
	}

	var similarityScore int = 0
	for i := 0; i < len(leftList); i++ {
		if rightListOccurences[leftList[i]] > 0 {
			similarityScore += leftList[i] * rightListOccurences[leftList[i]]
		}
	}
	fmt.Println("Similarity score: ", similarityScore)
}

func main() {
	fmt.Println("Day 1")
	day1()
}
