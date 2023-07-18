package http

// Router struct
type Router struct {
	Name string
}

// NewRouter new router
func NewRouter(name string) *Router {
	return &Router{
		Name: name,
	}
}
