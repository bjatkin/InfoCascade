package main

import (
	"fmt"
	"math/rand"
)

func example4(days int, networkType int) {
	truth := 0.9
	faith := 0.9
	network := oneCycle()
	if networkType == twoCycleNetwork {
		network = twoCycle()
	}
	if networkType == completeCycleNetwork {
		network = completeCycle()
	}
	if networkType == adHocNetwork {
		network = adHoc()
	}

	var allData stats
	for d := 0; d < days; d++ {
		var dayData stats
		var h1, h2, h3 string
		correct := randWell([]int{wellA, wellB, wellC})
		network.Reset()
		for network.Next() {
			h := network.HouseHold()
			h.faith = faith
			if h1 != "" && h2 != "" && h3 == "" {
				h3 = h.name
			}
			if h1 != "" && h2 == "" {
				h2 = h.name
			}
			if h1 == "" {
				h1 = h.name
			}
			signal := correct
			if rand.Float64() > truth {
				signal = not(correct)
			}

			observe := h.askFrineds()
			choice := h.informedObserveAndChoose(signal, observe)
			if h.name == h1 || h.name == h2 || h.name == h3 {
				choice = correct
				signal = correct
			}
			h.choice = choice
			s := stat{
				correctChoice: correct == choice,
				gotWater:      false,
				choice:        choice,
			}
			if s.correctChoice {
				s.gotWater = true
			}
			dayData.data = append(dayData.data, s)
			allData.data = append(allData.data, s)
		}
		fmt.Printf("day %d\nCorrect: %.2f%%(%d), First Choices [%s], First Household [%s, %s, %s], Got Water %.2f%%\n",
			d+1,
			dayData.Correct(),
			correct,
			dayData.Choices()[:27],
			h1, h2, h3,
			dayData.GotWater())
	}
	fmt.Printf("All Data:\nCorrect: %.2f%%, Got Water %.2f%%\n",
		allData.Correct(), allData.GotWater())

}

const oneCycleNetwork = 1

func oneCycle() houseHoldNetwork {
	a := houseHold{name: "a"}
	b := houseHold{name: "b"}
	c := houseHold{name: "c"}
	d := houseHold{name: "d"}
	e := houseHold{name: "e"}
	f := houseHold{name: "f"}
	g := houseHold{name: "g"}
	a.connections = []*houseHold{&b, &g}
	b.connections = []*houseHold{&a, &c}
	c.connections = []*houseHold{&b, &d}
	d.connections = []*houseHold{&c, &e}
	e.connections = []*houseHold{&d, &f}
	f.connections = []*houseHold{&e, &g}
	g.connections = []*houseHold{&f, &a}

	return houseHoldNetwork{
		households: []houseHold{
			a, b, c, d, e, f, g,
		},
		index: -1,
	}
}

const twoCycleNetwork = 2

func twoCycle() houseHoldNetwork {
	a := houseHold{name: "a"}
	b := houseHold{name: "b"}
	c := houseHold{name: "c"}
	d := houseHold{name: "d"}
	e := houseHold{name: "e"}
	f := houseHold{name: "f"}
	g := houseHold{name: "g"}
	a.connections = []*houseHold{&b, &g, &f, &c}
	b.connections = []*houseHold{&g, &a, &c, &d}
	c.connections = []*houseHold{&b, &a, &e, &d}
	d.connections = []*houseHold{&c, &b, &e, &f}
	e.connections = []*houseHold{&c, &d, &f, &g}
	f.connections = []*houseHold{&e, &d, &g, &a}
	g.connections = []*houseHold{&f, &e, &a, &b}

	return houseHoldNetwork{
		households: []houseHold{
			a, b, c, d, e, f, g,
		},
		index: -1,
	}
}

const completeCycleNetwork = 3

func completeCycle() houseHoldNetwork {
	a := houseHold{name: "a"}
	b := houseHold{name: "b"}
	c := houseHold{name: "c"}
	d := houseHold{name: "d"}
	e := houseHold{name: "e"}
	f := houseHold{name: "f"}
	g := houseHold{name: "g"}
	a.connections = []*houseHold{&b, &c, &d, &e, &f, &g}
	b.connections = []*houseHold{&a, &c, &d, &e, &f, &g}
	c.connections = []*houseHold{&a, &b, &d, &e, &f, &g}
	d.connections = []*houseHold{&a, &b, &c, &e, &f, &g}
	e.connections = []*houseHold{&a, &b, &c, &d, &f, &g}
	f.connections = []*houseHold{&a, &b, &c, &d, &e, &g}
	g.connections = []*houseHold{&a, &b, &c, &d, &e, &f}

	return houseHoldNetwork{
		households: []houseHold{
			a, b, c, d, e, f, g,
		},
		index: -1,
	}
}

const adHocNetwork = 4

func adHoc() houseHoldNetwork {
	a := houseHold{name: "a"}
	b := houseHold{name: "b"}
	c := houseHold{name: "c"}
	d := houseHold{name: "d"}
	e := houseHold{name: "e"}
	f := houseHold{name: "f"}
	g := houseHold{name: "g"}
	h := houseHold{name: "h"}
	i := houseHold{name: "i"}
	j := houseHold{name: "j"}
	k := houseHold{name: "k"}
	l := houseHold{name: "l"}
	m := houseHold{name: "m"}
	n := houseHold{name: "n"}
	o := houseHold{name: "o"}
	p := houseHold{name: "p"}
	q := houseHold{name: "q"}
	a.connections = []*houseHold{&b, &c}
	b.connections = []*houseHold{&a, &c}
	c.connections = []*houseHold{&a, &c, &b}
	d.connections = []*houseHold{&e, &h}
	e.connections = []*houseHold{&d, &g, &f}
	f.connections = []*houseHold{&e, &g, &j}
	g.connections = []*houseHold{&e, &f, &j, &h, &i}
	h.connections = []*houseHold{&d, &g, &i, &q}
	i.connections = []*houseHold{&g, &h, &j, &o}
	j.connections = []*houseHold{&g, &f, &i, &k}
	k.connections = []*houseHold{&j, &m, &l}
	l.connections = []*houseHold{&k, &m, &n}
	m.connections = []*houseHold{&o, &k, &l, &n}
	n.connections = []*houseHold{&l, &m, &o}
	o.connections = []*houseHold{&m, &i, &q, &p, &n}
	p.connections = []*houseHold{&o, &q, &n}
	q.connections = []*houseHold{&h, &o, &p}

	return houseHoldNetwork{
		households: []houseHold{
			a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q,
		},
		index: -1,
	}
}
