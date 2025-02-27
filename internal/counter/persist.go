package counter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/had-nu/NameSniper/internal/config"
)

// Counter tracks the daily query count
type Counter struct {
	Date  string `json:"date"`  // Date in "YYYY-MM-DD" format
	Count int    `json:"count"` // Number of queries made today
}

// LoadCounter retrieves the counter from file or creates a new one
func LoadCounter() Counter {
	data, err := ioutil.ReadFile(config.CounterFile)
	if err != nil {
		return Counter{Date: time.Now().UTC().Format("2006-01-02"), Count: 0}
	}

	var counter Counter
	if err := json.Unmarshal(data, &counter); err != nil {
		return Counter{Date: time.Now().UTC().Format("2006-01-02"), Count: 0}
	}
	return counter
}

// SaveCounter persists the counter to file
func SaveCounter(counter Counter) {
	data, err := json.Marshal(counter)
	if err != nil {
		fmt.Println("Error saving counter:", err)
		return
	}
	if err := ioutil.WriteFile(config.CounterFile, data, 0644); err != nil {
		fmt.Println("Error writing file:", err)
	}
}