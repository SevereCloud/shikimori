package internal_test

import (
	"context"
	"testing"
	"time"

	"github.com/SevereCloud/shikimori/internal"
)

func TestLimit_RateLimit(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	limit := internal.NewLimit(1, time.Second)

	first := time.Now() //nolint:ifshort

	_ = limit.RateLimit(ctx)
	_ = limit.RateLimit(ctx)

	if time.Since(first) < time.Second {
		t.Error("Since time < 1s")
	}
}
