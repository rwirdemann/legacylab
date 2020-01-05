package git

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Checkout clones or pulls the given repository.
func Checkout(url string) {
	p := repositoryPath(url)
	if exists(p) {
		os.Chdir(p)
		fmt.Printf("pulling repository [dir=%s]\n", p)
		run("git", "pull")
	} else {
		println("cloning")
		run("git", "clone", url, p)
	}
}

// ChangeFrequency calculates the number of changes applied to each file in the
// given local repository. The resulting list is sorted by number of changes in
// descending order. The result list is limited according to the given limit.
func ChangeFrequency(repository string, limit int) {
	os.Chdir(repository)
	c1 := exec.Command("git", "log", "--format=format:", "--name-only")
	c2 := exec.Command("egrep", "-v", "^[[:space:]]*$")
	c3 := exec.Command("sort")
	c4 := exec.Command("uniq", "-c")
	c5 := exec.Command("sort", "-r")
	c6 := exec.Command("head", fmt.Sprintf("-%d", limit))

	var b bytes.Buffer
	commands := []*exec.Cmd{c1, c2, c3, c4, c5, c6}
	var pipes []*io.PipeWriter
	for i, c := range commands {
		if i == len(commands)-1 {
			c.Stdout = &b
		} else {
			r, w := io.Pipe()
			pipes = append(pipes, w)
			c.Stdout = w
			next := commands[i+1]
			next.Stdin = r
		}
	}
	for _, c := range commands {
		c.Start()
	}

	for i, c := range commands {
		c.Wait()
		if i < len(commands)-1 {
			pipes[i].Close()
		}
	}

	s := b.String()
	fmt.Printf(s)
}

func run(name string, arg ...string) []byte {
	out, err := exec.Command("git", arg...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func repositoryPath(url string) string {
	parts := strings.Split(url, "/")
	name := strings.TrimSuffix(parts[len(parts)-1], ".git")
	return fmt.Sprintf("%s/tmp/%s", os.Getenv("HOME"), name)
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		return !os.IsNotExist(err)
	}
	return true
}
