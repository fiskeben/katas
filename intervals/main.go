package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var includesArg = flag.String("include", "", "a comma separated list of sequences to include")
	var excludesArg = flag.String("exclude", "", "a comma separated list of sequences to exclude")
	flag.Parse()

	includes, err := parseSequences(*includesArg)
	if err != nil || len(includes) == 0 {
		failAndExit(err)
	}
	excludes, err := parseSequences(*excludesArg)
	if err != nil {
		failAndExit(err)
	}

	includes, err = mergeSequences(includes)
	if err != nil {
		failAndExit(err)
	}

	if len(excludes) == 0 {
		fmt.Printf("Result:\n\n\t%v\n", includes)
		os.Exit(0)
	}

	excludes, err = mergeSequences(excludes)
	if err != nil {
		failAndExit(err)
	}

	seq, err := subtractSequences(includes, excludes)
	if err != nil {
		failAndExit(err)
	}

	fmt.Printf("Result:\n\n\t%v\n", seq)
}

func parseSequences(seq string) (intervalList, error) {
	if seq == "" {
		return intervalList{}, nil
	}

	intervals := strings.Split(seq, ",")
	sequences := make(intervalList, len(intervals))
	for i, s := range intervals {
		seq := interval{}
		numbers := strings.Split(s, "-")
		if len(numbers) == 1 {
			if numbers[0] == "" {
				continue
			}
			number, err := strconv.ParseUint(numbers[0], 10, 32)
			if err != nil {
				return nil, fmt.Errorf("'%s' is not a number (%v)", numbers[0], err)
			}
			sequences[i] = interval{From: number, To: number}
			continue
		}

		start, err := strconv.ParseUint(numbers[0], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse start of interval: %s", s)
		}
		seq.From = start

		end, err := strconv.ParseUint(numbers[1], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse end of interval: %s", s)
		}
		seq.To = end

		if seq.From > seq.To {
			return nil, fmt.Errorf("not a valid sequence: %s", s)
		}
		sequences[i] = seq
	}
	return sequences, nil
}

func mergeSequences(seqs intervalList) (intervalList, error) {
	if len(seqs) == 1 {
		return seqs, nil
	}

	sort.Sort(intervalList(seqs))

	i := 1
	curr := seqs[0]
	res := intervalList{}

	for i < len(seqs) {
		next := seqs[i]
		i++
		if overlapOrConsequtive(curr, next) {
			curr = merge(curr, next)
			if i == len(seqs) {
				res = append(res, curr)
			}
		} else {
			res = append(res, curr)
			if i == len(seqs) {
				res = append(res, next)
			}
			curr = next
		}
	}
	return res, nil
}

func merge(first, second interval) interval {
	start := first.From
	if start > second.From {
		start = second.From
	}
	end := first.To
	if end < second.To {
		end = second.To
	}

	return interval{From: start, To: end}
}

func overlap(a, b *interval) bool {
	if a == nil || b == nil {
		return false
	}
	return a.To >= b.From && b.To >= a.From
}

func overlapOrConsequtive(a, b interval) bool {
	return overlap(&a, &b) || b.From-1 == a.To
}

func failAndExit(err error) {
	flag.Usage()
	os.Exit(1)
}

func subtractSequences(includes, excludes intervalList) (intervalList, error) {
	if len(excludes) == 0 {
		return includes, nil
	}

	res := intervalList{}

	includeCounter := 0
	excludeCounter := 0

	toInclude := &includes[includeCounter]
	toExclude := &excludes[excludeCounter]

	for toInclude != nil || toExclude != nil {
		if less(toInclude, toExclude) {
			res = append(res, *toInclude)
			includeCounter++
			toInclude = get(includeCounter, includes)
			continue
		}

		if less(toExclude, toInclude) {
			excludeCounter++
			toExclude = get(excludeCounter, excludes)
			continue
		}

		if overlap(toInclude, toExclude) {
			subs := subtract(*toInclude, *toExclude)
			if len(subs) == 2 {
				res = append(res, subs[0])
				toInclude = &subs[1]
			} else {
				if less(&subs[0], toExclude) {
					res = append(res, subs[0])
					includeCounter++
					toInclude = get(includeCounter, includes)
				} else {
					toInclude = &subs[0]
					excludeCounter++
					toExclude = get(excludeCounter, excludes)
				}
			}
		}
	}

	return res, nil
}

func subtract(a, b interval) intervalList {
	res := intervalList{}
	if a.From < b.From {
		res = append(res, interval{a.From, b.From - 1})

	}
	if a.To > b.To {
		res = append(res, interval{b.To + 1, a.To})
	}
	return res
}

// less returns true if a's To is less than or equal to b's To
func less(a, b *interval) bool {
	if a == nil {
		return false
	}
	if b == nil {
		return true
	}
	return a.To < b.From
}

func get(i int, list intervalList) *interval {
	if i < len(list) {
		return &list[i]
	}
	return nil
}

type interval struct {
	From uint64
	To   uint64
}

func (i interval) String() string {
	return fmt.Sprintf("%d-%d", i.From, i.To)
}

type intervalList []interval

func (b intervalList) String() string {
	items := make([]string, len(b))
	for i, item := range b {
		items[i] = item.String()
	}
	return strings.Join(items, " ")
}

func (b intervalList) Len() int {
	return len(b)
}

func (b intervalList) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b intervalList) Less(i, j int) bool {
	return b[i].From < b[j].From
}
