## Moon is a wrapper for http.Handler, making it way way easier to declare and handles route
[![Build Status](https://travis-ci.org/monkeydioude/moon.svg?branch=master)](https://travis-ci.org/monkeydioude/moon)

Example:
```golang
	handler := moon.NewHandler(configuration)
	// Me API es su API
	handler.WithHeader("Access-Control-Allow-Origin", "*")

    	// Will call allMoons() func every time a GET on "/moon/all" URI is caught
	handler.Routes.AddGet("moon/all", allMoons)
    	// Will call lolStop() func every time a GET on "/moon/moon" URI is caught
	handler.Routes.AddGet("moon/moon", lolStop)

    	// Start standard http.Server
	server := &http.Server{
		Addr:           ":9393:",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("[INFO] Starting server")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
```

See [Moon's GoDoc](https://godoc.org/github.com/monkeydioude/moon) for Documentation.

Lacks:
- Tests
- Benchs
- Matching without using RegExp ?
- Response Caching ?
- More Lazy-Win funcs ?
