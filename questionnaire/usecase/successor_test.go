package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rwirdemann/legacylab.api/questionnaire/domain"
)

var questionnaire domain.Questionnaire
var q1, q2, q3 domain.Question

func init() {
	q1 = domain.Question{ID: 1}
	q2 = domain.Question{ID: 2}
	q3 = domain.Question{ID: 3}

	q1.Successor = q2.ID
	q2.Successor = q3.ID
	questionnaire.Questions = []domain.Question{q1, q2, q3}
}

func TestSuccessorWithoudDependcies(t *testing.T) {
	id, b := Successor(q1)
	assert.True(t, b)
	assert.Equal(t, q2.ID, id)

	id, b = Successor(q3)
	assert.False(t, b)
	assert.Equal(t, -1, id)
}
