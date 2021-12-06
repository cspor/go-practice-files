package errorHandler

// Check Checks an error and panics if it's not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
