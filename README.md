# ASCII Art Justify

A command-line utility that generates ASCII art with text justification options.

## Description

This program extends the basic ASCII art functionality by adding text justification capabilities. It allows aligning the ASCII art output to the left, right, or center of the terminal.

## Features

- All base functionality of the ascii-art program
- Support for text justification (left, right, center)
- Automatic terminal width detection
- Custom width specification

## Usage

```bash
go run . [STRING] [BANNER] [OPTION]
```

Options:
- `--align=left` (default)
- `--align=right`
- `--align=center`
- `--align=justify`

Examples:
```bash
go run . "Hello World" standard --align=right
go run . "Hello World" shadow --align=center
```

## Implementation Details

- Calculates appropriate padding based on terminal width
- Adjusts alignment for multi-line output
- Preserves original banner formatting

## Requirements

- Go 1.13 or higher
