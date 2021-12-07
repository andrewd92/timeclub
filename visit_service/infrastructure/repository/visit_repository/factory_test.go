package visit_repository

import (
	"github.com/andrewd92/timeclub/visit_service/infrastructure/dao/visit_dao"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Instance(t *testing.T) {
	instance := Instance()
	assert.NotNil(t, instance)
}

func Test_InstanceFromCache(t *testing.T) {
	repository = &VisitDbRepository{dao: visit_dao.Instance()}

	instance := Instance()

	assert.Equal(t, repository, instance)
}
