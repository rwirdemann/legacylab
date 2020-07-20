package git

import (
	"strconv"
	"strings"
)

type File struct {
	Name       string  `json:"name"`
	Commits    int     `json:"commits"`
	Lines      int     `json:"lines"`
	Complextiy float32 `json:"complextiy"`
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
