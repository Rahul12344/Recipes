package expand

import "testing"

/* UNIT TESTS FOR EXPAND */

func MapperTest(t *testing.T) {
	mapper := NewMapper("test")
	mapper.InitMap()

	if mapper.Map != "t.*e.*s.*t" {
		t.Errorf("Invalid map: expected: %s; got: %s", "t.*e.*s.*t", mapper.Map)
	}
}

func MatcherTest(t *testing.T) {
	var dataset []string
	matcher := NewMatcher("test", "tst", dataset...)
	matcher.Matcher()
}

func AssocationsTest(t *testing.T) {
	var dataset []string
	var alphabet []string

	threshold := 0.8
	association := NewAssociations(alphabet, dataset)
	if association.Associate(threshold) != 1 {
		t.Errorf("Associate failed")
	}

	threshold = 0.0
	association = NewAssociations(alphabet, dataset)
	if association.Associate(threshold) != 1 {
		t.Errorf("Associate failed")
	}
}
