package moon

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/monkeydioude/tools"
)

// Configuration holds service configuration. Typically used for passing configuration container in a file to the Guide
type Configuration map[string]string

// Routes matching URIs go here. Key of the map must hold the Regexp matching an URI
type Routes map[string]*Means

// ResponseHeader is a classic map[string]string header container,
// but explicitly describes that the header will be used for the Response.
// Example: "Access-Control-Allow-Origin" should be passed here
type ResponseHeader map[string]string

// Handler is the core. Contains the Configuration, Response Header and Routes
type Handler struct {
	config  *Configuration
	headers ResponseHeader
	Routes  Routes
}

// Guide kind of acts like Middleware between low levels HandlerFunc and the response, but not exactly a Middleware
type Guide func(*Request, *Configuration) ([]byte, int, error)

// Means defines how and what method shall handle a route
type Means struct {
	Guide  Guide
	Method string
}

// Request contains the wish of an user passed to a Guide
type Request struct {
	Matches     []string
	QueryString map[string]string
	Header      *http.Header
}

// WithHeader specifies headers used in the response
func (h *Handler) WithHeader(k, v string) {
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
func newRequest(m []string, h *http.Header, q map[string]string) *Request {
	return &Request{
		Matches:     m,
		Header:      h,
		QueryString: q,
	}
}

// ParseQueryString parses URI in search of query string
func ParseQueryString(queries string) (qs map[string]string) {
	for _, q := range strings.Split(queries, "&") {
		p := strings.Split(q, "=")
		if len(p) != 2 {
			continue
		}
		qs[p[0]] = p[1]
	}

	return
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

		var q map[string]string

		if len(uri) == 2 {
			q = ParseQueryString(uri[1])
		}

		data, _, err := route.Guide(newRequest(v, &r.Header, q), h.config)
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

// NewHandler generates a Handler from a *Configuration.
func NewHandler(conf *Configuration) *Handler {
	return &Handler{
		config: conf,
		Routes: make(Routes),
	}
}

// Moon Moon ??!1! lol stop it
func Moon(conf *Configuration) *Handler {
	return NewHandler(conf)
}

// Add writes a Means in the Routes map using the regexp that will match the URI, a method and a Guide definition
func (routes *Routes) Add(r, m string, g Guide) {
	(*routes)[r] = &Means{
		Method: m,
		Guide:  g,
	}
}

// AddGet is a wrapper around Add that forces GET method
func (routes *Routes) AddGet(r string, f Guide) {
	routes.Add(r, "GET", f)
}
