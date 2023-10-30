package main

import (
	"asciirev/asciirev"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input := os.Args[1:]

	banner := ""
	filename := ""
	inputStr := ""

	reg1 := regexp.MustCompile(`--reverse=\s*([^"]*)`)
	reg2 := regexp.MustCompile(`--output=\s*([^"]*)`)
	regB := regexp.MustCompile(`standard|shadow|thinkertoy`)

	if len(input) == 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if reg1.MatchString(input[0]) {
		filename = input[0][10:]

		asciirev.Reverse(filename)
		return
	} else if reg2.MatchString(input[0]) {
		if len(input) != 3 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
		filename = input[0][9:]
		inputStr = input[1]
		banner = input[2]

		if !asciirev.CheckFileHashing(banner) {
			fmt.Println("ERROR: the file was changed!")
			return
		}

		if !strings.HasSuffix(filename, ".txt") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
			return
		}

		if !regB.MatchString(banner) {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
			return
		}

		asciirev.Output(filename, inputStr, banner)
		return
	} else if !reg2.MatchString(input[0]) && !reg1.MatchString(input[0]) {
		fmt.Println("Usage: go run . [OPTION]\nEX: go run . --reverse=<fileName>\n\nOR\n\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	if len(input) == 1 {

		inputStr = input[0]
		banner = "standard"

	} else if len(input) == 2 {
		inputStr = input[0]
		banner = input[1]
	}
	fmt.Println(banner)
	if !regB.MatchString(banner) {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}

	asciirev.FS(inputStr, banner)
}

// 9hpbd
