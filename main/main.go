package main

import (
	"bufio"
	"fmt"
	"mathskills"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename.txt>")
		os.Exit(0)
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(0)
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			fmt.Printf("Warning: Line %d is empty. Skipping...\n", lineNumber)
			lineNumber++
			continue
		}

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("Error parsing value on line %d: %v\n", lineNumber, err)
			fmt.Printf("Please fix the file and remove any letters or invalid characters.\n")
			os.Exit(0)
		}

		data = append(data, value)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(0)
	}

	if len(data) == 0 {
		fmt.Println("Data file cannot be empty or contains only invalid data")
		os.Exit(0)
	}

	average := mathskills.CalculateAverage(data)
	median := mathskills.CalculateMedian(data)
	variance := mathskills.CalculateVariance(data, float64(average))
	std := mathskills.CalculateStandardDeviation(float64(variance))

	fmt.Printf("Average: %d\n", average)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", std)
}
