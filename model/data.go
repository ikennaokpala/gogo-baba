package model

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // This is a blank import
	yaml "gopkg.in/yaml.v2"
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

// DbConfig represents information about the database
type DbConfig struct {
	Adapter  string
	Database string
	Host     string
	User     string
	Password string
}

// Drive the connection parameters
// from the DbConfig properties
func (db DbConfig) ConnectString() string {
	// This is the connection String
	// meant for connecting to the
	connectString := db.User + ":" + db.Password + "@/" + db.Database
	return connectString
}

// Connect to Database
func (db DbConfig) Connect() *sql.DB {
	con, _ := sql.Open(db.Adapter, db.ConnectString())
	return con
}

// ReadJSON read and returns content of
// a yaml file passed to it
func ReadJSON() DbConfig {

	db := DbConfig{}
	path := DbConfigPath("json")

	data, _ := ioutil.ReadFile(path)
	err := json.Unmarshal([]byte(data), &db)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

// ReadYAML read and returns content of
// a yaml file passed to it
func ReadYAML() DbConfig {

	db := DbConfig{}
	path := DbConfigPath("yml")

	data, _ := ioutil.ReadFile(path)
	err := yaml.Unmarshal([]byte(data), &db)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// DbConfigPath returns database.yml path
func DbConfigPath(ext string) string {

	path, _ := os.Getwd()
	dbPath := path + "/config/database." + ext

	return dbPath
}
