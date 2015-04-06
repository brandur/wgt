package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yosssi/ace"
)

type Artist struct {
	Familiarity  float32
	Genres       []string
	Location     *ArtistLocation
	Name         string
	PopularSongs []string `json:"popular_songs"`
	Sample       *string
	URLs         ArtistURLs
}

type ArtistLocation struct {
	City    *string
	Country *string
}

type ArtistURLs struct {
	HomeURL      *string `json:"home_url"`
	WikipediaURL *string `json:"wikipedia_url"`
}

var (
	artists []Artist
)

func handler(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("index", "", &ace.Options{
		DynamicReload: true,
		FuncMap:       templateFuncMap(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"artists": artists,
	}
	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func readArtists() []Artist {
	dat, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	var artists []Artist
	err = json.Unmarshal(dat, &artists)
	if err != nil {
		panic(err)
	}

	return artists
}

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"JoinStrings": func(strs []string) string {
			return strings.Join(strs, ", ")
		},
		"SafeLink": func(uri *string, text string) template.HTML {
			if uri != nil {
				return template.HTML(fmt.Sprintf(`<a href="%s">%s</a>`, *uri, text))
			} else {
				return ""
			}
		},
		"SafeLocation": func(location *ArtistLocation) string {
			if location != nil && location.City != nil {
				return *location.City + ", " + *location.Country
			} else if location != nil {
				return *location.Country
			} else {
				return ""
			}
		},
		"TruncateFamiliarity": func(familiarity float32) string {
			return fmt.Sprintf("%.2f", familiarity)
		},
	}
}

func main() {
	artists = readArtists()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
