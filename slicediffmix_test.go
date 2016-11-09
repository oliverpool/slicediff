package slicediff

import "testing"

func TestMixedScenarios(t *testing.T) {
	testCharScenario(
		"abd",
		[]string{"abc"},
		"c",
		"d",
		t,
	)
	testCharScenario(
		"abd",
		[]string{"bc"},
		"c",
		"ad",
		t,
	)
}
