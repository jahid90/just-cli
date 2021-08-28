package lib

// Ellipsify Ellipsifies the provided string to the provided length
// Returns the original string if it fits
func Ellipsify(input string, length int) string {
	if len(input) <= length {
		return input
	}

	// the length is reduced by 4 to provision for the ellipses ... and a space
	truncated := input[:length-4]
	truncated = truncated + " ..."

	return truncated
}
