package main

type WebhookPayload struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

type PullRequest struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	State   string `json:"state"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	DiffURL string `json:"diff_url"`
}

type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}
