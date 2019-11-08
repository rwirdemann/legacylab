package domain

// Dependency models a dependency between questions.
type Dependency struct {
	ID       int
	Question Question
	Answers  []int
}
