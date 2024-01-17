package extract

import (
	"bufio"
	"os"
	"strings"

	"github.com/pericles-luz/go-base/pkg/utils"
)

func MuitoDiferente(a, b uint64) bool {
	if a > b {
		return a-b > 1000
	}
	return b-a > 1000
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

func StringToInt(s string) uint64 {
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "-", "")
	return uint64(utils.StringToInt(s))
}
