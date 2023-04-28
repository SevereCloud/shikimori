# shikimori

[![PkgGoDev](https://pkg.go.dev/badge/github.com/SevereCloud/shikimori/v2)](https://pkg.go.dev/github.com/SevereCloud/shikimori)
[![codecov](https://codecov.io/gh/SevereCloud/shikimori/branch/master/graph/badge.svg)](https://codecov.io/gh/SevereCloud/shikimori)
[![release](https://img.shields.io/github/v/tag/SevereCloud/shikimori?label=release)](https://github.com/SevereCloud/shikimori/releases)
[![license](https://img.shields.io/github/license/SevereCloud/shikimori.svg?maxAge=2592000)](https://github.com/SevereCloud/vksdk/blob/master/LICENSE)

```go
package main

import (
	"context"
	"log"

	"github.com/SevereCloud/shikimori"
)

func main() {
	shiki := shikimori.NewAPI()

	// shiki.UserAgent = "MyUserAgent"
	// shiki.AccessToken = os.Getenv("ACCESS_TOKEN")

	resp, err := shiki.UsersAnimeRates(
		context.TODO(),
		"SevereCloud",
		&shikimori.UsersAnimeRateParams{
			Limit:  1,
			Status: shikimori.Watching,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp[0].Anime.Name)
}

```

### Achievements

- [x] GET /api/achievements List user achievements

### Animes

- [ ] GET /api/animes List animes
- [ ] GET /api/animes/:id Show an anime
- [ ] GET /api/animes/:id/roles
- [ ] GET /api/animes/:id/similar
- [ ] GET /api/animes/:id/related
- [ ] GET /api/animes/:id/screenshots
- [ ] GET /api/animes/:id/franchise
- [ ] GET /api/animes/:id/external_links
- [ ] GET /api/animes/:id/topics

### Appear

- [ ] POST /api/appears Mark comments or topics as read

### Bans

- [ ] GET /api/bans List bans

### Calendars

- [x] GET /api/calendar Show a calendar

### Characters

- [ ] GET /api/characters/:id Show a character
- [ ] GET /api/characters/search Search characters

### Clubs

- [ ] GET /api/clubs List clubs
- [ ] GET /api/clubs/:id Show a club
- [ ] PATCH /api/clubs/:id Update a club
- [ ] PUT /api/clubs/:id Update a club
- [ ] GET /api/clubs/:id/animes Show club's animes
- [ ] GET /api/clubs/:id/mangas Show club's mangas
- [ ] GET /api/clubs/:id/ranobe Show club's ranobe
- [ ] GET /api/clubs/:id/characters Show club's characters
- [ ] GET /api/clubs/:id/collections
- [ ] GET /api/clubs/:id/clubs
- [ ] GET /api/clubs/:id/members Show club's members
- [ ] GET /api/clubs/:id/images Show club's images
- [ ] POST /api/clubs/:id/join Join a club
- [ ] POST /api/clubs/:id/leave Leave a club

### Comments

- [ ] GET /api/comments List comments
- [ ] GET /api/comments/:id Show a comment
- [ ] POST /api/comments Create a comment
- [ ] PATCH /api/comments/:id Update a comment
- [ ] PUT /api/comments/:id Update a comment
- [ ] DELETE /api/comments/:id Destroy a comment

### Constants

- [ ] GET /api/constants/anime
- [ ] GET /api/constants/manga
- [ ] GET /api/constants/user_rate
- [ ] GET /api/constants/club
- [ ] GET /api/constants/smileys

### Dialogs

- [ ] GET /api/dialogs List dialogs
- [ ] GET /api/dialogs/:id Show a dialog
- [ ] DELETE /api/dialogs/:id Destroy a dialog

### Favorites

- [ ] POST /api/favorites/:linked_type/:linked_id(/:kind) Create a favorite
- [ ] DELETE /api/favorites/:linked_type/:linked_id Destroy a favorite
- [ ] POST /api/favorites/:id/reorder Assign a new position to a favorite

### Forums

- [ ] GET /api/forums List of forums

### Friends

- [ ] POST /api/friends/:id Create a friend
- [ ] DELETE /api/friends/:id Destroy a friend

### Genres

- [ ] GET /api/genres List genres

### Ignores

### Mangas

- [ ] GET /api/mangas List mangas
- [ ] GET /api/mangas/:id Show a manga
- [ ] GET /api/mangas/:id/roles
- [ ] GET /api/mangas/:id/similar
- [ ] GET /api/mangas/:id/related
- [ ] GET /api/mangas/:id/franchise
- [ ] GET /api/mangas/:id/external_links
- [ ] GET /api/mangas/:id/topics

### Messages

- [ ] GET /api/messages/:id Show a message
- [ ] POST /api/messages Create a message
- [ ] PATCH /api/messages/:id Update a message
- [ ] PUT /api/messages/:id Update a message
- [ ] DELETE /api/messages/:id Destroy a message
- [ ] POST /api/messages/mark_read Mark messages as read or unread
- [ ] POST /api/messages/read_all Mark all messages as read
- [ ] POST /api/messages/delete_all Delete all messages

### People

- [ ] GET /api/people/:id Show a person
- [ ] GET /api/people/search Search people

### Publishers

- [ ] GET /api/publishers List publishers

### Ranobe

- [ ] GET /api/ranobe List ranobe
- [ ] GET /api/ranobe/:id Show a ranobe
- [ ] GET /api/ranobe/:id/roles
- [ ] GET /api/ranobe/:id/similar
- [ ] GET /api/ranobe/:id/related
- [ ] GET /api/ranobe/:id/franchise
- [ ] GET /api/ranobe/:id/external_links
- [ ] GET /api/ranobe/:id/topics

### Reviews

- [ ] POST /api/reviews Create a review
- [ ] PATCH /api/reviews/:id Update a review
- [ ] PUT /api/reviews/:id Update a review
- [ ] DELETE /api/reviews/:id Destroy a review

### Stats

- [ ] GET /api/stats/active_users Users having at least 1 completed animes and active during last month

### Studios

- [x] GET /api/studios List studios

### Styles

- [ ] GET /api/styles/:id Show a style
- [ ] POST /api/styles/preview Preview a style
- [ ] POST /api/styles Create a style
- [ ] PATCH /api/styles/:id Update a style
- [ ] PUT /api/styles/:id Update a style

### Topics

- [ ] GET /api/topics List topics
- [ ] GET /api/topics/updates NewsTopics about database updates
- [ ] GET /api/topics/hot Hot topics
- [ ] GET /api/topics/:id Show a topic
- [ ] POST /api/topics Create a topic
- [ ] PATCH /api/topics/:id Update a topic
- [ ] PUT /api/topics/:id Update a topic
- [ ] DELETE /api/topics/:id Destroy a topic

### User images

- [ ] POST /api/user_images Create an user image

### User rates

- [ ] DELETE /api/user_rates/:type/cleanup Delete entire user rates and history
- [ ] DELETE /api/user_rates/:type/reset Reset all user scores to 0

### Users

- [ ] GET /api/users List users
- [ ] GET /api/users/:id Show an user
- [ ] GET /api/users/:id/info Show user's brief info
- [ ] GET /api/users/whoami Show current user's brief info
- [ ] GET /api/users/sign_out Sign out the user
- [ ] GET /api/users/:id/friends Show user's friends
- [ ] GET /api/users/:id/clubs Show user's clubs
- [ ] GET /api/users/:id/anime_rates Show user's anime list
- [ ] GET /api/users/:id/manga_rates Show user's manga list
- [ ] GET /api/users/:id/favourites Show user's favourites
- [ ] GET /api/users/:id/messages Show current user's messages
- [ ] GET /api/users/:id/unread_messages Show current user's unread messages counts
- [ ] GET /api/users/:id/history Show user history
- [ ] GET /api/users/:id/bans Show user's bans

### Videos

- [ ] GET /api/animes/:anime_id/videos List videos
- [ ] POST /api/animes/:anime_id/videos Create a video
- [ ] DELETE /api/animes/:anime_id/videos/:id Destroy a video
