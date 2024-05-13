package rainbow

import (
	"bytes"
	"hash/fnv"
	"regexp"

	"github.com/fatih/color"
)

type Colorizer interface {
	Colorize([]byte) (color.Attribute, bool)
}

var _ Colorizer = &ColorString{}
var _ Colorizer = &ColorRegex{}

type ColorString struct {
	Find  string
	Color color.Attribute
}

func (c *ColorString) Colorize(b []byte) (color.Attribute, bool) {
	fnd := []byte(c.Find)
	if !bytes.Contains(b, fnd) {
		return color.Reset, false
	}
	return getColor(fnd, c.Color), true
}

type ColorRegex struct {
	Find  *regexp.Regexp
	Color color.Attribute
}

func (c *ColorRegex) Colorize(b []byte) (color.Attribute, bool) {
	fnd, ok := findMatch(c.Find, b)
	if !ok {
		return color.Reset, false
	}
	return getColor(fnd, c.Color), true
}

func findMatch(re *regexp.Regexp, b []byte) ([]byte, bool) {
	var found []byte
	resp := re.FindSubmatch(b)
	switch len(resp) {
	case 0: // no match
	case 1: // general match
		found = resp[0]
	case 2: // specific submatch
		found = resp[1]
	default: // more... specific matches?
		found = resp[1]
	}
	return found, len(found) > 0
}

//func colorize(seed []byte, col color.Attribute) color.Attribute {
//	if col != HashColor {
//		return col
//	}
//	return hashColor(seed)
//}
//
//func hashColor(seed []byte) color.Attribute {
//	hash := fnv.New32a()
//	if _, err := hash.Write(seed); err != nil {
//		return color.Reset
//	}
//	idx := int(hash.Sum32()) % len(colors)
//	return colors[idx]
//}

func getColor(seed []byte, col color.Attribute) color.Attribute {
	if col != HashColor {
		return col
	}
	hash := fnv.New32a()
	if _, err := hash.Write(seed); err != nil {
		return color.Reset
	}
	idx := int(hash.Sum32()) % len(colors)
	return colors[idx]
}
