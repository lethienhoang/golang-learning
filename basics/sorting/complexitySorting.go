package main

import (
	"fmt"
	"sort"
)

type Programmer struct {
	Age int
}

type byAge []Programmer

func (p byAge) Len() int {
	return len(p)
}

func (p byAge) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byAge) Less(i int, j int) bool {
	return p[i].Age < p[j].Age
}

func main() {
	programmers := []Programmer{
		Programmer{Age: 30},
		Programmer{Age: 20},
		Programmer{Age: 50},
		Programmer{Age: 1000},
	}

	byAges := byAge(programmers)

	for i := 0; i < byAges.Len(); i++ {
		for j := 1; j < byAges.Len(); j++ {
			if byAges.Less(i, j) {
				byAges.Swap(i, j)
			}
		}
	}

	fmt.Println(byAges)

	sort.Sort(byAge(programmers))

	fmt.Println(programmers)
}
