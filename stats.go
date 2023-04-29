package shikimori

import (
	"context"
)

type StatsActiveUsersParams struct{}

func (s *API) StatsActiveUsers(ctx context.Context, params *StatsActiveUsersParams) (resp []int, err error) {
	err = s.get(ctx, &resp, "stats/active_users", params)

	return
}
