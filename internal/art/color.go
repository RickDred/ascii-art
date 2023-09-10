package art

import (
	"regexp"
	"strconv"
	"strings"
)

type RGB struct {
	r uint8
	g uint8
	b uint8
}

func ConvertToRGB(color string) (RGB, bool) {
	color = strings.ToLower(color)

	colors := map[string]RGB{
		"red":    {255, 0, 0},
		"green":  {0, 255, 0},
		"blue":   {0, 0, 255},
		"yellow": {255, 255, 0},
		"orange": {255, 128, 0},
		"purple": {127, 0, 255},
		"cyan":   {0, 255, 255},
		"pink":   {255, 0, 255},
		"gray":   {128, 128, 128},
		"black":  {0, 0, 0},
		"white":  {255, 255, 255},
	}
	if colors[color] != (RGB{}) || color == "black" {
		return colors[color], true
	}
	if len(color) <= 4 {
		return RGB{}, false
	}

	if color[0] == '#' {
		values, err := strconv.ParseUint(string(color[1:]), 16, 32)
		if err != nil {
			return RGB{}, false
		}

		rgb := RGB{
			uint8(values >> 16),
			uint8((values >> 8) & 0xFF),
			uint8(values & 0xFF),
		}
		return rgb, true
	}

	if color[3] != '(' || color[len(color)-1] != ')' {
		return RGB{}, false
	}

	reg := regexp.MustCompile(`[0-9]+`)
	strNums := reg.FindAllString(color, -1)
	if len(strNums) != 3 {
		return RGB{}, false
	}
	var nums []uint8

	for _, s := range strNums {
		n, err := strconv.Atoi(s)
		if err != nil {
			return RGB{}, false
		}
		nums = append(nums, uint8(n))
	}

	switch color[:3] {
	case "rgb":
		return RGB{nums[0], nums[1], nums[2]}, true
	case "hsl":
		return RGB{}, false
	}

	return RGB{}, false
}
