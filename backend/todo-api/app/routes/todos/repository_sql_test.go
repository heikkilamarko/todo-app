package todos

import (
	"testing"
)

func TestSQLRepositoryImplementsRepository(t *testing.T) {
	var _ Repository = &SQLRepository{}
}
