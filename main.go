package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zmb3/spotify"
)

func main() {
	spotifyClientBuilder := NewSpotifyClientBuilder(nil)
	spotifyClient, err := spotifyClientBuilder.GetClient()
	if err != nil {
		log.Fatal(err)
	}
	tr := "medium"
	limit := 10
	tracks, err := spotifyClient.CurrentUsersTopTracksOpt(&spotify.Options{Timerange: &tr, Limit: &limit})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("---\ntitle: \"Spotify Top Tracks - %s\"\ndate: %s\ndraft: false\n---\n\n", time.Now().Format("Jan 2006"), time.Now().Format(time.RFC3339))
	fmt.Print("<p>PUT SOME COMMENTS HERE</p>\n\n")
	fmt.Println("<table>")
	for i := 0; i < len(tracks.Tracks); i++ {
		fmt.Println("\t<tr>")
		track := tracks.Tracks[i]
		artist := track.Artists[0]
		album := track.Album
		img := spotify.Image{} // todo: put in a placeholder image
		if album.Images != nil {
			for j := 0; j < len(album.Images); j++ {
				if album.Images[j].Width == 64 {
					img = album.Images[j]
				}
			}
		}
		fmt.Printf("\t\t<td width='72px'><img src='%s' alt='%s' width='150' height='150'/></td>\n<td><strong>%s</strong> - %s<br/>%s (%s)</td>\n",
			img.URL, album.Name, track.Name, artist.Name, album.Name, album.ReleaseDate[:4])
		fmt.Println("\t<tr>")
	}
	fmt.Println("</table>")
}
