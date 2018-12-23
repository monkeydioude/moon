package moon

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/monkeydioude/tools"
)

// ResponseHeader is a classic map[string]string header container,
// but explicitly describes that the header will be used for the Response.
// Example: "Access-Control-Allow-Origin" should be passed here
type ResponseHeader map[string]string

// Handler is the core. Contains the Configuration, Response Header and Routes
type Handler struct {
	headers ResponseHeader
	Routes  Routes
}

// NewHandler generates a Handler.
func NewHandler() *Handler {
	return &Handler{
		Routes: make(Routes),
	}
}

// Moon Moon ??!1! lol stop it
func Moon() *Handler {
	return NewHandler()
}

// Request contains data that should be passed to the function matching a route
// @see (routes *Routes) Add(r, m string, g func(*Request) ([]byte, int, error))
type Request struct {
	HTTPRequest *http.Request
	Matches     []string
	QueryString map[string]string
	Header      *http.Header
}

// AddHeader specifies headers used in the response
func (h *Handler) AddHeader(k, v string) {
	if h.headers == nil {
		h.headers = make(map[string]string)
	}
	h.headers[k] = v
}

func (h *Handler) applyHeaders(rw http.ResponseWriter) {
	for key, value := range h.headers {
		rw.Header().Set(key, value)
	}
}

// newRequest generates a Request from URI parsing & headers. Used in ServeHTTP
func newRequest(m []string, h *http.Header, q map[string]string, r *http.Request) *Request {
	return &Request{
		HTTPRequest: r,
		Matches:     m,
		Header:      h,
		QueryString: q,
	}
}

// ParseQueryString parses URI in search of query string
func ParseQueryString(queries string, qs *map[string]string) {
	for _, q := range strings.Split(queries, "&") {
		p := strings.Split(q, "=")
		if len(p) != 2 {
			continue
		}
		(*qs)[p[0]] = p[1]
	}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}

	h.applyHeaders(rw)

	for p, route := range h.Routes {
		if route.Method != r.Method {
			continue
		}

		uri := strings.Split(r.RequestURI, "?")
		v, err := tools.MatchAndFind(p, strings.Trim(uri[0], "/"))
		if err != nil {
			continue
		}

		q := make(map[string]string)

		if len(uri) == 2 {
			ParseQueryString(uri[1], &q)
		}

		data, _, err := route.Guide(newRequest(v, &r.Header, q, r))
		if err != nil {
			log.Printf("[ERR ] Error while Guiding. Reason: %s\n", err)
			tools.HttpNotFound(rw)
			return
		}

		fmt.Fprint(rw, string(data))
		return
	}
	log.Printf("[WARN] '%s' did not match any route\n", r.RequestURI)
	tools.HttpNotFound(rw)
}

func (h *Handler) MakeRouter(routes ...*Route) {
	for _, r := range routes {
		h.Routes[r.ID] = r
	}
}
