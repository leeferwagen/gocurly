package gocurly

import (
	"strings"
)

var colors map[string]*Color

const (
	minlen = 3
	maxlen = 12
)

func init() {
	colors = map[string]*Color{
		"b":         &Color{"\x1b[1m", "\x1b[22m"},
		"bold":      &Color{"\x1b[1m", "\x1b[22m"},
		"i":         &Color{"\x1b[3m", "\x1b[23m"},
		"italic":    &Color{"\x1b[3m", "\x1b[23m"},
		"u":         &Color{"\x1b[4m", "\x1b[24m"},
		"underline": &Color{"\x1b[4m", "\x1b[24m"},
		"blink":     &Color{"\x1b[5m", "\x1b[25m"},
		"r":         &Color{"\x1b[7m", "\x1b[27m"},
		"inverse":   &Color{"\x1b[7m", "\x1b[27m"},
		"black":     &Color{"\x1b[30m", "\x1b[39m"},
		"red":       &Color{"\x1b[31m", "\x1b[39m"},
		"green":     &Color{"\x1b[32m", "\x1b[39m"},
		"yellow":    &Color{"\x1b[33m", "\x1b[39m"},
		"blue":      &Color{"\x1b[34m", "\x1b[39m"},
		"magenta":   &Color{"\x1b[35m", "\x1b[39m"},
		"cyan":      &Color{"\x1b[36m", "\x1b[39m"},
		"white":     &Color{"\x1b[37m", "\x1b[39m"},
		"grey":      &Color{"\x1b[90m", "\x1b[39m"},
	}
}

type Color struct {
	Open  string
	Close string
}

type History []*Color

func (h *History) push(c *Color) {
	*h = append(*h, c)
}

func (h *History) pop() (*Color, bool) {
	l := len(*h)
	if l == 0 {
		return nil, false
	}
	c := (*h)[l-1]
	*h = (*h)[0 : l-1]
	return c, true
}

func (h *History) last() (*Color, bool) {
	l := len(*h)
	if l == 0 {
		return nil, false
	}
	return (*h)[l-1], true
}

func FormatString(s string) string {
	var history History
	result := ""
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

		result += s[lastIndex:startIndex]
		length := stopIndex - startIndex

		if length == 3 && s[startIndex+1:startIndex+2] == "}" {
			if lastColor, ok := history.pop(); ok {
				result += lastColor.Close
				if lastColor2, ok := history.last(); ok {
					result += lastColor2.Open
				}
				lastIndex = stopIndex
				continue
			}
		} else if length >= minlen && length <= maxlen && s[startIndex+1:startIndex+2] == "{" {
			colorName := s[startIndex+2 : stopIndex-1]
			if color, ok := colors[colorName]; ok {
				result += color.Open
				history.push(color)
				lastIndex = stopIndex
				continue
			}
		}
		// The content between '<' and '>' is invalid, so push the
		// character '<' to the result string and increase lastIndex by 1
		stopIndex = startIndex + 1
		lastIndex = stopIndex
		result += s[startIndex:stopIndex]
	}
	// Don't forget to copy the remaining content
	if lastIndex < len(s) {
		result += s[lastIndex:]
	}
	// Make sure that all opened colors are closed
	for {
		if color, ok := history.pop(); ok {
			result += color.Close
		} else {
			break
		}
	}
	return result
}

func FormatBytes(s []byte) []byte {
	return []byte(FormatString(string(s)))
}
