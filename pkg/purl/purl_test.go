package purl

import (
	"testing"
)

func TestICanParseFullUrl(t *testing.T) {
	p := New(NewKeyMatcher("{}"), NewQueryStringMatcher())

	if !p.Match("/{pouet}/", "/pouet/?roi=dadidou") ||
		p.GetPathMatches()["pouet"] != "pouet" ||
		p.GetQueryStringMatches()["roi"] != "dadidou" {
		t.Fail()
	}
}

func TestICanParseUrlWithoutQueryString(t *testing.T) {
	p := New(NewKeyMatcher("{}"), NewQueryStringMatcher())

	if !p.Match("/{isthislove}/", "/thatimfeelin/") ||
		p.GetPathMatches()["isthislove"] != "thatimfeelin" ||
		len(p.GetQueryStringMatches()) > 0 {
		t.Fail()
	}
}
