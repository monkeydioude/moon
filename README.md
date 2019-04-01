## Moon is a wrapper for http.Handler, making it way way easier to declare and handles route
[![Build Status](https://travis-ci.org/monkeydioude/moon.svg?branch=master)](https://travis-ci.org/monkeydioude/moon)

Simple example:
```golang
	func main() {
		server := moon.Moon()
		// Me API es su API
		server.AddHeader("Access-Control-Allow-Origin", "*")

		server.MakeRouter(
			// Will call routeHandler1() func every time a GET on "/handler1" URI is caught
			moon.Get("/handler1", routeHandler1),
			// This will match any "/handler1/*" GET request and create a "param1" parameter 
			// in moon.Request.Matches containing the matched value by the {param1} pattern
			moon.Get("/handler1/{param1}", routeHandler1bis),
		)

		moon.ServerRun(":8080", server)
	}
```

### What it does
- Define routes using Patterns and a HTTP method
- Not use Regexp for Route Pattern matching
- Give matched parts of the Route Pattern as well as any query strings inside the `moon.Request` 

### Handlers
```golang
	func aHandler(r *moon.Request) ([]byte, int, error)
```

Take:
- a `*moon.Request` containing a simplified HTTP Request (Parts of the uri matched, query string & headers)

Return:
- `[]byte` containing the text response
- `int`, as the http status code
- `error`, as an error, for logging purpose



### [Detailed example(s) available here](./examples)

See [Moon's GoDoc](https://godoc.org/github.com/monkeydioude/moon) for Documentation.

This project still lacks:
- Tests (ongoing)
- Benchs
- ~~Matching without using RegExp ?~~
- Response Caching ?
