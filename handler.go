package moon

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/monkeydioude/moon/pkg/purl"
	"github.com/monkeydioude/tools"
)

// ResponseHeader is a classic map[string]string header container,
// but explicitly describes that the header will be used for the Response.
// Example: "Access-Control-Allow-Origin" should be passed here
type ResponseHeader map[string]string

// Handler is the core. Contains the Configuration, Response Header and Routes
type Handler struct {
	headers ResponseHeader
	Routes  map[string]Routes
}

// NewHandler generates a Handler.
func NewHandler() *Handler {
	return &Handler{
		Routes: make(map[string]Routes),
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
	Matches     map[string]string
	QueryString url.Values
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
func newRequest(m map[string]string, h *http.Header, q url.Values, r *http.Request) *Request {
	return &Request{
		HTTPRequest: r,
		Matches:     m,
		Header:      h,
		QueryString: q,
	}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}

	h.applyHeaders(rw)

	for p, mRoute := range h.Routes {
		if _, ok := mRoute[r.Method]; !ok {
			continue
		}

		route := mRoute[r.Method]
		parser := purl.NewUrlParser()
		if !parser.Match(p, r.RequestURI) {
			continue
		}

		data, _, err := route.Guide(
			newRequest(
				parser.GetPathMatches(),
				&r.Header,
				parser.GetQueryStringMatches(),
				r,
			),
		)
		if err != nil {
			log.Printf("[ERR ] Error while Guiding. Reason: %s\n", err)
			tools.HttpNotFound(rw)
			return
		}

		fmt.Fprint(rw, string(data))
		return
	}
	log.Printf("[WARN] %s '%s' did not match any route\n", r.Method, r.RequestURI)
	tools.HttpNotFound(rw)
}

func (h *Handler) MakeRouter(routes ...*Route) {
	for _, r := range routes {
		if _, ok := h.Routes[r.ID]; !ok {
			h.Routes[r.ID] = make(Routes)
		}
		h.Routes[r.ID][r.Method] = r
	}
}
