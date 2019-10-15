package main

import (
	"strings"
)

type stat struct {
	correctChoice bool
	gotWater      bool
	liedTo        bool
	choice        int
}

type stats struct {
	data            []stat
	statsSet        bool
	lies            float64
	correct         float64
	correctFromLies float64
	gotWater        float64
	choices         string
}

func (s *stats) calculateStats() {
	if s.statsSet {
		return
	}
	correct := 0.0
	water := 0.0
	correctFromLies := 0.0
	totalLies := 0.0
	choices := []string{}

	for _, i := range s.data {
		if i.correctChoice {
			correct++
		}
		if i.gotWater {
			water++
		}
		if i.liedTo {
			totalLies++
			if i.correctChoice {
				correctFromLies++
			}
		}
		choices = append(choices, wellString(i.choice))
	}

	count := float64(len(s.data))
	s.correct = correct / count * 100
	s.lies = totalLies / count * 100
	s.correctFromLies = correctFromLies / totalLies * 100
	s.gotWater = water / count * 100
	s.choices = strings.Join(choices, ",")
	s.statsSet = true
}

func (s *stats) Lies() float64 {
	s.calculateStats()
	return s.lies
}

func (s *stats) Correct() float64 {
	s.calculateStats()
	return s.correct
}

func (s *stats) CorrectFromLies() float64 {
	s.calculateStats()
	return s.correctFromLies
}

func (s *stats) GotWater() float64 {
	s.calculateStats()
	return s.correct
}

func (s *stats) Choices() string {
	s.calculateStats()
	return s.choices
}

func bays(sequ []int, getWater float64) (float64, float64, float64) {
	invWater := 1 - getWater
	pMat := [3][3]float64{
		//				A			B			C
		[3]float64{getWater, invWater / 2, invWater / 2}, //P(a|state)
		[3]float64{invWater / 2, getWater, invWater / 2}, //P(b|state)
		[3]float64{invWater / 2, invWater / 2, getWater}, //P(c|state)
	}
	pSgiveA := 1.0
	pSgiveB := 1.0
	pSgiveC := 1.0

	for _, o := range sequ {
		row := 0
		if o == wellB {
			row = 1
		}
		if o == wellC {
			row = 2
		}

		pSgiveA *= pMat[row][0]
		pSgiveB *= pMat[row][1]
		pSgiveC *= pMat[row][2]
	}

	pSequ := pSgiveA*wellCorrect + pSgiveB*wellCorrect + pSgiveC*wellCorrect
	pA := pSgiveA * wellCorrect / pSequ
	pB := pSgiveB * wellCorrect / pSequ
	pC := pSgiveC * wellCorrect / pSequ
	return pA, pB, pC
}

func wellString(well int) string {
	if well == wellA {
		return "Well A"
	}
	if well == wellB {
		return "Well B"
	}
	return "Well C"
}
