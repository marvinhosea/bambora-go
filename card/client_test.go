package card

import (
	"github.com/marvinhosea/bambora-go"
	"testing"

	"github.com/marvinhosea/bambora-go/util"
	"github.com/stretchr/testify/assert"
)


func TestCardNew(t *testing.T) {
	merchantId := ""
	profilePasscode := ""
	paymentPasscode := ""

	_ = bambora.GeneratePasscodes(
		merchantId,
		profilePasscode,
		paymentPasscode)

	token, err := Tokenize(&bambora.CardParams{
		Number: util.String("4030000010001234"),
		ExpiryMonth: util.String("02"),
		ExpiryYear: util.String("20"),
		CVD: util.String("123"),
	})

	assert.Nil(t, err)
	assert.NotNil(t, token)
}