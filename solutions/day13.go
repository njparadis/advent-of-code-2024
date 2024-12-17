package solutions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ClawMachine struct {
	ax, ay         int
	bx, by         int
	prizeX, prizeY int
}

func Day13() {
	fmt.Println("Day 13")
	file, err := os.Open(fmt.Sprintf("%s/day_13.txt", INPUT_FOLDER))
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	machines := make([]ClawMachine, 0)
	buttonARegex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	buttonBRegex := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	lines := []string{}

	// Read all non-blank lines into a buffer
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // Skip blank lines
			lines = append(lines, line)
		}
	}

	// Now parse the data in groups of 3 lines
	for i := 0; i < len(lines); i += 3 {
		if i+2 >= len(lines) { // Ensure there are enough lines for a complete block
			break
		}

		lineA := lines[i]
		lineB := lines[i+1]
		linePrize := lines[i+2]

		// Extract numbers using regex
		aMatches := buttonARegex.FindStringSubmatch(lineA)
		bMatches := buttonBRegex.FindStringSubmatch(lineB)
		prizeMatches := prizeRegex.FindStringSubmatch(linePrize)

		if aMatches == nil || bMatches == nil || prizeMatches == nil {
			continue
		}

		// Convert matches to integers
		ax, _ := strconv.Atoi(aMatches[1])
		ay, _ := strconv.Atoi(aMatches[2])
		bx, _ := strconv.Atoi(bMatches[1])
		by, _ := strconv.Atoi(bMatches[2])
		prizeX, _ := strconv.Atoi(prizeMatches[1])
		prizeY, _ := strconv.Atoi(prizeMatches[2])

		// Append to the list of machines
		machines = append(machines, ClawMachine{
			ax: ax, ay: ay,
			bx: bx, by: by,
			prizeX: prizeX, prizeY: prizeY,
		})
	}

	// machines = []ClawMachine{
	// 	{ax: 94, ay: 34, bx: 22, by: 67, prizeX: 8400, prizeY: 5400},
	// 	{ax: 26, ay: 66, bx: 67, by: 21, prizeX: 12784, prizeY: 12176},
	// 	{ax: 17, ay: 86, bx: 84, by: 37, prizeX: 7870, prizeY: 6450},
	// 	{ax: 69, ay: 23, bx: 27, by: 71, prizeX: 18641, prizeY: 10279},
	// }
	day13part1(machines)
	day13part2(machines)
}

func day13part1(machines []ClawMachine) {
	total := 0
	for _, m := range machines {
		bNumerator := (m.ay * m.prizeX) - (m.ax * m.prizeY)
		bDenominator := (m.ay * m.bx) - (m.ax * m.by)

		bValid := true
		if bNumerator%bDenominator != 0 {
			bValid = false
		}
		bSolveB := bNumerator / bDenominator

		if (m.prizeX-(bSolveB*m.bx))%m.ax != 0 {
			bValid = false
		}
		aSolveB := (m.prizeX - (bSolveB * m.bx)) / m.ax

		aNumerator := (m.by * m.prizeX) - (m.bx * m.prizeY)
		aDenominator := (m.by * m.ax) - (m.ay * m.bx)

		aValid := true
		if aNumerator%aDenominator != 0 {
			aValid = false
		}
		aSolveA := aNumerator / aDenominator

		if (m.prizeX-(aSolveA*m.ax))%m.bx != 0 {
			aValid = false
		}
		bSolveA := (m.prizeX - (aSolveA * m.ax)) / m.bx
		if aSolveA < 0 || aSolveB < 0 || bSolveA < 0 || bSolveB < 0 {
			continue
		}

		var aPrice int
		if aValid {
			aPrice = aSolveA*3 + bSolveA
		}

		var bPrice int
		if bValid {
			bPrice = aSolveB*3 + bSolveB
		}

		if aValid && bValid {
			if aPrice < bPrice {
				total += aPrice
			} else {
				total += bPrice
			}
		} else if aValid {
			total += aPrice
		} else if bValid {
			total += bPrice
		}
	}
	fmt.Println("Total price:", total)
}

func day13part2(machines []ClawMachine) {
	total := 0
	for _, m := range machines {
		m.prizeX += 10000000000000
		m.prizeY += 10000000000000
		bNumerator := (m.ay * m.prizeX) - (m.ax * m.prizeY)
		bDenominator := (m.ay * m.bx) - (m.ax * m.by)

		bValid := true
		if bNumerator%bDenominator != 0 {
			bValid = false
		}
		bSolveB := bNumerator / bDenominator

		if (m.prizeX-(bSolveB*m.bx))%m.ax != 0 {
			bValid = false
		}
		aSolveB := (m.prizeX - (bSolveB * m.bx)) / m.ax

		aNumerator := (m.by * m.prizeX) - (m.bx * m.prizeY)
		aDenominator := (m.by * m.ax) - (m.ay * m.bx)

		aValid := true
		if aNumerator%aDenominator != 0 {
			aValid = false
		}
		aSolveA := aNumerator / aDenominator

		if (m.prizeX-(aSolveA*m.ax))%m.bx != 0 {
			aValid = false
		}
		bSolveA := (m.prizeX - (aSolveA * m.ax)) / m.bx
		if aSolveA < 0 || aSolveB < 0 || bSolveA < 0 || bSolveB < 0 {
			continue
		}

		var aPrice int
		if aValid {
			aPrice = aSolveA*3 + bSolveA
		}

		var bPrice int
		if bValid {
			bPrice = aSolveB*3 + bSolveB
		}

		if aValid && bValid {
			if aPrice < bPrice {
				total += aPrice
			} else {
				total += bPrice
			}
		} else if aValid {
			total += aPrice
		} else if bValid {
			total += bPrice
		}
	}
	fmt.Println("Total price (with updated units):", total)
}
