### 2.1.2 - 01/01/2019

- Fix case where Purl could segfault on trial string
- Remove whole URL from Purl matches URL has no query string
- Modify test to be a little more complicated
- Modify 404 log message to include HTTP method

### 2.1.1 - 26/12/2018

- Wops forgot fixes... really need tests from 2.1.0 in `handlers` 

### 2.1.0 - 26/12/2018

- Remove query string parser, use std lib one

### 2.0.1 - 24/12/2018

- Add purl tests to travis build
- Remove handler tests
- Now storing routes using Route pattern and Method as identifiers (meaning it is possible to apply multiple handlers to a same route but method has to be different, else handlers overwrite will happen)

### 2.0.0 - 24/12/2018

- Use of Purl inside routing system
- Moon routes are now defined using curly brackets patterns (aka Mustacho patterns). Ex: `/user/{action}/doggo`
- `Matches` received in `moon.Request` is now a `map[string]string`, previously `[]string`
- Mustacho patterns matching a route will now create a key containing its value in `request.Matches`. Ex: `/user/{action}/doggo` => `/user/good/doggo` => `request.Matches["action"] == "good`

### 1.6.1 - 23/12/2018

- Removing "/" trimming in URL matching

### 1.6.0 - 23/12/2018

- Split router methods
- Add conveniant way to create a router

### 1.5.0 - 19/12/2018

- Remove Configuration from handler func

### 1.4.0 - 04/10/2018

- Add example
- Fix bug in ParseQueryString and add tests for it
- Better README.md
- Thanks @scolalongo :)

### 1.3.0 - 03/08/2018

- Remove Guide type, makes things more complicated than anything else

### 1.0.1 - 10/06/2018

- Better logging in ServeHTTP.

### 1.0.0 - 09/06/2018

- Oh god it is born.
- Wrapper for http.Handler, making it way way easier to declare and handles route
