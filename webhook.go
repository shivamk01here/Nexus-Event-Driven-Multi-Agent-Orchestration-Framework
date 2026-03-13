package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// HandleGitHubWebhook processes incoming PR events
func HandleGitHubWebhook(w http.ResponseWriter, r *http.Request) {
	// 1. Only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Read the body (we need the raw body for signature validation)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 3. Validate the HMAC signature (Security first!)
	signature := r.Header.Get("X-Hub-Signature-256")
	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")
	if secret != "" && !verifySignature(signature, body, []byte(secret)) {
		log.Println("Unauthorized attempt: Invalid signature")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 4. Ensure it's a Pull Request event
	event := r.Header.Get("X-GitHub-Event")
	if event != "pull_request" {
		w.WriteHeader(http.StatusOK) // Return 200 so GitHub doesn't retry, but we ignore it
		log.Printf("Ignored event type: %s\n", event)
		return
	}

	// 5. Parse the JSON payload
	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 6. Log the successful extraction (Later, we drop this into a Queue)
	log.Printf("Received PR #%d [%s]: %s\n", payload.Number, payload.Action, payload.PullRequest.Title)
	log.Printf("Code diff can be fetched at: %s\n", payload.PullRequest.DiffURL)

	// Acknowledge receipt to GitHub
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received and verified"))
}

// verifySignature checks if the payload was actually sent by GitHub
func verifySignature(signature string, payload, secret []byte) bool {
	if len(signature) != 71 || signature[:7] != "sha256=" {
		return false
	}
	mac := hmac.New(sha256.New, secret)
	mac.Write(payload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(signature[7:]), []byte(expectedMAC))
}
