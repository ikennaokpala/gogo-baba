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

// Reads onfiguration YAML or JSON  and returns content of
// the coreesponding file passed to it
func ReadConfigFile(ext string) DbConfig {
	var err error

	db := DbConfig{}
	path := ConfigPath(ext)

	data, _ := ioutil.ReadFile(path)

	if ext == "json" {
		err = json.Unmarshal([]byte(data), &db) // TODO need to figure out better error capture
	} else {
		err = yaml.Unmarshal([]byte(data), &db) // TODO need to figure out better error capture
	}

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// ConfigPath returns database.yml path
func ConfigPath(ext string) string {

	path, _ := os.Getwd()
	dbPath := path + "/config/database." + ext

	return dbPath
}
