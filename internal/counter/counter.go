package counter

import (
	"fmt"
	"time"

	"github.com/had-nu/NameSniper/internal/config"
)

// CanQuery checks if queries are available within the daily limit
func CanQuery() bool {
	counter := LoadCounter()
	today := time.Now().UTC().Format("2006-01-02")

	if counter.Date != today {
		counter.Date = today
		counter.Count = 0
		SaveCounter(counter)
	}

	if counter.Count >= config.DailyLimit {
		return false
	}

	if counter.Count == config.DailyLimit-1 {
		fmt.Println("Warning: This is the last free query of the day!")
	}

	return true
}

// UpdateCounter increments the query count after a successful search
func UpdateCounter() {
	counter := LoadCounter()
	counter.Count++
	SaveCounter(counter)
	fmt.Printf("Queries used today: %d/%d\n", counter.Count, config.DailyLimit)
}