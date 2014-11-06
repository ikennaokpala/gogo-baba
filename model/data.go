package model

import (
	"log"
	"os"
)

type Site struct {
	Base Blog `json:"blog"`
}

type Blog struct {
	Post   Posts `json:"post"`
	Medium Media `json:"medium"`
}

type Posts map[string]int
type Media map[string]int

type Db struct{}

func ReadYml() string {
	return ""
}

func DbPath() string {

	path, err := os.Getwd()

	dbConfig := path + "/config/database.yml"

	if err != nil {
		log.Fatal(err)
	}

	return dbConfig
}
