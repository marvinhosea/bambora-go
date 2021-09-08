package card

import (
	"testing"

	"github.com/marvinhosea/bambora-go/util"
	go_bambora "github.com/marvinhosea/bambora-go"
	"github.com/stretchr/testify/assert"
)


func TestCardNew(t *testing.T) {
	card, err := New(&go_bambora.CardParams{
		Number: util.String("4030000010001234"),
		ExpiryMonth: util.String("02"),
		ExpiryYear: util.String("20"),
		CVD: util.String("123"),
	})

	assert.Nil(t, err)
	assert.NotNil(t, card)
}