package main

import (
	"math/rand"
)

type houseHold struct {
	faith       float64
	getWater    float64
	wellAUitl   float64
	wellBUtil   float64
	wellCUtil   float64
	choice      int
	connections []*houseHold
}

func (h houseHold) askFrineds() []int {
	ret := []int{}
	if len(h.connections) == 0 {
		return ret
	}

	for i := 0; i < len(h.connections); i++ {
		info := h.connections[i].choice
		if info != 0 {
			ret = append(ret, info)
		}
	}

	return ret
}

func (h houseHold) informedChoice(signal int) int {
	if rand.Float64() < h.faith {
		return signal
	}
	return not(signal)
}

func (h houseHold) informedObserveAndChoose(signal int, public []int) int {
	if len(public) == 0 || len(public) == 1 {
		return h.informedChoice(signal)
	}
	return h.observerAndChoose(public)
}

func (h houseHold) observerAndChoose(public []int) int {
	if len(public) == 0 || len(public) == 1 {
		return randWell([]int{wellA, wellB, wellC})
	}
	pA, pB, pC := bays(public, h.getWater)

	if pA > pB && pA > pC {
		return wellA
	}
	if pB > pC {
		return wellB
	}
	return wellC
}

func (h houseHold) utilityObserveAndChoose(public []int) int {
	pA, pB, pC := wellCorrect, wellCorrect, wellCorrect
	if len(public) != 0 && len(public) != 1 {
		pA, pB, pC = bays(public, h.getWater)
	}

	uA, uB, uC := h.wellAUitl*pA, h.wellBUtil*pB, h.wellCUtil*pC
	if uA > uB && uA > uC {
		return wellA
	}
	if uB > uC {
		return wellB
	}
	return wellC
}

type houseHoldNetwork struct {
	households []houseHold
	order      []int
	index      int
}

func (hn *houseHoldNetwork) Next() bool {
	if hn.index == -1 {
		l := len(hn.households)
		if l == 0 {
			return false
		}
		//shuffle the network
		hn.order = rand.Perm(l)
	}
	hn.index++
	if hn.index < len(hn.order) {
		return true
	}
	return false
}

func (hn *houseHoldNetwork) HouseHold() houseHold {
	return hn.households[hn.index]
}

func (hn *houseHoldNetwork) Reset() {
	hn.index = -1
}

func not(well int) int {
	ret := [2]int{wellA, wellB}
	if well == wellA {
		ret[0] = wellC
	}
	if well == wellB {
		ret[1] = wellC
	}

	if rand.Float64() < 0.5 {
		return ret[0]
	}
	return ret[1]
}

func randWell(wells []int) int {
	count := len(wells)
	if count == 1 {
		return wells[0]
	}

	w := rand.Float64()

	if w < 1/float64(count) {
		return wells[0]
	}
	return not(wells[0])
}
