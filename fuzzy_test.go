// Copyright 2020 Gon Yi
// https://gonyyi.com/copyright.txt

package amatch_test

import (
	"github.com/orangenumber/amatch"
	"testing"
)

func TestFuzzyScore(t *testing.T) {
	if score := amatch.FuzzyScore("gon", "georgia outdoor network"); score != 4 {
		t.Logf("exp: %d, act: %d", 4, score)
		t.Fail()
	}
	if score := amatch.FuzzyScore("gon", "go network"); score != 5 {
		t.Logf("exp: %d, act: %d", 5, score)
		t.Fail()
	}
	if score := amatch.FuzzyScore("gon", "general party time"); score != 0 {
		t.Logf("exp: %d, act: %d", 0, score)
		t.Fail()
	}
}

func TestFuzzy(t *testing.T) {
	if amatch.Fuzzy("gon", &[]string{"georgia outdoor network", "go network", "general party time", "gon"}) != 3 {
		t.Fail()
	}
	if score := amatch.Fuzzy("gond", &[]string{"georgia outdoor network", "go network", "general party time", "gon"}); score != -1 {
		t.Logf("exp: %d, act: %d", -1, score)
		t.Fail()
	}
}

func BenchmarkFuzzy(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// amatch.Fuzzy("gon", []string{"georgia outdoor network", "go network", "general party time", "gon"})
		amatch.FuzzyScore("gon", "georgia outdoor network")
	}
}
