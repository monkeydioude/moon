package purl

import "bytes"

const (
	queryStringTokens = "=&"
	keyMode           = 0
	valueMode         = 1
	mMode             = -1
)

// QueryStringMatcher implements purl.Matcher interface
// and is made for matching and retrieving query string
// parameters from an URL
type QueryStringMatcher struct {
	matches map[string]string
}

// NewQueryStringMatcher returns a pointer to QueryStringMatcher
// after initializing its map
func NewQueryStringMatcher() *QueryStringMatcher {
	return &QueryStringMatcher{
		matches: make(map[string]string),
	}
}

// writeMatch write a found match using pointers to key & value buffers
func (q *QueryStringMatcher) writeMatch(k *bytes.Buffer, v *bytes.Buffer) {
	if k.Len() > 0 {
		q.matches[k.String()] = v.String()
	}
	k.Reset()
	v.Reset()
}

// Match implements purl.Matcher interface. "match" is never used here
func (q *QueryStringMatcher) Match(trial, match []byte) bool {
	mode := keyMode
	buffers := map[int]*bytes.Buffer{
		keyMode:   &bytes.Buffer{},
		valueMode: &bytes.Buffer{},
	}

	c := bytes.Index(trial, []byte{'#'})
	if c != -1 {
		trial = trial[0:c]
	}

	trial = bytes.Trim(trial, "?&")

	for _, v := range trial {
		// Happily stocking keys (after a &) and values (after a =).
		// Using an automatic switch mode, not really sure about this yet,
		// works well for now.
		if v == queryStringTokens[mode] {
			if mode == valueMode {
				q.writeMatch(buffers[keyMode], buffers[valueMode])
			}
			mode = (mode + mMode) * mMode
			continue
		}
		// case of a wild & appearing before a =
		if mode == keyMode && v == '&' {
			q.writeMatch(buffers[keyMode], buffers[valueMode])
			continue
		}
		buffers[mode].WriteByte(v)
	}

	if buffers[keyMode].Len() > 0 {
		q.matches[buffers[keyMode].String()] = buffers[valueMode].String()
	}
	return true
}

// GetMatches implements purl.Matcher interface
func (q *QueryStringMatcher) GetMatches() map[string]string {
	return q.matches
}
