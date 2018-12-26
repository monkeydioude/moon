package purl

import (
	"net/url"
	"strings"
)

type Purl struct {
	PathMatcher       Matcher
	QueryStringValues url.Values
}

func New(path Matcher) *Purl {
	return &Purl{
		PathMatcher: path,
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

func (p *Purl) MatchQueryString(u string) bool {
	v, err := url.ParseQuery(u[strings.IndexByte(u, '?')+1:])

	if err != nil {
		return false
	}

	p.QueryStringValues = v
	return true
}

func (p *Purl) GetQueryStringMatches() url.Values {
	return p.QueryStringValues
}

func (p *Purl) Match(match, u string) bool {
	return p.MatchPath(match, u) && p.MatchQueryString(u)
}

func NewUrlParser() *Purl {
	return New(
		NewKeyMatcher("{}"),
	)
}
