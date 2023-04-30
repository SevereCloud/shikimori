package shikimori

import (
	"context"
)

type AnimeKind string

const (
	AnimeKindTv      AnimeKind = "tv"
	AnimeKindMovie   AnimeKind = "movie"
	AnimeKindOva     AnimeKind = "ova"
	AnimeKindOna     AnimeKind = "ona"
	AnimeKindSpecial AnimeKind = "special"
	AnimeKindMusic   AnimeKind = "music"
	// TV13, TV24 and TV48 used in AnimesParams.
	AnimeKindTV13 AnimeKind = "tv_13"
	AnimeKindTV24 AnimeKind = "tv_24"
	AnimeKindTV48 AnimeKind = "tv_48"
)

type AnimeStatus string

const (
	AnimeStatusAnons    AnimeStatus = "anons"
	AnimeStatusOngoing  AnimeStatus = "ongoing"
	AnimeStatusReleased AnimeStatus = "released"
)

type ConstantsAnime struct {
	Kind   []AnimeKind   `json:"kind"`
	Status []AnimeStatus `json:"status"`
}

type ConstantsAnimeParams struct{}

func (s *API) ConstantsAnime(ctx context.Context, params *ConstantsAnimeParams) (resp ConstantsAnime, err error) {
	err = s.get(ctx, &resp, "constants/anime", params)

	return
}

type MangaKind string

const (
	MangaKindManga      MangaKind = "manga"
	MangaKindManhwa     MangaKind = "manhwa"
	MangaKindManhua     MangaKind = "manhua"
	MangaKindLightNovel MangaKind = "light_novel"
	MangaKindNovel      MangaKind = "novel"
	MangaKindOneShot    MangaKind = "one_shot"
	MangaKindDoujin     MangaKind = "doujin"
)

type MangaStatus string

const (
	MangaStatusAnons        MangaStatus = "anons"
	MangaStatusOngoing      MangaStatus = "ongoing"
	MangaStatusReleased     MangaStatus = "released"
	MangaStatusPaused       MangaStatus = "paused"
	MangaStatusDiscontinued MangaStatus = "discontinued"
)

type ConstantsManga struct {
	Kind   []MangaKind   `json:"kind"`
	Status []MangaStatus `json:"status"`
}

type ConstantsMangaParams struct{}

func (s *API) ConstantsManga(ctx context.Context, params *ConstantsMangaParams) (resp ConstantsManga, err error) {
	err = s.get(ctx, &resp, "constants/manga", params)

	return
}

type UserRateStatus string

const (
	UserRatePlanned    UserRateStatus = "planned"
	UserRateWatching   UserRateStatus = "watching"
	UserRateRewatching UserRateStatus = "rewatching"
	UserRateCompleted  UserRateStatus = "completed"
	UserRateOnHold     UserRateStatus = "on_hold"
	UserRateDropped    UserRateStatus = "dropped"
)

type ConstantsUserRate struct {
	Status []UserRateStatus `json:"status"`
}

type ConstantsUserRateParams struct{}

func (s *API) ConstantsUserRate(
	ctx context.Context, params *ConstantsUserRateParams,
) (resp ConstantsUserRate, err error) {
	err = s.get(ctx, &resp, "constants/user_rate", params)

	return
}

type ClubJoinPolicy string

const (
	ClubJoinPolicyFree         ClubJoinPolicy = "free"
	ClubJoinPolicyMemberInvite ClubJoinPolicy = "member_invite"
	ClubJoinPolicyAdminInvite  ClubJoinPolicy = "admin_invite"
	ClubJoinPolicyOwnerInvite  ClubJoinPolicy = "owner_invite"
)

type ClubCommentPolicy string

const (
	ClubCommentPolicyFree    ClubCommentPolicy = "free"
	ClubCommentPolicyMembers ClubCommentPolicy = "members"
	ClubCommentPolicyAdmins  ClubCommentPolicy = "admins"
)

type ClubImageUploadPolicy string

const (
	ClubImageUploadPolicyMembers ClubImageUploadPolicy = "members"
	ClubImageUploadPolicyAdmins  ClubImageUploadPolicy = "admins"
)

type ConstantsClub struct {
	JoinPolicy        []ClubJoinPolicy        `json:"join_policy"`
	CommentPolicy     []ClubCommentPolicy     `json:"comment_policy"`
	ImageUploadPolicy []ClubImageUploadPolicy `json:"image_upload_policy"`
}

type ConstantsClubParams struct{}

func (s *API) ConstantsClub(
	ctx context.Context, params *ConstantsClubParams,
) (resp ConstantsClub, err error) {
	err = s.get(ctx, &resp, "constants/club", params)

	return
}

type ConstantsSmiley struct {
	Bbcode string `json:"bbcode"`
	Path   string `json:"path"`
}

type ConstantsSmileysParams struct{}

func (s *API) ConstantsSmileys(
	ctx context.Context, params *ConstantsSmileysParams,
) (resp []ConstantsSmiley, err error) {
	err = s.get(ctx, &resp, "constants/smileys", params)

	return
}
