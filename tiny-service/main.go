package main

import (
	"crypto/subtle"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

// getExpectedToken returns the secret token the service expects.
// It is read from the API_TOKEN environment variable, mirroring the way
// you would configure an ANTHROPIC_API_KEY. A default is provided so the
// service is usable out of the box.
func getExpectedToken() string {
	if t := os.Getenv("API_TOKEN"); t != "" {
		return t
	}
	return "sk-tiny-1234567890"
}

// extractToken reads the token from the request headers.
// It supports both the Anthropic-style "x-api-key" header and the
// "Authorization: Bearer <token>" header.
func extractToken(r *http.Request) string {
	if v := r.Header.Get("x-api-key"); v != "" {
		return v
	}
	if v := r.Header.Get("Authorization"); v != "" {
		return strings.TrimSpace(strings.TrimPrefix(v, "Bearer "))
	}
	return ""
}

// tokensEqual compares two tokens in constant time to avoid timing attacks.
func tokensEqual(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func handler(w http.ResponseWriter, r *http.Request) {
	provided := extractToken(r)
	expected := getExpectedToken()

	if provided == "" {
		log.Printf("request from %s: no token provided", r.RemoteAddr)
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"status":  "fail",
			"message": "❌ Missing token — provide it via the 'x-api-key' header",
		})
		return
	}

	if tokensEqual(provided, expected) {
		log.Printf("request from %s: token OK", r.RemoteAddr)
		writeJSON(w, http.StatusOK, map[string]string{
			"status":  "success",
			"message": "✅ Token is valid — welcome! 🎉",
		})
		return
	}

	log.Printf("request from %s: token INVALID (provided: %q)", r.RemoteAddr, provided)
	writeJSON(w, http.StatusForbidden, map[string]string{
		"status":   "fail",
		"message":  "❌ Invalid token — access denied",
		"provided": provided,
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "6565"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/health", healthHandler)

	// Bind to 0.0.0.0 so the service is reachable on eth0 and can be
	// published to the host with `sbx ports`.
	addr := "0.0.0.0:" + port
	log.Printf("🐳 tiny-service listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
