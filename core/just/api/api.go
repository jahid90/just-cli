package api

// The API for the `just` module
type JustApi struct {
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
}
