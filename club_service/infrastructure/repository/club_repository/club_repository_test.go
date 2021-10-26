package club_repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClubInMemoryRepository_getAll(t *testing.T) {
	repository := Instance()

	all := repository.GetAll()
	assert.Equal(t, 1, len(all), all)
}
