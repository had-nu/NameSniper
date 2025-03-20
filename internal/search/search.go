package search

import (
	"encoding/json"
	"fmt"
	"io"
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

// BuildQuery constructs a map of categorised, URL-safe queries
func BuildQuery(firstName, surname string) map[string]string {
    baseQuery := firstName
    if surname != "" {
        baseQuery += " " + surname
    }
    exactQuery := fmt.Sprintf(`"%s"`, baseQuery)

    // Define the categorised dorking queries
    queries := map[string]string{
        "Name Search": url.QueryEscape(exactQuery),
        "Social Media":     url.QueryEscape(fmt.Sprintf(`%s site:linkedin.com OR site:facebook.com OR site:instagram.com OR site:twitter.com`, exactQuery)),
        "Documents":        url.QueryEscape(fmt.Sprintf(`%s filetype:pdf OR filetype:doc OR filetype:xlsx`, exactQuery)),
        "Professional Info": url.QueryEscape(fmt.Sprintf(`%s intitle:curriculo OR intitle:CV`, exactQuery)),
        "Personal Info":    url.QueryEscape(fmt.Sprintf(`%s intext:address OR intext:contact OR intext:email OR intext:username`, exactQuery)),
    }

    return queries
}

// SearchGoogle performs a Google API search for each query and groups results by category
func SearchGoogle(queries map[string]string) map[string][]SearchResult {
    client := &http.Client{Timeout: config.APITimeout}
    categorisedResults := make(map[string][]SearchResult)

    for category, query := range queries {
        // Map to deduplicate results within this category
        resultMap := make(map[string]SearchResult)

        url := fmt.Sprintf("%s?key=%s&cx=%s&q=%s", config.GoogleURL, config.GoogleAPIKey, config.GoogleCX, query)
        fmt.Println("Request URL for", category, ":", url)

        resp, err := client.Get(url)
        if err != nil {
            fmt.Println("Request error for", category, ":", err)
            continue
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            fmt.Printf("API error for %s: %s\n", category, resp.Status)
            body, _ := io.ReadAll(resp.Body)
            fmt.Println("Raw response:", string(body))
            continue
        }

        var response GoogleSearchResponse
        if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
            fmt.Println("Response decoding error for", category, ":", err)
            continue
        }

        // Add results to the map (deduplicate by URL)
        for _, item := range response.Items {
            resultMap[item.Link] = SearchResult{
                Title:   item.Title,
                Link:    item.Link,
                Snippet: item.Snippet,
            }
        }

        // Convert map to slice
        var results []SearchResult
        for _, result := range resultMap {
            results = append(results, result)
        }

        categorisedResults[category] = results
    }

    return categorisedResults
}