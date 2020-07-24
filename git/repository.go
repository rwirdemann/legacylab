package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Checkout clones or pulls the given repository.
func Checkout(url string) string {
	p := repositoryPath(url)
	if exists(p) {
		os.Chdir(p)
		run("git", "pull")
	} else {
		run("git", "clone", url, p)
	}
	return p
}

// ChangeFrequency calculates the number of changes applied to each file in the
// given local repository. The resulting list is sorted by number of changes in
// descending order. The result list is limited according to the given limit.
func ChangeFrequency(repository string, limit int) []File {
	cmd := fmt.Sprintf("git log --format=format: --name-only | egrep -v '^$' | sort | uniq -c | sort -r | head -%d", limit)
	out, err := exec.Command("bash", "-c", cmd).Output()
	check(err)
	return toArray(string(out))
}

func toArray(s string) []File {
	a := strings.Split(s, "\n")
	var result []File
	for _, v := range a {
		if isLanguageFile(v) && !isTest(v) {
			result = append(result, NewFile(v))
		}
	}
	return result
}

func isLanguageFile(s string) bool {
	return strings.HasSuffix(s, ".go") || strings.HasSuffix(s, ".java")
}

func isTest(s string) bool {
	return strings.HasSuffix(s, "Tests.java")
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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
