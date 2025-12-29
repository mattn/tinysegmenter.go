package tinysegmenter

import (
	"github.com/mattn/tinysegmenter.go/internal"
)

func Segment(s string) []string {
	return internal.Segment(s)
}
