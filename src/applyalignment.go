package src

import "strings"

// Apply alignment to the lines
func ApplyAlignment(lines []string, align string, width int) []string {
	switch align {
	case "center":
		return alignCenter(lines, width)
	case "right":
		return alignRight(lines, width)
	case "justify":
		return alignJustify(lines, width)
	default:
		return alignLeft(lines) // default to left
	}
}

// Alignment functions
func alignLeft(lines []string) []string {
	return lines
}

func alignCenter(lines []string, width int) []string {
	for i, line := range lines {
		spaces := (width - len(line)) / 2
		if spaces > 0 {
			lines[i] = strings.Repeat(" ", spaces) + line
		}
	}
	return lines
}

func alignRight(lines []string, width int) []string {
	for i, line := range lines {
		spaces := width - len(line)
		if spaces > 0 {
			lines[i] = strings.Repeat(" ", spaces) + line
		}
	}
	return lines
}

func alignJustify(lines []string, width int) []string {
	for i, line := range lines {
		// Split each line into individual characters
		characters := strings.Split(line, "")

		// Count the total number of characters and spaces between them
		totalChars := len(strings.Join(characters, ""))
		totalSpaces := width - totalChars

		// If there are no spaces or just one character, continue to the next line
		if totalSpaces <= 0 || len(characters) == 1 {
			continue
		}

		// Calculate the number of spaces to be added between each character
		spacesBetween := totalSpaces / (len(characters) - 1)
		extraSpaces := totalSpaces % (len(characters) - 1) // Extra spaces to be distributed

		// Construct the justified line
		justifiedLine := characters[0]
		for j := 1; j < len(characters); j++ {
			// Add spaces between characters
			spacesToAdd := spacesBetween
			if extraSpaces > 0 {
				spacesToAdd++
				extraSpaces--
			}
			justifiedLine += strings.Repeat(" ", spacesToAdd) + characters[j]
		}

		// Replace the original line with the justified one
		lines[i] = justifiedLine
	}
	return lines
}
