package main

import (
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
	days := 5
	households := 500
	ex := strings.ToLower(os.Args[1])
	if ex == "1" || ex == "all" {
		example1(days, households)
	}
	if ex == "2" || ex == "all" {
		example2(days, households)
	}
	if ex == "2b" || ex == "all" {
		example2B(days, households)
	}
	if ex == "3" || ex == "all" {
		example3(days, households)
	}
	if ex == "4A" || ex == "all" {
		example4(days, oneCycleNetwork)
	}
	if ex == "4B" || ex == "all" {
		example4(days, twoCycleNetwork)
	}
	if ex == "4C" || ex == "all" {
		example4(days, completeCycleNetwork)
	}
	if ex == "4D" || ex == "all" {
		example4(days, adHocNetwork)
	}
}
