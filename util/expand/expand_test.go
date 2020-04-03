package expand

import "testing"

/* UNIT TESTS FOR EXPAND */

func RgxMprTest(t *testing.T) {
	rgxMpr := NewRgxMpr("test")
	rgxMpr.InitMap(rgxMpr.Rgxmpr)

	if rgxMpr.Rgxmpr != "t.*e.*s.*t" {
		t.Errorf("Invalid map: expected: %s; got: %s", "t.*e.*s.*t", rgxMpr.Rgxmpr)
	}
}

func MatcherTest(t *testing.T) {
	matcher := NewMatcher("test", "tst", nil, nil)
	matcher.Matcher()
}

func AssocationsTest(t *testing.T) {

}
