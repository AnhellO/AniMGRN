package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Song struct.
type Song struct {
	ID                  int
	ArtistString        string
	CreateDate          string
	DefaultName         string
	DefaultNameLanguage string
	FavoritedTimes      int
	LengthSeconds       int
	Name                string
	PVServices          string
	PublishDate         string
	RatingScore         int
	SongType            string
	Status              string
	Version             int
}

func main() {
	response, err := http.Get("https://vocadb.net/api/songs")

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
