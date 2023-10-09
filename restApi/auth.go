package restApi

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

func generateAPIKey(passphrase string) string {
	hash := sha256.Sum256([]byte(passphrase))
	return hex.EncodeToString(hash[:])
}

func validateHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-KEY")
		apiPassphrase := r.Header.Get("API-PASSPHRASE")

		if apiKey == "" || apiPassphrase == "" {
			http.Error(w, "Missing API-KEY or API-PASSPHRASE header", http.StatusUnauthorized)
			return
		}

		// Validate that the API-KEY matches the hashed API-PASSPHRASE
		expectedAPIKey := generateAPIKey(apiPassphrase)
		if apiKey != expectedAPIKey {
			http.Error(w, "Invalid API-KEY", http.StatusUnauthorized)
			return
		}

		// If the headers are valid, call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, you've passed the validation!")
}

// func main() {
// 	http.Handle("/", validateHeadersMiddleware(http.HandlerFunc(mainHandler)))
// 	http.ListenAndServe(":8080", nil)
// }
