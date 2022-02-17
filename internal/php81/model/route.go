package model

type Route struct {
	Name   string
	Method string
	URL    string
}

func NewRoute(name, method, url string) *Route {
	return &Route{
		Name:   name,
		Method: method,
		URL:    url,
	}
}
