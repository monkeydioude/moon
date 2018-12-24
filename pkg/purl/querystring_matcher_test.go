package purl

import (
	"testing"
)

func TestICanParseQueryString(t *testing.T) {
	q := NewQueryStringMatcher()
	trial := "?wesh=alors&&salut=ccool&dora&&#&no=nono"
	q.Match([]byte(trial), nil)

	res := q.GetMatches()

	if len(res) != 3 || res["wesh"] != "alors" || res["salut"] != "ccool" || res["dora"] != "" {
		t.Fail()
	}
}
