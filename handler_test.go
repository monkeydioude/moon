package moon

import "testing"

func TestICanParseSimpleQueries(t *testing.T) {
	qs := make(map[string]string)
	ParseQueryString("1=1&2=2", &qs)

	if qs["1"] != "1" || qs["2"] != "2" {
		t.Error("Did not parse query string as intended")
	}
}

func TestIDoNotParseOnEmptyQueries(t *testing.T) {
	qs := make(map[string]string)
	ParseQueryString("", &qs)

	if len(qs) > 0 {
		t.Error("ParseQueryString's second parameter should have stayed empty")
	}
}
