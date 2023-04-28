package shikimori_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalendar(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Calendar(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}
