package club_controller

import (
	"errors"
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/club_repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var getAll func() ([]*club.Club, error)

type mockRepository struct {
}

func (m mockRepository) Save(_ *club.Club) (*club.Club, error) {
	panic("implement me")
}

func (m mockRepository) GetAll() ([]*club.Club, error) {
	return getAll()
}

func (m mockRepository) GetById(_ int64) (*club.Club, error) {
	panic("implement me")
}

func TestMain(m *testing.M) {
	club_repository.MockTestRepository(mockRepository{})

	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
	getAll = func() ([]*club.Club, error) {
		return []*club.Club{
			club.DefaultClub(),
		}, nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	All(c)

	expected := `[{"id":1,"name":"Club A","open_time":"12:00","currency":"USD","prices":[{"price_period":{"from":0,"to":60},"value_per_minute":10}]}]`
	response, _ := io.ReadAll(w.Body)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, string(response))
}

func TestAll_WhenDbError(t *testing.T) {
	getAll = func() ([]*club.Club, error) {
		return nil, errors.New("test error")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	All(c)

	expected := "DB error"
	response, _ := io.ReadAll(w.Body)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, string(response))
}
