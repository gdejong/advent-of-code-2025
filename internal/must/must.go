package must

import "strconv"

// NoError takes in some value and an error, and panics if the error is not nil.
func NoError[T any](something T, err error) T {
	if err != nil {
		panic(err)
	}

	return something
}

func Int(s string) int {
	return NoError(strconv.Atoi(s))
}
