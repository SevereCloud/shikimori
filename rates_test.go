package shikimori_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/shikimori"
)

func TestUserRate(t *testing.T) {
	t.Parallel()

	resp, err := shiki.UserRate(context.Background(), 57370056, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestUserRates(t *testing.T) {
	t.Parallel()
	// To fix unproccessable entity error, convert everything to string
	// For example {"user_id": 299749} -> {"user_id": "299749"}
	t.Skip("FIXME: undefined method `split' for 299749:Integer")

	resp, err := shiki.UserRates(context.Background(), &shikimori.UserRatesParams{ //nolint:exhaustruct,exhaustivestruct
		UserID: SevereCloudUser,
	})

	NoError(t, err)
	NotEmpty(t, resp)
}
