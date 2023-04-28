package shikimori_test

import (
	"os"
	"testing"

	"github.com/SevereCloud/shikimori"
)

// func needAccessToken(t *testing.T) {
// 	t.Helper()

// 	if authShiki == nil {
// 		t.Skip("ACCESS_TOKEN empty")
// 	}
// }

var shiki, authShiki *shikimori.API //nolint:gochecknoglobals

func TestMain(m *testing.M) {
	shiki = shikimori.NewAPI()

	if token := os.Getenv("ACCESS_TOKEN"); token != "" {
		authShiki = shikimori.NewAPI()
		authShiki.AccessToken = token
	}

	runTests := m.Run()
	os.Exit(runTests)
}
