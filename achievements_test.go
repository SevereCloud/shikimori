package shikimori_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/shikimori"
	"github.com/stretchr/testify/assert"
)

func TestAchievements(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Achievements(context.Background(), &shikimori.AchievementsParams{
		UserID: 299749,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}
