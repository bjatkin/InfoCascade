package main

import (
	"fmt"
	"math/rand"
)

func example3(days, household int) {
	getWater := .9
	utils := [3]float64{0.8, 0.15, 0.05}
	var allData stats
	for d := 0; d < days; d++ {
		observations := []int{}
		correct := randWell([]int{wellA, wellB, wellC})
		var dayData stats
		for h := 0; h < household; h++ {
			house := houseHold{
				wellAUitl: utils[0],
				wellBUtil: utils[1],
				wellCUtil: utils[2],
				getWater:  getWater,
			}

			choice := house.utilityObserveAndChoose(observations)
			s := stat{
				correctChoice: false,
				gotWater:      false,
				choice:        choice,
			}
			if choice == correct {
				s.correctChoice = true
				if rand.Float64() < getWater {
					s.gotWater = true
				}
			}
			dayData.data = append(dayData.data, s)
			allData.data = append(allData.data, s)
		}
		fmt.Printf("day %d\nCorrect: %.2f%%, First Choices [%s], Got Water %.2f%%\n",
			d, dayData.Correct(), dayData.Choices()[:27], dayData.GotWater())
	}
	fmt.Printf("All Data:\nCorrect: %.2f%%, Got Water %.2f%%\n",
		allData.Correct(), allData.GotWater())
}
