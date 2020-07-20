package main

import (
	"flag"
	"fmt"
	"github.com/rwirdemann/legacylab/complexity"
	"os"
	"text/tabwriter"

	"github.com/rwirdemann/legacylab/git"
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
		var found bool
		f.Lines, f.Complextiy, found = complexity.Calc(fmt.Sprintf("%s/%s", path, f.Name))
		if found {
			fmt.Fprintf(w, "%d\t%s\t%d\t%.2f\n", f.Commits, f.Name, f.Lines, f.Complextiy)
		}
	}
	w.Flush()
}
