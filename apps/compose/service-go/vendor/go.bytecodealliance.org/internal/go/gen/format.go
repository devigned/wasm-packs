package gen

import "strings"

const (
	DocCommentPrefix = "//"
	LineLength       = 80
)

// FormatDocComments formats documentation comment text (without // or /*)
// into multiple lines of max length LineLength, prefixed by //, suitable
// for inclusion as documentation comments in Go source code.
func FormatDocComments(docs string, indent bool) string {
	if docs == "" {
		return ""
	}
	space := ' '
	if indent {
		space = '\t'
	}
	var b strings.Builder
	var lineLength = 0
	for _, c := range docs {
		if lineLength == 0 {
			b.WriteString(DocCommentPrefix)
			lineLength = len(DocCommentPrefix)
		}
		switch c {
		case '\n':
			b.WriteRune('\n')
			lineLength = 0
			continue
		case ' ':
			switch {
			case lineLength == len(DocCommentPrefix):
				// Ignore leading spaces
				continue
			case lineLength > LineLength:
				b.WriteRune('\n')
				lineLength = 0
				continue
			}
		default:
			if lineLength == len(DocCommentPrefix) {
				b.WriteRune(space)
				lineLength++
			}
		}
		b.WriteRune(c)
		lineLength++
	}
	if lineLength != 0 {
		b.WriteRune('\n')
	}
	return b.String()
}
