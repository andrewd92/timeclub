package visit_service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Instance(t *testing.T) {
	testingService := Instance()

	assert.NotNil(t, testingService)
}
