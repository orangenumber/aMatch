// Copyright 2020 Gon Yi
// https://gonyyi.com/copyright.txt

package amatch

// Fuzzy compares a with candidates and return index of candidates
// if no candidates are suitable return -1
func Fuzzy(a string, candidates *[]string) int {
	if len(*candidates) == 0 {
		return -1
	}

	scores := make([]int, len(*candidates))
	for idx, val := range *candidates {
		if a == val { // if identical, just pick that
			return idx
		}
		scores[idx] = FuzzyScore(a, val)
	}
	var maxScore int // hold highest score
	var maxIdx int   // hold index of highest score
	for i, e := range scores {
		if i == 0 || e > maxScore {
			maxScore = e
			maxIdx = i
		}
	}
	// if max score was 0, then return -1
	if maxScore == 0 {
		return -1
	}
	return maxIdx
}

func FuzzyScore(a string, b string) int {
	var lastIndex int
	var score int

	// for each character of `a`
	for aIdx, aVal := range a {
		found := false
		for bIdx, bVal := range b[lastIndex:] {
			if aVal == bVal {
				score += 1                  // if found a match, add a point
				if aIdx == 0 && bIdx == 0 { // if first char match, add a point
					score += 1
				} else if bIdx == 0 { // if not the first but a consequence number, add a point
					score += 1
				}

				lastIndex += bIdx + 1
				found = true
				break
			}
		}
		// if match not found, given 0 score
		if !found {
			return 0
		}
	}
	return score
}
