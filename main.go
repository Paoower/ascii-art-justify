package main

import (
	src "ascii-art-justify/src"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}
	var input, style, align string
	style = "standard"
	align = "left"
	argIndex := 1
	for argIndex < len(os.Args) {
		arg := os.Args[argIndex]
		switch {
		case strings.HasPrefix(arg, "--align="):
			align = strings.TrimPrefix(arg, "--align=")
			if align != "left" && align != "center" && align != "right" && align != "justify" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
				return
			}
		case input == "":
			input = arg
		case style == "standard":
			style = arg
		default:
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			return
		}
		argIndex++
	}
	if input == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}
	// Use correct Banner files
	bannerFile, err := src.GetBannerFile(style)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	lines := make([]string, 0)
	words := strings.Split(input, "\\n")
	if align == "justify" {
		numOfSpaces := strings.Count(input, " ")
		for _, word := range words {
			if word == "" {
				lines = append(lines, "")
			} else {
				getW, err := src.GetWord(Trim(word), bannerFile)
				if err != nil {
					fmt.Println(err)
				}
				lines = append(lines, getW...)
			}
		}
		lenOfOneLine := len(lines[0])
		spacesToRemain := getTerminalWidth() - lenOfOneLine
		lines = []string{}
		lines = append(lines, src.GetJustifiedAscii(input, style, spacesToRemain, numOfSpaces)...)
	}
	if align == "center" {
		for _, word := range words {
			if word == "" {
				lines = append(lines, "")
			} else {
				getW, err := src.GetWord(word, bannerFile)
				if err != nil {
					fmt.Println(err)
				}
				lines = append(lines, getW...)
			}
		}
		lenOfOneLine := len(lines[0])
		spacesToRemain := getTerminalWidth() - lenOfOneLine
		lines = []string{}
		lines = append(lines, src.GetCenteredAscii(input, style, spacesToRemain, 2)...)
	}
	if align == "right" {
		for _, word := range words {
			if word == "" {
				lines = append(lines, "")
			} else {
				getW, err := src.GetWord(word, bannerFile)
				if err != nil {
					fmt.Println(err)
				}
				lines = append(lines, getW...)
			}
		}
		lenOfOneLine := len(lines[0])
		spacesToRemain := getTerminalWidth() - lenOfOneLine
		lines = []string{}
		lines = append(lines, src.GetCenteredAscii(input, style, spacesToRemain, 1)...)
	}
	if align == "left" {
		for _, word := range words {
			if word == "" {
				lines = append(lines, "")
			} else {
				getW, err := src.GetWord(word, bannerFile)
				if err != nil {
					fmt.Println(err)
				}
				lines = append(lines, getW...)
			}
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Get the width of the terminal
func getTerminalWidth() int {
	var dimensions [4]uint16
	retCode, _, errno := syscall.Syscall6(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&dimensions)),
		0, 0, 0)
	if int(retCode) == -1 {
		fmt.Println("Error getting terminal size:", errno)
		return 80 // default width
	}
	return int(dimensions[1])
}
func Trim(s string) string {
	str := ""
	for _, char := range s {
		if char != ' ' {
			str += string(char)
		}
	}
	return str
}
