package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func parsed(lines []string) int {

	var check [9]string = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var ans int32

	mn := math.MaxInt
	mx := math.MinInt
	for _, line := range lines {
		mn = math.MaxInt
		mx = math.MinInt
		var firstVal int
		var lastVal int
		for i, num := range check {

			curr := strings.Index(line, num)
			if curr != -1 {
				if mn > curr {
					firstVal = i + 1
					mn = curr
				}
			}
			curr = strings.LastIndex(line, num)
			if curr != -1 {
				if curr > mx {
					mx = curr
					lastVal = i + 1
				}
			}
		}
		digFirst := strings.IndexAny(line, "123456789")
		digLast := strings.LastIndexAny(line, "123456789")
		if digFirst != -1 {
			if mn > digFirst {
				firstVal = int(line[digFirst] - '0')
				mn = digFirst
			}
			if mx < digLast {
				lastVal = int(line[digLast] - '0')
				mx = digLast
			}
		}
		ans += int32(firstVal*10 + lastVal)
	}

	fmt.Println("answer", ans)

	return 0

}

func parse(lines []string) int {
	var check [9]string = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var ans int

	for _, line := range lines {
		firstocc := -1
		lastocc := -1
		firstnum := int(line[strings.IndexAny(line, "123456789")])
		lastnum := int(line[strings.LastIndexAny(line, "123456789")])
		var lastIndex int = int(line[strings.LastIndexAny(line, "123456789")])
		var firstIndex int = int(line[strings.IndexAny(line, "123456789")])
		for i, num := range check {
			curr := strings.LastIndex(line, num)
			if lastocc < curr && curr > lastIndex {
				lastnum = i
				lastIndex = curr
			}
			lastocc = max(strings.LastIndex(line, num), lastocc)
			curr = strings.Index(line, num)
			if firstocc > curr && curr < firstIndex {
				firstnum = i
				firstIndex = curr
			}
			firstocc = min(strings.Index(line, num), firstocc)
		}
		ans += firstnum*10 + lastnum
	}
	fmt.Println(ans)
	return ans
}

func main() {
	// Open the file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	parsed(lines)

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
