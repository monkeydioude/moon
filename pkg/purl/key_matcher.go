package purl

import (
	"bytes"
)

type KeyMatcher struct {
	pattern *Pattern
	matches map[string]string
}

func NewKeyMatcher(pattern string) *KeyMatcher {
	return &KeyMatcher{
		pattern: NewPattern(pattern),
		matches: make(map[string]string),
	}
}

func (m *KeyMatcher) Match(trial, match []byte) bool {
	key := &bytes.Buffer{}
	value := &bytes.Buffer{}
	it := 0
	im := 0
	lm := len(match)
	lt := len(trial)
	for im < lm {
		if !m.pattern.Match(match[im]) {
			if it < lt && match[im] != trial[it] {
				return false
			}
			it++
			im++
			continue
		}

		// Match the Pattern segment against the "match" string to find
		// the parameter name.
		// Ex, if the defined pattern is "{}", in /user/{wesh},
		// it will try to extract "wesh"
		for !m.pattern.IsComplete() && im < lm {
			if m.pattern.Match(match[im]) {
				m.pattern.Next()
			} else {
				key.WriteByte(match[im])
			}
			im++
		}

		// Could not find any param name
		if key.Len() == 0 {
			continue
		}

		if im == lm {
			im--
		}
		// Recovers the value from "trial", until it meets the same
		// character (:< ghetto) right after the end of the pattern
		// in "match"
		for it < len(trial) && trial[it] != match[im] && trial[it] != '/' {
			value.WriteByte(trial[it])
			it++
		}

		m.matches[key.String()] = value.String()
		key.Reset()
		value.Reset()
		m.pattern.Reset()
	}

	fLen := lm
	for key, value := range m.matches {
		fLen = fLen - len(m.pattern.p) - len(key) + len(value)
	}

	return fLen == lt
}

func (m *KeyMatcher) GetMatches() map[string]string {
	return m.matches
}

type Pattern struct {
	p  []byte
	it int
}

func NewPattern(p string) *Pattern {
	return &Pattern{
		p: []byte(p),
	}
}

func (p *Pattern) Match(c byte) (res bool) {
	return p.it < len(p.p) && p.p[p.it] == c
}

func (p *Pattern) Next() {
	p.it++
}

func (p *Pattern) IsComplete() bool {
	return p.it == len(p.p)
}

func (p *Pattern) IsMatching() bool {
	return p.it > 0
}

func (p *Pattern) Reset() {
	p.it = 0
}
