package shikimori

type AnimeRateStatus string

const (
	Planned    AnimeRateStatus = "planned"
	Watching   AnimeRateStatus = "watching"
	Rewatching AnimeRateStatus = "rewatching"
	Completed  AnimeRateStatus = "completed"
	OnHold     AnimeRateStatus = "on_hold"
	Dropped    AnimeRateStatus = "dropped"
)
