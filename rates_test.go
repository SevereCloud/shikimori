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
	t.Skip("FIXME: undefined method `split' for 299749:Integer")

	resp, err := shiki.UserRates(context.Background(), &shikimori.UserRatesParams{ //nolint:exhaustruct,exhaustivestruct
		UserID: SevereCloudUser,
	})

	NoError(t, err)
	NotEmpty(t, resp)
}
