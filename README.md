# aMatch

Fuzzy matching algorithm written in Go

## Usage

`amatch.Fuzzy(a string, candidates *[]string) int`

Fuzzy will evaluate (pointer of) a string slice. Returns an index of highest scored element.


## Calculation

Each of below will count 1 point. Fuzzy will return an index of highest scored item.
Note that any item identical to keyword will be picked if found any. Also, if not everything
in the keyword has not matched, it will get (or reset to) 0 point.

- First character of keyword matches first character of a candidate.
- Any match in order will get 1 point.
- Consecutive match

Example: if a keyword is `abc` and candidate is `abdc`, then it will be 5 points total.

- `a` from the keyword will get 2 point, (match point + first character point)
- `b` from the keyword will get 2 point, (match point + consecutive point)
- `c` from the keyword will get 1 point. (match point)


## Usage example

~~~go
package main

import (
	"fmt"
	"github.com/orangenumber/amatch"
)

func main() {
	candidates := []string{
		"georgia outdoor network",
		"go network",
		"gon",
		"ago network",
	}
	keyword := "gon"

	for idx, c := range candidates {
		fmt.Printf("%d: %s vs %s --> %d\n", idx, keyword, c, amatch.FuzzyScore(keyword, c))
	}

	best := amatch.Fuzzy(keyword, &candidates)
	fmt.Printf("Best: %s\n", candidates[best])
}
~~~
