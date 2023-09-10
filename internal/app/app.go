package app

import (
	"github.com/RickDred/ascii-art/internal/art"
	"github.com/RickDred/ascii-art/internal/validator"
	"github.com/RickDred/ascii-art/pkg/console"
	"github.com/RickDred/ascii-art/pkg/files"
	"flag"
	"fmt"
)

type Font struct {
	path    string
	md5Hash string
}

const (
	usageExample = "Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>"
)

func Run() {
	// fonts
	fonts := map[string]Font{
		"standard": {
			path:    "./fonts/standard.txt",
			md5Hash: "ac85e83127e49ec42487f272d9b9db8b",
		},
		"shadow": {
			path:    "./fonts/shadow.txt",
			md5Hash: "a49d5fcb0d5c59b2e77674aa3ab8bbb1",
		},
		"thinkertoy": {
			path:    "./fonts/thinkertoy.txt",
			md5Hash: "86d9947457f6a41a18cb98427e314ff8",
		},
	}

	// validator
	v := validator.New()

	// default
	text := ""
	banner := "standard"
	colorize := ""

	// available options
	availableFonts := []string{"standard", "shadow", "thinkertoy"}
	availableAlignments := []string{"center", "left", "right", "justify"}

	// flags
	align := flag.String("align", "left", "it should be string")
	output := flag.String("output", "", "it should be filename")
	color := flag.String("color", "white", "it should be string")
	reverse := flag.String("reverse", "", "is should be filename")

	flag.Parse()

	if isFlagPassed("output") && (isFlagPassed("align") || isFlagPassed("color")) {
		fmt.Println("error: cannot use output option with align or color flags at the same time")
		return
	}

	args := flag.Args()

	isFlagPas := make(map[string]bool)
	isFlagPas["reverse"] = isFlagPassed("reverse")
	isFlagPas["output"] = isFlagPassed("output")
	isFlagPas["color"] = isFlagPassed("color")
	isFlagPas["align"] = isFlagPassed("align")

	switch len(args) {
	case 0:
		if !isFlagPas["reverse"] {
			fmt.Println(usageExample)
			return
		}
	case 1:
		switch {
		case isFlagPas["reverse"] && isFlagPas["color"]:
			colorize = args[0]
		case isFlagPas["reverse"]:
			banner = args[0]
		default:
			text = args[0]
		}
	case 2:
		switch {
		case isFlagPas["reverse"] && isFlagPas["color"]:
			colorize = args[0]
			banner = args[1]
		case isFlagPas["reverse"]:
			fmt.Println(usageExample)
			return
		case isFlagPas["color"]:
			colorize = args[0]
			text = args[1]
		default:
			text = args[0]
			banner = args[1]
		}
	case 3:
		switch {
		case isFlagPas["reverse"] || !isFlagPas["color"]:
			fmt.Println(usageExample)
			return
		default:
			colorize = args[0]
			text = args[1]
			banner = args[2]
		}
	default:
		fmt.Println(usageExample)
		return
	}

	if ok := v.IsTextAscii(text); !ok {
		fmt.Println("text is not valid, use characters only in Ascii")
		return
	}

	if ok := v.IsArrayContainsString(availableAlignments, *align); !ok {
		fmt.Printf("alignment option is not appropriate, you can use only: %v\n", availableAlignments)
		return
	}

	if ok := v.IsArrayContainsString(availableFonts, banner); !ok {
		fmt.Printf("banner is not appropriate, you can use only: %v\n", availableFonts)
		return
	}

	path := fonts[banner].path
	hash := fonts[banner].md5Hash

	if ok := v.IsMd5HashEqual(path, hash); !ok {
		fmt.Println("File with font was changed")
		return
	}

	if isFlagPassed("color") {
		if _, ok := art.ConvertToRGB(*color); !ok {
			fmt.Println("invalid color input")
			return
		}
	}

	if colorize == "" {
		colorize = text
	}

	var f func(string, string, string, string, string, func() int) (string, bool)

	var funcGetWidth func() int = console.GetWidth

	if isFlagPas["reverse"] {
		text = *reverse
		f = art.ReverseAsciiArt
	} else {
		f = art.AsciiArt
	}

	if isFlagPas["output"] {
		funcGetWidth = files.GetWidth
		result, ok := f(text, path, *align, *color, colorize, funcGetWidth)
		if !ok {
			return
		}
		if ok := v.IsFileHasTxtExtension(*output); !ok {
			fmt.Println("file must has extension <.txt>")
			return
		}
		if *output == "fonts/shadow.txt" || *output == "fonts/standard.txt" || *output == "fonts/thinkertoy.txt" {
			fmt.Println("you cannot save to font files")
			return
		}
		files.Write(*output, result)
	} else {
		result, ok := f(text, path, *align, *color, colorize, funcGetWidth)
		if !ok {
			return
		}
		fmt.Print(result)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
