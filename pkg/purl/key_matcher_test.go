package purl

import (
	"testing"
)

func TestKeyMatcherCanMatchSimpleString(t *testing.T) {
	trial := "/a/b/c/d/"
	url := "/a/b/c/d/"
	p := NewKeyMatcher("{}")

	if !p.Match([]byte(url), []byte(trial)) {
		t.Fail()
	}
}
func TestKeyMatcherCanMatch(t *testing.T) {
	trial := "/test/{wesh}/3/{salut}/{nada}//{}}{/{}/}{/}}}}}{/}}}{{{{{{"
	url := "/test/alors/3/ccool//2/3/4"
	p := NewKeyMatcher("{}")

	p.Match([]byte(url), []byte(trial))
	m := p.GetMatches()

	if len(m) != 3 {
		t.Fail()
	}
	if v, ok := m["wesh"]; !ok || v != "alors" {
		t.Fail()
	}
	if v, ok := m["salut"]; !ok || v != "ccool" {
		t.Fail()
	}
	if v, ok := m["nada"]; !ok || v != "" {
		t.Fail()
	}
}

func TestKeyMatcherFailOnNonMatchingString(t *testing.T) {
	trial := "/b/a/c/d/"
	url := "/a/b/c/d/"
	p := NewKeyMatcher("{}")

	if p.Match([]byte(trial), []byte(url)) {
		t.Fail()
	}
}
