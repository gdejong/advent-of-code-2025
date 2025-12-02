package must

// NoError takes in some value and an error, and panics if the error is not nil.
func NoError[T any](something T, err error) T {
	if err != nil {
		panic(err)
	}

	return something
}
