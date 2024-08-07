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

	if !p.Match("/{isthislove}/isthislove", "/thatimfeelin/isthislove") ||
		p.GetPathMatches()["isthislove"] != "thatimfeelin" {
		t.Fail()
	}
}

func TestIShouldFailParsingPartialChunk(t *testing.T) {
	p := New(NewKeyMatcher("{}"))

	if p.Match("/{should}", "/notmatch/salut") {
		t.Fail()
	}
}
