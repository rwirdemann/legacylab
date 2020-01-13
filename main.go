package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/rwirdemann/legacylab/git"
	"github.com/rwirdemann/legacylab/identation"
)

func main() {
	url := flag.String("url", "https://github.com/spring-projects/spring-data-jpa.git", "repository url")
	limit := flag.Int("limit", 30, "file limit")
	flag.Parse()
	path := git.Checkout(*url)
	files := git.ChangeFrequency(path, *limit)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", "Commits", "File", "Lines", "Complexity")
	for _, f := range files {
		path := fmt.Sprintf("%s/%s", path, f.Name)
		lines, found := identation.ReadFile(path)
		if found {
			lines = identation.ReplaceTabs(lines)
			lines = identation.RemoveComments(lines)
			lines = identation.RemoveEmptyLines(lines)
			var indent int
			for _, l := range lines {
				indent = indent + identation.LeadingSpaces(l)
			}
			f.Lines = len(lines)
			f.Complextiy = float32(indent) / float32(f.Lines)
			fmt.Fprintf(w, "%d\t%s\t%d\t%.2f\n", f.Commits, f.Name, f.Lines, f.Complextiy)
		}
	}
	w.Flush()
}
