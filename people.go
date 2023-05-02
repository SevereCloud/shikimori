package shikimori

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"
	"time"
)

type RoleStat struct {
	Name  string
	Count int
}

func (r *RoleStat) UnmarshalJSON(data []byte) error {
	var (
		vaule [2]interface{}
		ok    bool //nolint:varnamelen
	)

	err := json.Unmarshal(data, &vaule)
	if err != nil {
		return err //nolint:wrapcheck
	}

	r.Name, ok = vaule[0].(string)
	if !ok {
		return &json.UnmarshalTypeError{
			Value: string(data),
			Type:  reflect.TypeOf(r),
		}
	}

	count, ok := vaule[1].(float64)
	if !ok {
		return &json.UnmarshalTypeError{
			Value: string(data),
			Type:  reflect.TypeOf(r),
		}
	}

	r.Count = int(count)

	return nil
}

type PersonDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type PersonCharacter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Image   Image  `json:"image"`
	URL     string `json:"url"`
}

type PersonRole struct {
	Characters []PersonCharacter `json:"characters"`
	Animes     []Anime           `json:"animes"`
}

type PersonWork struct {
	Anime Anime       `json:"anime"`
	Manga interface{} `json:"manga"`
	Role  string      `json:"role"`
}

type PersonBase struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Image   Image  `json:"image"`
	URL     string `json:"url"`
}

type Person struct {
	PersonBase
	Japanese         string       `json:"japanese"`
	JobTitle         string       `json:"job_title"`
	BirthOn          PersonDate   `json:"birth_on"`
	DeceasedOn       PersonDate   `json:"deceased_on"`
	Website          string       `json:"website"`
	GrouppedRoles    []RoleStat   `json:"groupped_roles"`
	Roles            []PersonRole `json:"roles"`
	Works            []PersonWork `json:"works"`
	TopicID          int          `json:"topic_id"`
	PersonFavoured   bool         `json:"person_favoured"`
	Producer         bool         `json:"producer"`
	ProducerFavoured bool         `json:"producer_favoured"`
	Mangaka          bool         `json:"mangaka"`
	MangakaFavoured  bool         `json:"mangaka_favoured"`
	Seyu             bool         `json:"seyu"`
	SeyuFavoured     bool         `json:"seyu_favoured"`
	UpdatedAt        time.Time    `json:"updated_at"`
	ThreadID         int          `json:"thread_id"`
	Birthday         PersonDate   `json:"birthday"`
}

type PeopleParams struct{}

func (s *API) People(ctx context.Context, id int, params *PeopleParams) (resp Person, err error) {
	err = s.get(ctx, &resp, "people/"+strconv.Itoa(id), params)

	return
}
