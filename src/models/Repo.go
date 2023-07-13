package models

type Repo struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	StarsCount int    `json:"watchers_count"`
}
