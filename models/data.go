package model

type Site struct {
	Base Blog `json:"blog"`
}

type Blog struct {
	Post   Posts `json:"post"`
	Medium Media `json:"medium"`
}

type Posts map[string]int
type Media map[string]int
