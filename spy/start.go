package spy

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zmb3/spotify/v2"
)

func DownloadPlaylist(ctx context.Context, pid string) string {
	user := InitAuth()
	cli := UserData{
		UserClient: user,
	}
	playlistID := spotify.ID(pid)

	trackListJSON, err := cli.UserClient.GetPlaylistItems(ctx, playlistID)

	if err != nil {
		fmt.Println("Playlist not found!", err)
		os.Exit(1)
	}
	for _, val := range trackListJSON.Items {
		cli.TrackList = append(cli.TrackList, *val.Track.Track)
	}

	for page := 0; ; page++ {
		err := cli.UserClient.NextPage(ctx, trackListJSON)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, val := range trackListJSON.Items {
			cli.TrackList = append(cli.TrackList, *val.Track.Track)
		}
	}

	return DownloadTrackList(cli)
}

func DownloadTrackList(cli UserData) string {
	fmt.Println("Found", len(cli.TrackList), "tracks")
	fmt.Println("Searching and downloading tracks")

	for _, val := range cli.TrackList {
		var artistNames []string
		for _, artistInfo := range val.Artists {
			artistNames = append(artistNames, artistInfo.Name)
		}
		searchTerm := strings.Join(artistNames, " ") + " " + val.Name
		youtubeID, err := GetYoutubeId(searchTerm, val.Duration/1000)
		if err != nil {
			log.Printf("Error occurred for %s error: %s", val.Name, err)
			continue
		}
		cli.YoutubeIDList = append(cli.YoutubeIDList, youtubeID)
	}

	var name string

	for index, track := range cli.YoutubeIDList {
		ytURL := "https://www.youtube.com/watch?v=" + track
		name = Downloader(ytURL, cli.TrackList[index])
		fmt.Println()
	}
	fmt.Println("Download complete!" + name)
	// this return only the latest download
	return name
}
