package shikimori_test

import (
	"context"
	"testing"
)

func TestCalendar(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Calendar(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}
