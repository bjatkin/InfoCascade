package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const wellA = 1
const wellB = 2
const wellC = 3
const wellCorrect = 1.0 / 3.0

func main() {
	rand.Seed(time.Now().Unix())
	days := 20
	households := 50
	ex := strings.ToLower(os.Args[1])
	if ex == "1" || ex == "all" {
		fmt.Printf("Example 1\n")
		example1(days, households)
		fmt.Printf("\n\n")
	}
	if ex == "2" || ex == "all" {
		fmt.Printf("Example 1\n")
		example2(days, households)
		fmt.Printf("\n\n")
	}
	if ex == "2b" || ex == "all" {
		fmt.Printf("Example 2\n")
		example2B(days, households)
		fmt.Printf("\n\n")
	}
	if ex == "3" || ex == "all" {
		fmt.Printf("Example 3\n")
		example3(days, households)
		fmt.Printf("\n\n")
	}
	if ex == "3b" || ex == "all" {
		fmt.Printf("Example 3B\n")
		example3B(days, households)
		fmt.Printf("\n\n")
	}
	if ex == "4a" || ex == "all" {
		fmt.Printf("Example 4A\n")
		example4(days, oneCycleNetwork)
		fmt.Printf("\n\n")
	}
	if ex == "4b" || ex == "all" {
		fmt.Printf("Example 4B\n")
		example4(days, twoCycleNetwork)
		fmt.Printf("\n\n")
	}
	if ex == "4c" || ex == "all" {
		fmt.Printf("Example 4C\n")
		example4(days, completeCycleNetwork)
		fmt.Printf("\n\n")
	}
	if ex == "4d" || ex == "all" {
		fmt.Printf("Example 4D\n")
		example4(days, adHocNetwork)
		fmt.Printf("\n\n")
	}
}
