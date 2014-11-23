package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// Db represents information about the database
type Db struct {
	Adapter, Database, Host, User, Password string
}

// Drive the connection parameters
// from the Db properties
func (db Db) ConnectString() string {
	// This is the connection String
	// meant for connecting to the database
	return db.User + ":" + db.Password + "@/" + db.Database
}

// Connect to Database
func (db Db) Connect() *sql.DB {
	con, err := sql.Open(db.Adapter, db.ConnectString())

	if err != nil {
		panic(err.Error())
	}

	return con
}

func (db Db) Read(sql string) {
	conn := db.Connect()
	rows, err := conn.Query(sql)
	values, err := conn.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	cols, _ := rows.Columns()

	fmt.Println("Columns: \n", cols)
	fmt.Println("Exec : \n", values)
	fmt.Println("Query: \n", rows)
}

// Reads configuration YAML or JSON  and returns content of
// the coresponding file passed to it
func SetupDB(ext string) Db {
	var err error

	db := Db{}
	path := ConfigPath(ext)

	data, _ := ioutil.ReadFile(path)

	if ext == "json" {
		err = json.Unmarshal([]byte(data), &db) // TODO need to figure out better error capture
	} else {
		err = yaml.Unmarshal([]byte(data), &db) // TODO need to figure out better error capture
	}

	if err != nil {
		panic(err.Error())
	}

	return db
}

// ConfigPath returns database.yml path
func ConfigPath(ext string) string {

	path, _ := os.Getwd()
	dbPath := path + "/config/database." + ext

	return dbPath
}

// Run function that sets thing up
func Run() {
	db := SetupDB("yml")
	db.Read("SELECT * FROM GFB_OPTIONS")
}
