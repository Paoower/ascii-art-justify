# Ascii Art Justify

It's a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.
An optional arguments is the to change the alignment of the output be possible to use a flag --align=<type>, in which type can be center,left,right,justify.

## **[main.go](main.go)**

First, arguments are being handled, and saved correctly. Each line of the input string is turned into ASCII art character by character.

## **[src/getword.go](src/getword.go)**

The ```GetWord``` function takes a string as argument and applies the ```GetLetter``` function to each character. It returns a slice of string representing each line to display.

## **[src/getletter.go](src/getletter.go)**

The ```GetLetter``` function takes an int as argument repesenting the ASCII code of the character that needs to be turned into ASCII art. Its position in the banner file is calculated. A string is returned containing the ASCII art for the provided character.

## **[src/getbanner.go](src/getbanner.go)**

The ```GetBanner``` function takes a string as an argument representing the desired banner style. It returns the file path of the corresponding banner file. If the style provided is not recognized, an error is returned indicating the available styles.