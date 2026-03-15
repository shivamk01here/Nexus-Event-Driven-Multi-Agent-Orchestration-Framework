package main

import (
	"fmt"
	"io"
	"net/http"
)

// DownloadDiff fetches the diff of a PR given the diffURL.
func DownloadDiff(diffURL string, gitHubToken string) (string, error) {
	req, err := http.NewRequest("GET", diffURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// GitHub API usually expects proper headers
	req.Header.Set("Accept", "application/vnd.github.v3.diff")
	if gitHubToken != "" {
		req.Header.Set("Authorization", "token "+gitHubToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch diff: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d fetching diff", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read diff body: %w", err)
	}

	return string(bodyBytes), nil
}
