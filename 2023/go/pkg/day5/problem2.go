package day5

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Seed struct {
	start  int
	length int
}

type Map struct {
	from   string
	to     string
	ranges []Range
}

type Range struct {
	dest int
	src  int
	rng  int
}

func SolveProblem2(lines []string) string {
	rangesStr := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seeds := []Seed{}

	for i := 0; i < len(rangesStr); i += 2 {
		start, _ := strconv.Atoi(rangesStr[i])
		length, _ := strconv.Atoi(rangesStr[i+1])

		seeds = append(seeds, Seed{start, length})
	}

	highestValuePossible := 0
	maps := map[string]Map{}

	current := "seed"

	// create ranges
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}

		if !unicode.IsDigit(rune(line[0])) {
			args := strings.Split(line, "-")
			from := args[0]
			to := strings.Split(args[2], " ")[0]

			current = from
			maps[current] = Map{from, to, []Range{}}

			continue
		}

		rng := createRange(line)

		if rng.dest+rng.rng > highestValuePossible {
			highestValuePossible = rng.dest + rng.rng
		}

		if rng.src+rng.rng > highestValuePossible {
			highestValuePossible = rng.src + rng.rng
		}

		ranges := maps[current].ranges
		ranges = append(ranges, rng)
		maps[current] = Map{maps[current].from, maps[current].to, ranges}
	}

	// create negative ranges
	for _, m := range maps {
		sort.Slice(m.ranges, func(i, j int) bool {
			return m.ranges[i].src < m.ranges[j].src
		})

		start := 0

		for i := 0; i < len(m.ranges); i++ {
			rng := m.ranges[i]

			if rng.src > start {
				slices.Insert(m.ranges, i, Range{start, start, rng.src - start})
				i++
			}

			start = rng.src + rng.rng
		}
	}

	// lowest := math.MaxInt64

	// for _, seed := range seeds {
	// 	fmt.Printf("starting %v\n", seed)

	// 	remaining := seed.length
	// 	start := seed.start

	// 	for remaining > 0 {
	// 		startLocation, consumed := walk(start, remaining, "seed", maps)

	// 		remaining -= consumed
	// 		start += consumed

	// 		if consumed > 1 {
	// 			fmt.Printf("consumed %v\n", consumed)
	// 		}

	// 		if startLocation < lowest {
	// 			lowest = startLocation
	// 		}
	// 	}

	// 	fmt.Printf("finished %v\n", seed)
	// }

	return fmt.Sprint(maps["soil"])
}

func walk(start, remaining int, s string, maps map[string][]Range) {
	panic("unimplemented")
}

func createRange(line string) Range {
	items := strings.Fields(line)

	dest, _ := strconv.Atoi(items[0])
	src, _ := strconv.Atoi(items[1])
	rng, _ := strconv.Atoi(items[2])

	return Range{dest, src, rng}
}
