package visit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactoryImpl_Create(t *testing.T) {
	cardId := int64(473)
	clubId := int64(3)

	factory := FactoryImpl{}
	visit := factory.Create(clubId, cardId)

	assert.NotNil(t, visit)
	assert.Equal(t, clubId, visit.ClubId())
	assert.Equal(t, cardId, visit.CardId())
	assert.Equal(t, "Guest", visit.ClientName())
}
