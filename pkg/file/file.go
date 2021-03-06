package file

import (
	"bufio"
	"os"
)

func GetEachLine(path string) ([]string, error) {
	lines := []string{}
	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, err
}
