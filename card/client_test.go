package card

import (
	bambora "github.com/marvinhosea/bambora-go"
	"testing"

	"github.com/marvinhosea/bambora-go/util"
	"github.com/stretchr/testify/assert"
)


func TestCardNew(t *testing.T) {

	bambora.MerchantId = "383610147"
	bambora.AccountPasscode = "78BCE183B92E49EAA3C4F97CEDDE8539"
	bambora.Passcode = bambora.GeneratePasscode()

	card, err := Tokenize(&bambora.CardParams{
		Number: util.String("4030000010001234"),
		ExpiryMonth: util.String("02"),
		ExpiryYear: util.String("20"),
		CVD: util.String("123"),
	})

	assert.Nil(t, err)
	assert.NotNil(t, card)
}