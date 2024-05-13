package rainbow

import (
	"io"
	"regexp"

	"github.com/fatih/color"
)

const HashColor color.Attribute = -1

var (
	DefaultColors = []*ColorRegex{
		{Find: FileNameRegex, Color: HashColor},
		{Find: LogNameRegex, Color: HashColor},
	}

	// FileNameRegex works with hclog.LoggerOptions{IncludeLocation: true}
	FileNameRegex = regexp.MustCompile(`(\S+\.go):\d+:\s`)
	// LogNameRegex works with hclog default config.
	LogNameRegex = regexp.MustCompile(`(\S+):\s`)

	// HashColor will pick amongst these colors
	colors = []color.Attribute{
		color.FgRed,
		color.FgGreen,
		color.FgYellow,
		color.FgBlue,
		color.FgMagenta,
		color.FgCyan,
		color.FgWhite,
	}
)

func NewColorWriter(w io.Writer) *ColorWriter {
	c := &ColorWriter{
		Writer: w,
		//Colors: DefaultColors, // Cannot use 'DefaultColors' (type []*ColorRegex) as the type []Colorizer
	}
	// TODO: this seems real dumb.
	for _, col := range DefaultColors {
		c.Colors = append(c.Colors, col)
	}
	return c
}

var _ io.Writer = &ColorWriter{}

type ColorWriter struct {
	io.Writer

	// Colors pick a color to apply to each Write. The first successful match
	// will be applied, so put more-specific matchers earlier in the slice.
	Colors []Colorizer

	// Ignore does no colorizing for any matching lines.
	// Ignoring important lines might actually help them stand out,
	// since they would be white, and never faint, in a sea of mixed colors.
	Ignore []*regexp.Regexp
}

func (c *ColorWriter) WithColor(m ...Colorizer) *ColorWriter {
	// prepend to give these precedence over the default
	c.Colors = append(m, c.Colors...)
	return c
}

func (c *ColorWriter) WithIgnore(i ...*regexp.Regexp) *ColorWriter {
	c.Ignore = append(c.Ignore, i...)
	return c
}

func (c *ColorWriter) Write(b []byte) (n int, err error) {
	for _, i := range c.Ignore {
		if i.Match(b) {
			return c.Writer.Write(b)
		}
	}

	col := color.New()

	level := getLevel(b)
	switch level {
	case "DEBUG", "TRACE":
		col = col.Add(color.Faint)
	case "ERROR", "FATAL":
		col = col.Add(color.Bold)
	}

	for _, m := range c.Colors {
		if matched, ok := m.Colorize(b); ok {
			col = col.Add(matched)
			break
		}
	}

	return col.Fprint(c.Writer, string(b))
}

func getLevel(b []byte) string {
	re := regexp.MustCompile(`\[(TRACE|DEBUG|INFO|WARN|ERROR|FATAL)]`)
	level, ok := findMatch(re, b)
	if !ok {
		return ""
	}
	return string(level)
}
