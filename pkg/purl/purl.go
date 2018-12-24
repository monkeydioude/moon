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
	return p.PathMatcher.Match([]byte(match), []byte(path))
}

func (p *Purl) GetPathMatches() map[string]string {
	return p.PathMatcher.GetMatches()
}

func (p *Purl) MatchQueryString(url string) bool {
	return true
}

func (p *Purl) GetQueryStringMatcheS() map[string]string {
	return p.QueryStringMatcher.GetMatches()
}

func (p *Purl) Match(url string) bool {
	return true
}
