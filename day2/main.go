package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ans int
	curr := 0
	for scanner.Scan() {

		curr++
		line := strings.Split(scanner.Text(), ":")
		games := strings.Split(line[1], ";")
		maxGreen := 1
		maxBlue := 1
		maxRed := 1
		for _, round := range games {

			colors := strings.Split(round, ",")
			for i := 0; i < len(colors); i++ {
				currnum, _ := strconv.Atoi(strings.Trim(colors[i], "greenredblue, "))
				currstr := strings.Trim(colors[i], ", 1234567890")
				fmt.Println(currstr, currnum)
				if currstr == "green" {
					maxGreen = max(maxGreen, currnum)
				}
				if currstr == "red" {
					maxRed = max(maxRed, currnum)
				}
				if currstr == "blue" {
					maxBlue = max(maxBlue, currnum)
				}

			}
		}
		fmt.Println(maxRed, maxGreen, maxBlue)
		ans += maxBlue * maxRed * maxGreen

	}
	fmt.Println(ans)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
