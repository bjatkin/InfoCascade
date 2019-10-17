package main

import (
	"fmt"
	"math/rand"
)

func example1(days, households int) {
	truth := 0.1
	faith := 0.1
	var allData stats
	for d := 0; d < days; d++ {
		var dayData stats
		correct := randWell([]int{wellA, wellB, wellC})
		for i := 0; i < households; i++ {
			h := houseHold{
				faith: faith,
			}
			signal := correct

			if rand.Float64() > truth {
				signal = not(correct)
				// fmt.Printf("Lied!\n")
			}

			choice := h.informedChoice(signal)
			result := stat{
				correctChoice: false,
				gotWater:      false,
				liedTo:        signal != correct,
				choice:        choice,
			}
			if choice == correct {
				result.correctChoice = true
				result.gotWater = true
			}

			dayData.data = append(dayData.data, result)
			allData.data = append(allData.data, result)
		}
		fmt.Printf("day %d\nCorrect: %.2f%%, Lies %.2f%%, Correct from Lies %.2f%%, Got Water %.2f%%\n",
			d+1, dayData.Correct(), dayData.Lies(), dayData.CorrectFromLies(), dayData.GotWater())
		dayData = stats{}
	}
	fmt.Printf("All Data:\nCorrect: %.2f%%, Lies %.2f%%, Correct from Lies %.2f%%, Got Water %.2f%%\n",
		allData.Correct(), allData.Lies(), allData.CorrectFromLies(), allData.GotWater())
}
