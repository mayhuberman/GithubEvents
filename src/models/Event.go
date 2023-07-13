package models

type Event struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Actor   Actor   `json:"actor"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Commits []Commit `json:"commits"`
}

type Commit struct {
	Author Author `json:"author"`
}

type Author struct {
	Email string `json:"email"`
}
