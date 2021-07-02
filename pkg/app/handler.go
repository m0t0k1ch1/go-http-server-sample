package app

// HandlerFunc is a function to serve HTTP requests.
type HandlerFunc func(env *Env, c *Context) error
