package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/had-nu/NameSniper/internal/config"
)

// SearchResult represents a single Google search result
type SearchResult struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}

// GoogleSearchResponse holds the API response structure
type GoogleSearchResponse struct {
	Items []SearchResult `json:"items"`
}

// BuildQuery constructs a URL-safe query from name and optional surname
func BuildQuery(firstName, surname string) string {
	if surname == "" {
		return url.QueryEscape(firstName)
	}
	return url.QueryEscape(fmt.Sprintf("%s %s", firstName, surname))
}

// SearchGoogle performs a Google API search with the given query
func SearchGoogle(query string) []SearchResult {
	url := fmt.Sprintf("%s?key=%s&cx=%s&q=%s", config.GoogleURL, config.GoogleAPIKey, config.GoogleCX, query)
	fmt.Println("Request URL:", url)

	client := &http.Client{Timeout: config.APITimeout}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Request error:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API error: %s\n", resp.Status)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Raw response:", string(body))
		return nil
	}

	var response GoogleSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Response decoding error:", err)
		return nil
	}

	return response.Items
}