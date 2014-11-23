package gocurly

import (
	"strings"
)

var (
	colors    map[string]*color_t
	optimizer *strings.Replacer
)

func init() {
	textAttributes := map[string]*color_t{
		// Text attributes
		"b":         &color_t{"\x1b[1m", "\x1b[22m"},
		"bold":      &color_t{"\x1b[1m", "\x1b[22m"},
		"i":         &color_t{"\x1b[3m", "\x1b[23m"},
		"italic":    &color_t{"\x1b[3m", "\x1b[23m"},
		"u":         &color_t{"\x1b[4m", "\x1b[24m"},
		"underline": &color_t{"\x1b[4m", "\x1b[24m"},
		"blink":     &color_t{"\x1b[5m", "\x1b[25m"},
		"r":         &color_t{"\x1b[7m", "\x1b[27m"},
		"inverse":   &color_t{"\x1b[7m", "\x1b[27m"},
	}
	fgColors := map[string]*color_t{
		"black":   &color_t{"\x1b[30m", "\x1b[39m"},
		"red":     &color_t{"\x1b[31m", "\x1b[39m"},
		"green":   &color_t{"\x1b[32m", "\x1b[39m"},
		"yellow":  &color_t{"\x1b[33m", "\x1b[39m"},
		"blue":    &color_t{"\x1b[34m", "\x1b[39m"},
		"magenta": &color_t{"\x1b[35m", "\x1b[39m"},
		"cyan":    &color_t{"\x1b[36m", "\x1b[39m"},
		"white":   &color_t{"\x1b[37m", "\x1b[39m"},
	}
	bgColors := map[string]*color_t{
		"b:black":   &color_t{"\x1b[40m", "\x1b[49m"},
		"b:red":     &color_t{"\x1b[41m", "\x1b[49m"},
		"b:green":   &color_t{"\x1b[42m", "\x1b[49m"},
		"b:yellow":  &color_t{"\x1b[43m", "\x1b[49m"},
		"b:blue":    &color_t{"\x1b[44m", "\x1b[49m"},
		"b:magenta": &color_t{"\x1b[45m", "\x1b[49m"},
		"b:cyan":    &color_t{"\x1b[46m", "\x1b[49m"},
		"b:white":   &color_t{"\x1b[47m", "\x1b[49m"},
	}
	colors = make(map[string]*color_t)
	for name, color := range textAttributes {
		colors[name] = color
	}
	for name, color := range fgColors {
		colors[name] = color
	}
	for name, color := range bgColors {
		colors[name] = color
	}

	optimizer = strings.NewReplacer(
		"\x1b[39m\x1b[30m", "\x1b[30m",
		"\x1b[39m\x1b[31m", "\x1b[31m",
		"\x1b[39m\x1b[32m", "\x1b[32m",
		"\x1b[39m\x1b[33m", "\x1b[33m",
		"\x1b[39m\x1b[34m", "\x1b[34m",
		"\x1b[39m\x1b[35m", "\x1b[35m",
		"\x1b[39m\x1b[36m", "\x1b[36m",
		"\x1b[39m\x1b[37m", "\x1b[37m",
		"\x1b[49m\x1b[40m", "\x1b[40m",
		"\x1b[49m\x1b[41m", "\x1b[41m",
		"\x1b[49m\x1b[42m", "\x1b[42m",
		"\x1b[49m\x1b[43m", "\x1b[43m",
		"\x1b[49m\x1b[44m", "\x1b[44m",
		"\x1b[49m\x1b[45m", "\x1b[45m",
		"\x1b[49m\x1b[46m", "\x1b[46m",
		"\x1b[49m\x1b[47m", "\x1b[47m",
	)
}

func FormatString(s string) string {
	var history history_t
	var res result_t
	lastIndex := 0
	for {
		startIndex := strings.IndexByte(s[lastIndex:], '<')
		if startIndex == -1 {
			break
		}
		startIndex += lastIndex

		stopIndex := strings.IndexByte(s[startIndex:], '>')
		if stopIndex == -1 {
			break
		}
		stopIndex += startIndex + 1

		res.append(s[lastIndex:startIndex])
		length := stopIndex - startIndex

		if length == 3 && s[startIndex+1:startIndex+2] == "}" {
			if lastColor, ok := history.pop(); ok {
				res.append(lastColor.Close)
				if lastColor2, ok := history.last(); ok {
					res.append(lastColor2.Open)
				}
				lastIndex = stopIndex
				continue
			}
		} else if length >= 3 && s[startIndex+1:startIndex+2] == "{" {
			colorName := s[startIndex+2 : stopIndex-1]
			if color, ok := colors[colorName]; ok {
				res.append(color.Open)
				history.push(color)
				lastIndex = stopIndex
				continue
			}
		}
		// The content between '<' and '>' is invalid, so push the
		// character '<' to the result string and increase lastIndex by 1
		stopIndex = startIndex + 1
		lastIndex = stopIndex
		res.append(s[startIndex:stopIndex])
	}
	// Don't forget to copy the remaining content
	if lastIndex < len(s) {
		res.append(s[lastIndex:])
	}
	// Make sure that all opened colors are closed
	for {
		if color, ok := history.pop(); ok {
			res.append(color.Close)
		} else {
			break
		}
	}
	return res.result()
}

func FormatBytes(s []byte) []byte {
	return []byte(FormatString(string(s)))
}

func FormatStringOptimize(s string) string {
	return OptimizeString(FormatString(s))
}

func FormatBytesOptimize(s []byte) []byte {
	return OptimizeBytes(FormatBytes(s))
}

func OptimizeString(s string) string {
	return optimizer.Replace(s)
}

func OptimizeBytes(s []byte) []byte {
	return []byte(optimizer.Replace(string(s)))
}
