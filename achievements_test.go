package shikimori_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/shikimori"
)

func TestAchievements(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Achievements(context.Background(), &shikimori.AchievementsParams{
		UserID: 299749,
	})

	NoError(t, err)
	NotEmpty(t, resp)
}
