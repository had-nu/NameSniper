package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/had-nu/NameSniper/internal/config"
    "github.com/had-nu/NameSniper/internal/counter"
    "github.com/had-nu/NameSniper/internal/search"
    "github.com/had-nu/NameSniper/internal/ui"
)

func main() {
    // Display stylized title using our custom banner package
    ui.PrintBanner()

    // Load environment variables
    if err := config.LoadEnv(); err != nil {
        fmt.Println("Error loading .env:", err)
        os.Exit(1)
    }

    // Prompt for user input
    fmt.Print(">>> Enter name to search <FirstName Surname>: ")
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        os.Exit(1)
    }

    // Parse input into firstName and surname
    input = strings.TrimSpace(input)
    if input == "" {
        fmt.Println("Error: No name provided. Please enter a name to search.")
        os.Exit(1)
    }

    parts := strings.Split(input, " ")
    firstName := parts[0]
    surname := ""
    if len(parts) > 1 {
        surname = parts[1]
    }

    query := search.BuildQuery(firstName, surname)
    fmt.Printf("Searching for: %s %s\n", firstName, surname)

    // Check daily query limit
    if !counter.CanQuery() {
        fmt.Println("Daily limit of 100 free queries reached. Try again tomorrow.")
        os.Exit(1)
    }

    // Perform Google search
    results := search.SearchGoogle(query)
    if len(results) > 0 {
        fmt.Println("Results found:")
        for i, result := range results {
            fmt.Printf("%d. %s\n   URL: %s\n   Snippet: %s\n", i+1, result.Title, result.Link, result.Snippet)
        }
    } else {
        fmt.Println("No results found.")
    }

    // Update query counter
    counter.UpdateCounter()
}