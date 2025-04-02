package main

import (
	"fmt"
	"net/http"
	"os"

	"context"
	"hackspy/spy"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ctx context.Context
)

func downloader(playlistURL string) string {
	ctx = context.Background()

	spotifyURL := playlistURL

	splitURL := strings.Split(spotifyURL, "/")

	spotifyID := splitURL[len(splitURL)-1]

	if strings.Contains(spotifyID, "?") {
		spotifyID = strings.Split(spotifyID, "?")[0]
	}

	return spy.DownloadPlaylist(ctx, spotifyID)
}

type PlaylistRequest struct {
	PlaylistURL string `json:"playlistURL"`
}

func main() {
	router := gin.Default()

	router.POST("/d/spotify", func(c *gin.Context) {
		var request PlaylistRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		filePath := downloader(request.PlaylistURL) + ".mp3"

		fmt.Println(filePath)

		if _, err := os.Stat(filePath); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}

		c.Header("Content-Disposition", "attachment; filename=Careless Whisper.mp3")
		c.Header("Content-Type", "audio/mpeg")

		c.File(filePath)
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
