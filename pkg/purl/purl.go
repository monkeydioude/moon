package purl

import "strings"

type Purl struct {
	PathMatcher        Matcher
	QueryStringMatcher Matcher
}

func New(path Matcher, queryString Matcher) *Purl {
	return &Purl{
		PathMatcher:        path,
		QueryStringMatcher: queryString,
	}
}

type Matcher interface {
	Match([]byte, []byte) bool
	GetMatches() map[string]string
}

func (p *Purl) MatchPath(match, url string) bool {
	index := strings.IndexByte(url, '?')
	if index == -1 {
		index = len(url)
	}
	path := url[0:index]
	return p.PathMatcher.Match([]byte(path), []byte(match))
}

func (p *Purl) GetPathMatches() map[string]string {
	return p.PathMatcher.GetMatches()
}

func (p *Purl) MatchQueryString(url string) bool {
	index := strings.IndexByte(url, '?')
	if index == -1 {
		index = len(url)
	}
	q := url[index:]
	return p.QueryStringMatcher.Match([]byte(q), nil)
}

func (p *Purl) GetQueryStringMatches() map[string]string {
	return p.QueryStringMatcher.GetMatches()
}

func (p *Purl) Match(match, url string) bool {
	return p.MatchPath(match, url) && p.MatchQueryString(url)
}

func NewUrlParser() *Purl {
	return New(
		NewKeyMatcher("{}"),
		NewQueryStringMatcher(),
	)
}
