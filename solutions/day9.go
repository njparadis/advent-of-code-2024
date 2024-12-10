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

	var diskUsage []int
	fileSizes := make(map[int]int)

	for i, v := range diskMap {
		fileId := -1
		if i%2 == 0 {
			fileId = i / 2
			fileSizes[fileId] = v
		}
		block := make([]int, v)
		for j := 0; j < v; j++ {
			block[j] = fileId
		}
		diskUsage = append(diskUsage, block...)
	}
	fmt.Println(fileSizes)
	diskUsage = []int{
		0, 0, -1, -1, -1,
		1, 1, 1, -1, -1, -1,
		2, -1, -1, -1,
		3, 3, 3, -1,
		4, 4, -1,
		5, 5, 5, 5, -1,
		6, 6, 6, 6, -1,
		7, 7, 7, -1,
		8, 8, 8, 8,
		9, 9,
	}

	diskCopy := make([]int, len(diskUsage))
	copy(diskCopy, diskUsage)

	// day9part1(diskUsage)
	day9part2(diskCopy, fileSizes)
}

func day9part1(disk []int) {
	left := 0
	right := len(disk) - 1

	for left < right {
		for left < len(disk) && disk[left] >= 0 {
			left++
		}

		// Move the right pointer to the first positive number
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

func day9part2(disk []int, fileSizes map[int]int) {
	fmt.Println(disk)
	// Try to move files in decreasing order of file IDs
	for fileID := len(fileSizes); fileID > 0; fileID-- {
		// Get the size of the file
		fileSize := fileSizes[fileID]

		// Try to find a contiguous block of free space large enough
		spaceStart := -1
		freeSpaceCount := 0
		for i := 0; i < len(disk); i++ {
			if disk[i] == -1 { // Free space
				if spaceStart == -1 {
					spaceStart = i // Mark the start of the free space block
				}
				freeSpaceCount++
			} else {
				// If we hit a non-free space, reset the free space count
				spaceStart = -1
				freeSpaceCount = 0
			}

			// If we found enough contiguous free space, try to move the file
			if freeSpaceCount >= fileSize {
				// Move the file to the leftmost available space
				for j := 0; j < fileSize; j++ {
					disk[spaceStart+j] = fileID // Place the file in the free space
				}

				// Clear the old position of the file (assuming the file has a size)
				for k := 0; k < len(disk); k++ {
					if disk[k] == fileID {
						disk[k] = -1
					}
				}

				// fmt.Printf("Moved file ID %d of size %d to position %d\n", fileID, fileSize, spaceStart)
				break
			}
		}
	}
	fmt.Println(disk)
	fmt.Println("Checksum:", checksum(disk))

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
