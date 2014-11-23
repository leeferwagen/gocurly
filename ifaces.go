package gocurly

import (
	"strings"
)

type color_t struct {
	Open  string
	Close string
}

type history_t []*color_t

func (h *history_t) push(c *color_t) {
	*h = append(*h, c)
}

func (h *history_t) pop() (*color_t, bool) {
	l := len(*h)
	if l == 0 {
		return nil, false
	}
	c := (*h)[l-1]
	*h = (*h)[0 : l-1]
	return c, true
}

func (h *history_t) last() (*color_t, bool) {
	l := len(*h)
	if l == 0 {
		return nil, false
	}
	return (*h)[l-1], true
}

type result_t struct {
	List []string
}

func (this *result_t) append(s string) {
	(*this).List = append(*this, s)
}

func (this *result_t) result() string {
	return strings.Join((*this).List, "")
}
