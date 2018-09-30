## Moon is a wrapper for http.Handler, making it way way easier to declare and handles route
[![Build Status](https://travis-ci.org/monkeydioude/moon.svg?branch=master)](https://travis-ci.org/monkeydioude/moon)

Simple example:
```golang
	func routeHandler1(r *moon.Request, c *moon.Configuration) ([]byte, int, error) {
		return []byte("hello"), 200, nil
	}

	func main() {
		h := moon.Moon(nil)
		// Me API es su API
		h.WithHeader("Access-Control-Allow-Origin", "*")

		// Will call routeHandler1() func every time a GET on "/route1" URI is caught
		h.Routes.AddGet("route1", routeHandler1)

		// Server will run localhost port 8080
		moon.ServerRun(":8080", h)
	}
```

### Handlers
Take:
- a `*moon.Request` containing a simplified HTTP Request (Parts of the uri matched, query string & headers)
- a `*moon.Configuration` basically a map[string]string containing the data passed to Moon (`config` in the example above)

Return:
- `[]byte` containing the text response
- `int`, as the http status code
- `error`, as an error, for logging purpose

### [Detailed example(s) available here](./examples)

See [Moon's GoDoc](https://godoc.org/github.com/monkeydioude/moon) for Documentation.

This project still lacks:
- Tests
- Benchs
- Matching without using RegExp ?
- Response Caching ?
