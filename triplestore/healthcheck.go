package main

import (
	"net/http"
	"os"
	"time"
)

func main() {
	// try to get the list of repositories
	client := http.Client{Timeout: 1 * time.Second}
	res, err := client.Get("http://127.0.0.1:7200/rest/repositories")
	if err != nil {
		os.Exit(1)
	}

	// should be either 200 (when auth is not enabled) or 401 (when auth is enabled)
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusUnauthorized {
		os.Exit(1)
	}

	os.Exit(0)
}
