package purl

import "testing"

func TestICanParseFullUrl(t *testing.T) {
	p := New(&dummyMatcher{}, &dummyMatcher{})
	if !p.MatchPath("/{pouet}/", "/pouet/?king=dadidou") {
		t.Fail()
	}
}

type dummyMatcher struct {
	matches map[string]string
}

func (d *dummyMatcher) Match(t, p []byte) bool {
	return true
}

func (d *dummyMatcher) GetMatches() map[string]string {
	return d.matches
}
