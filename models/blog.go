package model

type Site struct {
	Base Blog
}

type Post struct {
}

type Blog struct {
	Post   Posts
	Medium Media
}

type Posts map[string]int
type Media map[string]int
