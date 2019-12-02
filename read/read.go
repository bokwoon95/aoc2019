package read

import (
	"bufio"
	"os"
)

// Accumulator ...
type Accumulator interface {
	Append(string) error
}

// Input ...
func Input(filename string, acc Accumulator) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err = acc.Append(scanner.Text())
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}
