package api

// Context The API context
type Context struct {
	KeepGoing bool
}

func NewApiContext() *Context {
	return &Context{
		KeepGoing: false,
	}
}

func (c *Context) WithSkipFailures() {
	c.KeepGoing = true
}
