package identation

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(path string) ([]string, bool) {
	file, err := os.Open(path)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, true
}

func LeadingSpaces(line string) int {
	return len(line) - len(strings.TrimLeft(line, " "))
}

func ReplaceTabs(lines []string) []string {
	var result []string
	for _, l := range lines {
		l = strings.ReplaceAll(l, "\t", "    ")
		result = append(result, l)
	}
	return result
}

func RemoveComments(lines []string) []string {
	var result []string
	for _, l := range lines {
		trimmed := strings.Trim(l, " ")
		if !strings.HasPrefix(trimmed, "/*") && !strings.HasPrefix(trimmed, "*") && !strings.HasPrefix(trimmed, "//") {
			result = append(result, l)
		}
	}
	return result
}

func RemoveEmptyLines(lines []string) []string {
	var result []string
	for _, l := range lines {
		if len(strings.Trim(l, " ")) > 0 {
			result = append(result, l)
		}
	}
	return result
}
