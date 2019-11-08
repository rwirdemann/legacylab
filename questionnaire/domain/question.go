package domain

// Question models a question.
type Question struct {
	ID           int
	Text         string
	Answers      []Answer
	Successor    int
	Dependencies []Dependency
}
