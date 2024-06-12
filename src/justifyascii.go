package src

import (
	"os"
	"strings"
)

const LETTER_HEIGHT = 8

func GetJustifiedAscii(input, style string, spacesRemaining, numOfSpaces int) []string {
	bannerFile := ""
	switch style {
	case "standard":
		bannerFile = "banners/standard.txt"
	case "shadow":
		bannerFile = "banners/shadow.txt"
	case "thinkertoy":
		bannerFile = "banners/thinkertoy.txt"
	default:
		bannerFile = ""
	}

	lines := make([]string, 0)
	inputLines := strings.Split(input, "\\n")

	for _, line := range inputLines {
		if line == "" {
			lines = append(lines, "")
			continue
		}
		linesOfline := getLine(line, bannerFile, spacesRemaining, numOfSpaces)
		lines = append(lines, linesOfline...)
	}
	return lines
}

func getLine(input, bannerFile string, spacesRemaining, numOfSpaces int) []string {
	lines := make([]string, 8)
	f, err := os.ReadFile(bannerFile)
	if err != nil {
		return []string{}
	}

	content := strings.ReplaceAll(string(f), "\r\n", "\n")
	for _, char := range input {
		if char == ' ' {
			for x := 0; x < spacesRemaining/numOfSpaces; x++ {
				for i := 0; i < len(lines); i++ {
					lines[i] += " "
				}
			}
			continue
		}
		c := strings.Split(GetLetter(content, int(char)), "\n")
		for i := 0; i < len(lines); i++ {
			lines[i] += c[i]
		}
	}
	return lines
}
