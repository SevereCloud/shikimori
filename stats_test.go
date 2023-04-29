package shikimori_test

import (
	"context"
	"testing"
)

func TestStatsActiveUsers(t *testing.T) {
	t.Parallel()

	resp, err := shiki.StatsActiveUsers(context.Background())

	NoError(t, err)
	NotEmpty(t, resp)
}
