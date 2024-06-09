package bufiofiles

import (
	"bufio"
	"fmt"
	"os"
)

// Read lines from a file
func ReadLines(filename string) (out []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out, err
}

// example for writing data to file
func writeLineExample(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	writer.Flush()

	return nil
}
