package git

import (
	"strconv"
	"strings"
)

type File struct {
	Name       string
	Commits    int
	Lines      int
	Complextiy float32
}

func NewFile(s string) File {
	a := strings.Split(strings.Trim(s, " "), " ")
	commits, err := strconv.Atoi(a[0])
	if err != nil {
		panic(err)
	}
	return File{
		Name:    a[1],
		Commits: commits,
	}
}
