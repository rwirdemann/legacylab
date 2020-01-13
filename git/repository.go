package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
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
		if isLanguageFile(v) {
			result = append(result, NewFile(v))
		}
	}
	return result
}

func Lines(respository, file string) int {
	f := fmt.Sprintf("%s/%s", respository, file)
	out, err := exec.Command("wc", "-l", f).Output()
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(out), " ")
	a := strings.Split(s, " ")
	i, err := strconv.Atoi(a[0])
	if err != nil {
		panic(err)
	}
	return i
}

func isLanguageFile(s string) bool {
	return strings.HasSuffix(s, ".go") || strings.HasSuffix(s, ".java")
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
