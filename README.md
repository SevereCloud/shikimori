# shikimori

[![PkgGoDev](https://pkg.go.dev/badge/github.com/SevereCloud/shikimori/v2)](https://pkg.go.dev/github.com/SevereCloud/shikimori)
[![codecov](https://codecov.io/gh/SevereCloud/shikimori/branch/master/graph/badge.svg)](https://codecov.io/gh/SevereCloud/shikimori)
[![release](https://img.shields.io/github/v/tag/SevereCloud/shikimori?label=release)](https://github.com/SevereCloud/shikimori/releases)
[![license](https://img.shields.io/github/license/SevereCloud/shikimori.svg?maxAge=2592000)](https://github.com/SevereCloud/vksdk/blob/master/LICENSE)

Status - https://github.com/SevereCloud/shikimori/milestone/1

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

	ctx := context.Background()

	resp, err := shiki.UsersAnimeRates(ctx, "SevereCloud", &shikimori.UsersAnimeRateParams{
		Limit:  1,
		Status: shikimori.Watching,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp[0].Anime.Name)
}

```
