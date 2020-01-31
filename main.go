package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
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

	funcMap := template.FuncMap{
		"now":              time.Now,
		"dateFmtMMYYYY":    func(t time.Time) string { return t.Format("Jan 2006") },
		"dateFmtRFC3339":   func(t time.Time) string { return t.Format(time.RFC3339) },
		"albumReleaseYear": func(album *spotify.SimpleAlbum) string { return album.ReleaseDate[:4] },
	}

	f, err := os.Open("index.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.New("toptracks").Funcs(funcMap).Parse(string(b)))
	t.Execute(os.Stdout, tracks.Tracks)
}
