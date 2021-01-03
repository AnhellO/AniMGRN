package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/AniMGRN/contenedor-c/pkg/mongo"
)

const (
	url          = "localhost:27017"
	username     = ""
	password     = ""
	databaseName = "anime_db"
)

func main() {
	session, err := mongo.NewConnection(url, username, password)
	if err != nil {
		fmt.Println(err)
	}
	col := session.DB(databaseName).C("anime")
	datab := session.DB(databaseName)
	count, err2 := col.Count()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("Database Name:", datab.Name)
	fmt.Println("Collection FullName:", col.FullName)
	fmt.Println(fmt.Sprintf("Documents count: %d", count))

	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

}
