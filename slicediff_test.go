package slicediff

import (
	"strings"
	"testing"
)

var emptySlice = make([]string, 0)

func testScenario(init []string, updates [][]string, expectedAdd []string, expectedDel []string, t *testing.T) {
	sd := New()
	sd.Append(init)
	var add []string
	var del []string
	for _, update := range updates {
		add, del = sd.SortedDiff(update)
	}
	if !sliceEq(add, expectedAdd) {
		t.Error("add != expectedAdd", add, expectedAdd)
	}
	if !sliceEq(del, expectedDel) {
		t.Error("del != expectedDel", del, expectedDel)
	}
}

func testCharScenario(cInit string, cUpdates []string, cExpectedAdd string, cExpectedDel string, t *testing.T) {
	init := strings.Split(cInit, "")
	updates := extractCSlices(cUpdates)
	expectedAdd := strings.Split(cExpectedAdd, "")
	expectedDel := strings.Split(cExpectedDel, "")
	testScenario(init, updates, expectedAdd, expectedDel, t)
}
func extractCSlices(s []string) [][]string {
	new := make([][]string, 0, len(s))
	for _, c := range s {
		new = append(new, strings.Split(c, ""))
	}
	return new
}

func TestEmptyInitScenarios(t *testing.T) {
	testCharScenario(
		"",
		emptySlice,
		"",
		"",
		t,
	)

	testCharScenario(
		"",
		[]string{"abd"},
		"abd",
		"",
		t,
	)

	testCharScenario(
		"",
		[]string{"abd", ""},
		"",
		"abd",
		t,
	)
}

func TestSmallInitScenarios(t *testing.T) {
	testCharScenario(
		"abd",
		emptySlice,
		"",
		"",
		t,
	)
	testCharScenario(
		"abd",
		[]string{"abd"},
		"",
		"",
		t,
	)
}

func TestSmallInitDeletionScenarios(t *testing.T) {
	testCharScenario(
		"abd",
		[]string{""},
		"",
		"abd",
		t,
	)
	testCharScenario(
		"abd",
		[]string{"a"},
		"",
		"bd",
		t,
	)
	testCharScenario(
		"abd",
		[]string{"b"},
		"",
		"ad",
		t,
	)
	testCharScenario(
		"abd",
		[]string{"d"},
		"",
		"ab",
		t,
	)
	testCharScenario(
		"abd",
		[]string{"bd"},
		"",
		"a",
		t,
	)
	testCharScenario(
		"abd",
		[]string{"ad"},
		"",
		"b",
		t,
	)
	testCharScenario(
		"abd",
		[]string{"ab"},
		"",
		"d",
		t,
	)
}

func TestSmallInitAdditionScenarios(t *testing.T) {
	testCharScenario(
		"abd",
		[]string{"abcd"},
		"c",
		"",
		t,
	)
	testCharScenario(
		"bcd",
		[]string{"abcd"},
		"a",
		"",
		t,
	)
	testCharScenario(
		"abc",
		[]string{"abcd"},
		"d",
		"",
		t,
	)

	testCharScenario(
		"ab",
		[]string{"abcd"},
		"cd",
		"",
		t,
	)
	testCharScenario(
		"bc",
		[]string{"abcd"},
		"ad",
		"",
		t,
	)
	testCharScenario(
		"cd",
		[]string{"abcd"},
		"ab",
		"",
		t,
	)
	testCharScenario(
		"ac",
		[]string{"abcd"},
		"bd",
		"",
		t,
	)
	testCharScenario(
		"bd",
		[]string{"abcd"},
		"ac",
		"",
		t,
	)
	testCharScenario(
		"ad",
		[]string{"abcd"},
		"bc",
		"",
		t,
	)

	testCharScenario(
		"a",
		[]string{"abcd"},
		"bcd",
		"",
		t,
	)
	testCharScenario(
		"b",
		[]string{"abcd"},
		"acd",
		"",
		t,
	)
	testCharScenario(
		"c",
		[]string{"abcd"},
		"abd",
		"",
		t,
	)
	testCharScenario(
		"d",
		[]string{"abcd"},
		"abc",
		"",
		t,
	)
}

func sliceEq(a, b []string) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
