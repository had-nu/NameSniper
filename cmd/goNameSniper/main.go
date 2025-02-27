package main

import (
    "fmt"
    "os"

    "github.com/common-nighthawk/go-figure"
    "github.com/had-nu/NameSniper/internal/config"
    "github.com/had-nu/NameSniper/internal/counter"
    "github.com/had-nu/NameSniper/internal/search"
)

func main() {
    // Display stylized title
    banner := figure.NewFigure("NameSniper", "slant", true)
    banner.Print()
    fmt.Println("\nVersion 1.0 - Powered by hadnu\n")

    // Load environment variables
    if err := config.LoadEnv(); err != nil {
        fmt.Println("Error loading .env:", err)
        os.Exit(1)
    }

    // Check command-line arguments
    if len(os.Args) < 2 {
        fmt.Println("Usage: namesniper [first_name] [surname]")
        os.Exit(1)
    }

    firstName := os.Args[1]
    surname := ""
    if len(os.Args) > 2 {
        surname = os.Args[2]
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