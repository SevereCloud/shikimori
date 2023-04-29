package shikimori

import (
	"context"
)

func (s *API) StatsActiveUsers(ctx context.Context) (resp []int, err error) {
	err = s.get(ctx, &resp, "stats/active_users", nil)

	return
}
