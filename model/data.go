package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Site is an exported type that
// contains a Base of type Blog.
type Site struct {
	Base Blog `json:"blog"`
}

// Blog is an exported type that
// contains a Post for Posts and Medium for Media.
type Blog struct {
	Post   Posts `json:"post"`
	Medium Media `json:"medium"`
}

// Posts holds collection of post data
type Posts map[string]int

//Media holds a collection of photos, videos and document data
type Media map[string]int

// Db represents information about the database
type Db struct {
	Adapter  string
	Database string
	Host     string
	User     string
	Password string
}

// ReadJSON read and returns content of
// a yaml file passed to it
func ReadJSON() Db {

	db := Db{}
	path := DbPath()

	data, _ := ioutil.ReadFile(path)
	err := json.Unmarshal([]byte(data), &db)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// DbPath returns database.yml path
func DbPath() string {

	path, _ := os.Getwd()
	dbPath := path + "/config/database.json"

	return dbPath
}

func connect() {

}
