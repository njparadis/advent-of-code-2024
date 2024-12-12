package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day9() {
	fmt.Println("Day 9")
	file, err := os.Open(fmt.Sprintf("%s/day_9.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var diskMap []int

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")

		for _, c := range chars {
			num, err := strconv.Atoi(c)
			if err != nil {
				fmt.Println(err)
				continue
			}
			diskMap = append(diskMap, num)
		}
	}

	day9part1(diskUsage(diskMap))
	day9part2(diskUsage(diskMap))
}

func day9part1(disk []int) {
	left := 0
	right := len(disk) - 1

	for left < right {
		for left < len(disk) && disk[left] >= 0 {
			left++
		}

		for right >= 0 && disk[right] < 0 {
			right--
		}

		// If the pointers have not crossed, swap the values
		if left < right {
			disk[left], disk[right] = disk[right], disk[left]
		}
	}

	fmt.Println("Checksum:", checksum(disk))
}

func day9part2(disk []int) {
	for end := len(disk) - 1; end >= 0; end-- {
		if disk[end] < 0 {
			continue
		}

		// find where the start of the block is
		var start int
		for i := end; i >= 0; i-- {
			if disk[i] != disk[end] {
				start = i + 1
				break
			}
		}
		length := end - start + 1

		found := false
		for j := 0; j < start; j++ {
			if disk[j] >= 0 {
				continue
			}

			enoughSpace := true
			for k := j; k < j+length; k++ {
				if disk[k] >= 0 {
					enoughSpace = false
					break
				}

			}

			if enoughSpace {
				for k := 0; k < length; k++ {
					disk[j+k], disk[start+k] = disk[start+k], disk[j+k]
				}
				found = true
				break
			}

		}
		if !found {
			end = start
		}
	}
	fmt.Println("Checksum:", checksum(disk))
}

func diskUsage(diskMap []int) []int {
	usage := make([]int, 0)
	for i, v := range diskMap {
		fileId := -1
		if i%2 == 0 {
			fileId = i / 2
		}
		block := make([]int, v)
		for j := 0; j < v; j++ {
			block[j] = fileId
		}
		usage = append(usage, block...)
	}
	return usage
}

func checksum(disk []int) int {
	sum := 0
	for i, v := range disk {
		if v >= 0 {
			sum += i * v
		}
	}
	return sum
}
