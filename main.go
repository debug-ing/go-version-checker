package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("Error: go.mod file not found in the repository")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "go ") {
			version := strings.TrimSpace(strings.TrimPrefix(line, "go "))
			fmt.Printf("The Go version specified in go.mod is: %s\n", version)
			return
		}
	}
	fmt.Println("No Go version found in go.mod")
}
