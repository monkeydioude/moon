package moon

// Routes matching URIs go here. Key of the map must hold the Regexp matching an URI
type Routes map[string]*Route

// Route defines how and what method shall handle a route
type Route struct {
	Guide  func(*Request) ([]byte, int, error)
	Method string
	ID     string
}

// Add writes a Route in the Routes map using the regexp that will match the URI, a method and a Guide definition
// Guide type is a callback as such function(*Request) ([]byte, int, error)
func (routes *Routes) Add(r, m string, g func(*Request) ([]byte, int, error)) {
	(*routes)[r] = &Route{
		Method: m,
		Guide:  g,
		ID:     r,
	}
}

// Get is a wrapper around Add that forces GET method
// Guide type is a callback as such function(*Request) ([]byte, int, error)
func (routes *Routes) Get(r string, f func(*Request) ([]byte, int, error)) {
	routes.Add(r, "GET", f)
}

// Post is a wrapper around Add that forces GET method
// Guide type is a callback as such function(*Request) ([]byte, int, error)
func (routes *Routes) Post(r string, f func(*Request) ([]byte, int, error)) {
	routes.Add(r, "POST", f)
}

// NewRoute returns a pointer to Route using its name as ID, a HTTP Method name & a func
func NewRoute(r, m string, f func(*Request) ([]byte, int, error)) *Route {
	return &Route{
		Method: m,
		Guide:  f,
		ID:     r,
	}
}

// Get is a shortcut to NewRoute with forced GET Method
func Get(r string, f func(*Request) ([]byte, int, error)) *Route {
	return NewRoute(r, "GET", f)
}

// Post is a shortcut to NewRoute with forced POST Method
func Post(r string, f func(*Request) ([]byte, int, error)) *Route {
	return NewRoute(r, "POST", f)
}

// Delete is a shortcut to NewRoute with forced DELETE Method
func Delete(r string, f func(*Request) ([]byte, int, error)) *Route {
	return NewRoute(r, "DELETE", f)
}

// Options is a shortcut to NewRoute with forced OPTIONS Method
func Options(r string, f func(*Request) ([]byte, int, error)) *Route {
	return NewRoute(r, "OPTIONS", f)
}
