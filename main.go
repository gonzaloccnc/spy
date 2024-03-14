package main

import (
	"context"
	"hackspy/spy"
	"strings"
)

var (
	ctx context.Context
)

func main() {
	ctx = context.Background()

	spotifyURL := "https://open.spotify.com/playlist/5gCiaLib6kvYuYHiizg2Qa?si=6d90b58f8f3d408b"

	splitURL := strings.Split(spotifyURL, "/")

	spotifyID := splitURL[len(splitURL)-1]

	if strings.Contains(spotifyID, "?") {
		spotifyID = strings.Split(spotifyID, "?")[0]
	}

	spy.DownloadPlaylist(ctx, spotifyID)

}
