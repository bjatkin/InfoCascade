package main

import (
	"fmt"
	"math/rand"
)

func example2(days, households int) {
	truthPercent := 0.9
	var allData stats
	for d := 0; d < days; d++ {
		observations := []int{}
		correct := randWell([]int{wellA, wellB, wellC})
		var dayData stats
		for h := 0; h < households; h++ {
			house := houseHold{
				truthPercent: truthPercent,
			}
			choice := house.observerAndChoose(observations)
			observations = append(observations, choice)
			s := stat{
				correctChoice: false,
				gotWater:      false,
				choice:        choice,
			}
			if choice == correct {
				s.correctChoice = true
				s.gotWater = true
			}
			dayData.data = append(dayData.data, s)
			allData.data = append(allData.data, s)
		}
		fmt.Printf("day %d\nCorrect: %.2f%%, First Choices [%s], Got Water %.2f%%\n",
			d+1, dayData.Correct(), dayData.Choices()[:27], dayData.GotWater())
	}
	fmt.Printf("All Data:\nCorrect: %.2f%%, Got Water %.2f%%\n",
		allData.Correct(), allData.GotWater())
}

func example2B(days, households int) {
	truth := 0.9
	faith := 0.95
	var allData stats
	for d := 0; d < days; d++ {
		observations := []int{}
		correct := randWell([]int{wellA, wellB, wellC})
		var dayData stats
		for h := 0; h < households; h++ {
			house := houseHold{
				truthPercent: truth,
				faith:        faith,
			}
			signal := correct
			if rand.Float64() > truth {
				signal = not(correct)
			}

			choice := house.informedObserveAndChoose(signal, observations)
			observations = append(observations, choice)
			s := stat{
				correctChoice: false,
				gotWater:      false,
				choice:        choice,
				liedTo:        signal == correct,
			}
			if choice == correct {
				s.correctChoice = true
				s.gotWater = true
			}
			dayData.data = append(dayData.data, s)
			allData.data = append(allData.data, s)
		}
		fmt.Printf("day %d\nCorrect: %.2f%%, First Choices [%s], Got Water %.2f%%\n",
			d+1, dayData.Correct(), dayData.Choices()[:27], dayData.GotWater())
	}
	fmt.Printf("All Data:\nCorrect: %.2f%%, Got Water %.2f%%\n",
		allData.Correct(), allData.GotWater())
}
