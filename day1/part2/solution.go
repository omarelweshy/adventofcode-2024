package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) distanceBetweenLists(firstColumn, secondColumn []int) int {
	secondColumnMap := make(map[int]int)
	for i := 0; i < len(secondColumn); i++ {

		_, exists := secondColumnMap[secondColumn[i]]
		if !exists {
			secondColumnMap[secondColumn[i]] = 0
		}
		secondColumnMap[secondColumn[i]]++
	}
	result := 0
	for i := 0; i < len(firstColumn); i++ {
		value, exists := secondColumnMap[firstColumn[i]]
		if exists {
			result += value * firstColumn[i]
			fmt.Println(value)
		}
	}
	return int(math.Abs(float64(result)))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var firstColumn []int
	var secondColumn []int

	for {
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		parts := strings.Fields(line)

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input.")
			continue
		}

		firstColumn = append(firstColumn, num1)
		secondColumn = append(secondColumn, num2)
	}

	solution := Solution{}
	result := solution.distanceBetweenLists(firstColumn, secondColumn)
	fmt.Println(result)
}
