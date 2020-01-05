package main

import (
	"flag"

	"github.com/rwirdemann/legacylab/git"
)

func main() {
	url := flag.String("url", "https://github.com/spring-projects/spring-data-jpa.git", "repository url")
	flag.Parse()
	path := git.Checkout(*url)
	git.ChangeFrequency(path, 30)
}
