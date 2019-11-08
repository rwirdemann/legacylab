package usecase

import (
	"github.com/rwirdemann/legacylab.api/questionnaire/domain"
)

// Successor returns the successor question of q regarding its dependencies.
func Successor(q domain.Question) (int, bool) {
	if q.Successor == 0 {
		return -1, false
	}
	return q.Successor, true
}
