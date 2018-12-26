package purl

import (
	"testing"
)

func TestICanParseFullUrl(t *testing.T) {
	p := New(NewKeyMatcher("{}"))

	if !p.Match("/{pouet}/", "/pouet/?roi=dadidou") ||
		p.GetPathMatches()["pouet"] != "pouet" ||
		p.GetQueryStringMatches().Get("roi") != "dadidou" {
		t.Fail()
	}
}

func TestICanParseUrlWithoutQueryString(t *testing.T) {
	p := New(NewKeyMatcher("{}"))

	if !p.Match("/{isthislove}/", "/thatimfeelin/") ||
		p.GetPathMatches()["isthislove"] != "thatimfeelin" {
		t.Fail()
	}
}
