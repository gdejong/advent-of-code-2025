package input

import (
	"bufio"
	"os"
)

func Content(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	lines := make([]string, 0)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}
