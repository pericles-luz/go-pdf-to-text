package extract

import (
	"bufio"
	"os"
)

func MuitoDiferente(a, b uint64) bool {
	if a > b {
		return a-b > 10
	}
	return b-a > 10
}

func ReadLinesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
