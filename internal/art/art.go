package art

import (
	"github.com/RickDred/ascii-art/pkg/files"
	"fmt"
	"log"
	"strings"
)

func ReverseAsciiArt(pathToArt, pathToBanner, align, color, colorize string, getWidth func() int) (string, bool) {
	art, err := files.Read(pathToArt)
	if err != nil {
		log.Fatal(err)
	}

	art = strings.ReplaceAll(art, "$", "")
	strs := strings.Split(art, "\n")
	artStr := ""
	result := ""

	runeColor := ""
	reset := ""
	hs := make(map[rune]bool)

	if strings.ToLower(color) != "white" {
		rgb, ok := ConvertToRGB(color)
		if !ok {
			return "", false
		}
		runeColor = fmt.Sprintf("\033[38;2;%v;%v;%vm", rgb.r, rgb.g, rgb.b)
		reset = "\033[0m"
	}

	for _, ch := range colorize {
		hs[ch] = true
	}

	bannerData, err := files.Read(pathToBanner)
	if err != nil {
		log.Fatal(err)
	}

	font := strings.Split(bannerData, "\n")

	count := 0
	for i := 0; i < len(strs); i++ {
		if strs[i] == "" {
			artStr += "\n"
			continue
		}
		if len(strs)-i < 8 {
			break
		}
		isEqual := true
		for ch := ' '; ch <= '~'; ch++ {
			isEqual = true
			line := (int(ch)-int(' '))*9 + 1
			for j := 0; j < 8; j++ {
				if len(strs[i+j]) < count+len(font[line+j]) {
					isEqual = false
					break
				}

				if font[line+j] != strs[i+j][count:count+len(font[line+j])] {
					isEqual = false
					break
				}
			}
			if isEqual {
				if hs[ch] {
					artStr += runeColor + string(ch) + reset
				} else {
					artStr += string(ch)
				}
				count += len(font[line])
				break
			}
		}
		if isEqual {
			i--
		} else {
			count = 0
			artStr += "\n"
			i += 7
		}

	}

	if len(artStr) != 0 {
		artStr = artStr[:len(artStr)-1]
	}
	

	result, ok := justify(artStr, align, getWidth)

	if colorize == "" {
		result = runeColor + result + reset
	}

	return result, ok
}

func justify(text, align string, getWidth func() int) (string, bool) {
	result := ""
	strs := strings.Split(text, "\n")
	width := getWidth()
	switch align {
	case "left":
		result = text
	case "right":
		for i, s := range strs {
			if len(s) > width {
				result += s[:width] + "\n"
				tempText, ok := justify(s[width:], align, getWidth)
				if !ok {
					return "", false
				}
				result += tempText
			} else {
				for len(s) < width {
					s = " " + s
				}
				result += s
				if i < len(strs)-1 {
					result += "\n"
				}
			}
		}
	case "center":
		for i, s := range strs {
			if len(s) > width {
				result += s[:width] + "\n"
				tempText, ok := justify(s[width:], align, getWidth)
				if !ok {
					return "", false
				}
				result += tempText
			} else {
				for len(s)+1 < width {
					s = " " + s + " "
				}
				if len(s) < width {
					s = " " + s
				}
				result += s
				if i < len(strs)-1 {
					result += "\n"
				}
			}
		}
	case "justify":
		for _, s := range strs {

			if len(s) > width {
				temp, ok := justify(s[:width], align, getWidth)
				if !ok {
					return "", false
				}
				result += temp
				s = s[width:]
			}
			words := strings.Fields(s)
			if len(words) <= 1 {
				temp, ok := justify(s, "left", getWidth)
				if !ok {
					return "", false
				}
				result += temp + "\n"
				continue
			}
			freeSpace := width
			for _, w := range words {
				freeSpace -= len(w)
			}
			spaces := ""
			for i := 0; i < freeSpace/(len(words)-1); i++ {
				spaces += " "
			}

			for i, w := range words {
				result += w
				if i != len(words)-1 {
					result += spaces
				}
			}
			result += "\n"
		}
	default:
		return "", false
	}
	return result, true
}

func AsciiArt(text, pathToBanner, align, color, colorize string, getWidth func() int) (string, bool) {
	if text == "" {
		return "", true
	}

	rgb, ok := ConvertToRGB(color)
	if !ok {
		return "", false
	}

	ans := ""
	temp := strings.Split(text, "\\n")
	lines := make([]string, 0)

	for _, w := range temp {
		lines = append(lines, strings.Split(w, "\n")...)
	}

	if len(lines) >= 2 {
		if lines[0] == "" && lines[1] == "" {
			lines = lines[1:]
		}
	}

	bannerData, err := files.Read(pathToBanner)
	if err != nil {
		log.Fatal(err)
	}

	font := strings.Split(bannerData, "\n")
	for _, line := range lines {
		ans += ascii(line, font, align, rgb, colorize, getWidth)
	}

	return ans, true
}

func ascii(word string, font []string, align string, rgb RGB, colorize string, getWidth func() int) string {
	if word == "" {
		return "\n"
	}

	const heightOfCharacter = 9
	y := getWidth()
	str := ""
	nextLine := ""
	hs := make(map[rune]bool)
	color := fmt.Sprintf("\033[38;2;%v;%v;%vm", rgb.r, rgb.g, rgb.b)
	reset := "\033[0m"

	if rgb.r == 255 && rgb.b == 255 && rgb.g == 255 {
		color = ""
		reset = ""
	}

	// characters that should be colorized
	for _, ch := range colorize {
		hs[ch] = true
	}

	// check len, does it fit to terminal
	l := 0
	for j, ch := range word {
		line := (int(ch)-int(' '))*heightOfCharacter + 1

		if l+len(font[line]) > y {
			nextLine = word[j:]
			word = word[:j]
			break
		}

		l += len(font[line])
	}

	if align == "justify" {
		words := strings.Fields(word)

		switch len(words) {
		case 0, 1:
			// print as usual
		default:
			l = 0
			for _, el := range words {
				for _, ch := range el {
					line := (int(ch)-int(' '))*heightOfCharacter + 1
					l += len(font[line])
				}
			}
			freeSpace := y - l
			spacesBetween := freeSpace / (len(words) - 1)
			spaces := ""
			for i := 0; i < spacesBetween; i++ {
				spaces += " "
			}
			for i := 0; i < 8; i++ {
				temp := ""
				for j, el := range words {
					for _, ch := range el {
						line := (int(ch)-int(' '))*heightOfCharacter + i + 1
						charLine := font[line]
						if hs[ch] {
							charLine = color + charLine + reset
						}
						temp += charLine
					}
					if j != len(words)-1 {
						temp += spaces
					}
				}
				str += temp + "\n"
			}

			if nextLine != "" {
				return str + ascii(nextLine, font, align, rgb, colorize, getWidth)
			}
			return str
		}
	}

	for i := 0; i < 8; i++ {
		temp := ""
		lineL := l
		for _, ch := range word {
			line := (int(ch)-int(' '))*heightOfCharacter + i + 1
			charLine := font[line]
			if hs[ch] {
				charLine = color + charLine + reset
			}
			temp += charLine
		}
		switch align {
		case "left":
			// nothing
		case "right":
			for lineL < y {
				temp = " " + temp
				lineL++
			}
		case "center":
			for lineL+1 < y {
				temp = " " + temp + " "
				lineL += 2
			}
		}
		str += temp + "\n"
	}

	if nextLine != "" {
		return str + ascii(nextLine, font, align, rgb, colorize, getWidth)
	}
	return str
}
