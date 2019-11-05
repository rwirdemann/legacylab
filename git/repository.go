package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Checkout clones or pulls the given repository.
func Checkout(url string) {
	p := repositoryPath(url)
	if exists(p) {
		println("pulling")
		os.Chdir(p)
		run("git", "pull")
	} else {
		println("cloning")
		run("git", "clone", url, p)
	}
}

func run(name string, arg ...string) {
	cmd := exec.Command("git", arg...)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
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
