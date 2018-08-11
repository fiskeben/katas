package main

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// generator holds the state of available numbers and a random number generator.
type generator struct {
	randomizer *rand.Rand
	// groups is a 2D array of available numbers grouped by tens.
	groups [][]int
}

// makeGenerator creates a new generator with a pool of random numbers grouped by tens.
func makeGenerator() *generator {
	t := time.Now()
	src := rand.NewSource(t.Unix())
	randomizer := rand.New(src)
	groups := makeGroups(randomizer)
	return &generator{randomizer: randomizer, groups: groups}
}

// MakeRandomBoard creates the six cards of three rows.
func (g generator) MakeRandomBoard() board {
	rows := []row{}
	for i := 0; i < 18; i++ {
		r := g.generateRow()
		rows = append(rows, r)
	}
	b := board{}
	for j := 0; j < len(rows); j += 3 {
		c := rows[j : j+3]
		b = append(b, c)
	}

	return b
}

//makeGroups randomizes the 90 numbers and puts them into groups of tens.
func makeGroups(randomizer *rand.Rand) [][]int {
	res := [][]int{
		[]int{},
		[]int{},
		[]int{},
		[]int{},
		[]int{},
		[]int{},
		[]int{},
		[]int{},
		[]int{},
	}

	nums := randomizer.Perm(90)
	for _, n := range nums {
		n = n + 1
		index := mapNumberToIndex(n)
		res[index] = append(res[index], n)
	}
	return res
}

// mapNumberToIndex calculates the group a given number should go into.
func mapNumberToIndex(n int) int {
	if n == 90 {
		return 8
	}
	return int(math.Floor(float64(n) / 10))
}

// generateRow generates a row of five random numbers and removes them from the pool.
// The five numbers will not be from the same group.
func (g generator) generateRow() row {
	indexes := g.indexesBasedOnLongestGroups()
	res := []int{}

	for _, i := range indexes {
		r := g.groups[i]
		if len(r) == 0 {
			continue
		}
		n, rest := r[0], r[1:]
		res = append(res, n)
		g.groups[i] = rest
		if len(res) == 5 {
			break
		}
	}

	sort.Ints(res)
	return res
}

// groupWithSize is used for sorting groups of numbers based on size.
type groupWithSize struct {
	Index int
	Size  int
}

// bySize implements sort.Interface
type bySize []groupWithSize

// Len returns the length of the slice.
func (a bySize) Len() int { return len(a) }

// Swap swaps two items identified by indexes.
func (a bySize) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less returns true if the item at position i is less than the one at position j.
func (a bySize) Less(i, j int) bool { return a[i].Size < a[j].Size }

// indexesBasedOnLongestGroups generates a list of pool group indexes to take numbers from.
// It favors groups with the most numbers to avoid exhausting some groups and ending up with
// groups of numbers that can't be used. The smaller groups' indexes are randomized.
func (g generator) indexesBasedOnLongestGroups() []int {
	items := bySize{}

	for i, r := range g.groups {
		items = append(items, groupWithSize{Index: i, Size: len(r)})
	}

	sort.Sort(items)

	last := len(items) - 1
	longestItems := []int{items[last].Index}
	rest := []int{}
	largestLength := items[last].Size

	for i := last - 1; i >= 0; i-- {
		if items[i].Size == largestLength {
			longestItems = append(longestItems, items[i].Index)
			continue
		}
		rest = append(rest, items[i].Index)
	}

	randomIndexes := g.randomizer.Perm(len(rest))
	for _, i := range randomIndexes {
		longestItems = append(longestItems, rest[i])
	}
	return longestItems
}
