package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day5() {
	fmt.Println("Day 5")
	file, err := os.Open(fmt.Sprintf("%s/day_5.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	rules := make(map[int][]int)
	total := 0
	totalFixed := 0
	isSecondHalf := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			isSecondHalf = true
			continue
		}

		if !isSecondHalf {
			var number1, number2 int
			_, err := fmt.Sscanf(line, "%d|%d", &number1, &number2)

			if err != nil {
				fmt.Println("Error converting to integer:", err)
			} else {
				if _, ok := rules[number1]; !ok {
					rules[number1] = []int{}
				}
				exists := false
				for _, v := range rules[number1] {
					if v == number2 {
						exists = true
						break
					}
				}
				if !exists {
					rules[number1] = append(rules[number1], number2)
				}
			}

		} else {
			numberStrings := strings.Split(line, ",")
			numbers := make([]int, len(numberStrings))
			lastSeen := make(map[int]int)
			for i, n := range numberStrings {
				num, err := strconv.Atoi(strings.TrimSpace(n))
				if err != nil {
					fmt.Printf("Error parsing number '%s': %v\n", n, err)
					break
				}
				numbers[i] = num
			}

			valid := true
			for i, num := range numbers {
				lastSeen[num] = i
				if rule, ok := rules[num]; ok {
					for _, r := range rule {
						if pos, seen := lastSeen[r]; seen {
							if pos < i {
								valid = false
								break
							}
						}
					}
				}
				if !valid {
					// we need to fix the order
					numbers = topologicalSort(numbers, rules)
					break
				}
			}
			mid := (len(numbers) - 1) / 2
			if valid {
				total += numbers[mid]
			} else {
				totalFixed += numbers[mid]
			}
		}
	}

	fmt.Println("Total: ", total)
	fmt.Println("Total (fixed): ", totalFixed)
}

func topologicalSort(numbers []int, rules map[int][]int) []int {
	// Step 1: Filter relevant rules
	relevantRules := make(map[int][]int)

	// Only keep rules where both num and r are in the numbers slice
	for _, num := range numbers {
		if rule, ok := rules[num]; ok {
			for _, r := range rule {
				// Only add the rule if both num and r are in the sequence
				if contains(numbers, r) {
					relevantRules[num] = append(relevantRules[num], r)
				}
			}
		}
	}

	// Step 2: Initialize in-degree and adjacency list
	inDegree := make(map[int]int)
	adjList := make(map[int][]int)

	// Ensure every number in the sequence has an entry in inDegree
	for _, num := range numbers {
		if _, exists := inDegree[num]; !exists {
			inDegree[num] = 0
		}
	}

	// Create adjacency list and calculate in-degrees
	for src, dests := range relevantRules {
		for _, dest := range dests {
			adjList[src] = append(adjList[src], dest)
			inDegree[dest]++
		}
	}

	// Step 3: Initialize the queue with all nodes having in-degree 0
	queue := []int{}
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	// Step 4: Process the nodes in topological order
	topOrder := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		topOrder = append(topOrder, node)

		// Decrease the in-degree of all successors
		for _, neighbor := range adjList[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Step 5: Check for cycle
	if len(topOrder) != len(numbers) {
		fmt.Println("Cycle detected")
		return []int{}
	}

	// Return the topologically sorted order
	return topOrder
}

// Helper function to check if a number is in the sequence
func contains(numbers []int, num int) bool {
	for _, n := range numbers {
		if n == num {
			return true
		}
	}
	return false
}
