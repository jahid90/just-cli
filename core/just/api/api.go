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

func (c *Context) EnableSkipFailures() {
	c.KeepGoing = true
}

// JustAPI The API for the `just` module
type JustApi struct {
	// The API context
	Ctx *Context

	// Function to get the version of the config
	Version func() string

	// Function to format the output in the supported formats
	// The param determines the format
	Format func(string) ([]byte, error)

	// Function to show a listing of the config
	// The param determines whether long or short format is returned
	ShowListing func(bool) ([]byte, error)

	// Function to execute a provided alias
	Execute func(string) error

	// Function to show the command(s) corresponding to an alias
	ShowCommand func(string) (string, error)
}
