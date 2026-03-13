package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("GITHUB_WEBHOOK_SECRET", "super-secret-key")

	mux := http.NewServeMux()

	mux.HandleFunc("/webhook", HandleGitHubWebhook)

	port := ":8080"
	log.Printf("Nexus Swarm Orchestrator starting on port %s...\n", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
