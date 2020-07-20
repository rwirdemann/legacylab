package complexity

import (
	"github.com/rwirdemann/legacylab/identation"
)

func Calc(path string) (int, float32, bool) {
	lines, found := identation.ReadFile(path)
	if found {
		lines = identation.ReplaceTabs(lines)
		lines = identation.RemoveComments(lines)
		lines = identation.RemoveEmptyLines(lines)
		var indent int
		for _, l := range lines {
			indent = indent + identation.LeadingSpaces(l)
		}
		return len(lines), float32(indent) / float32(len(lines)), true
	}
	return 0, 0, false
}
