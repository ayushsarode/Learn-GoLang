package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type Result struct {
    URL   string
    Data  map[string]any
    Error error
    Took  time.Duration
}

func fetchWithRetry(ctx context.Context, url string, maxRetries int) Result {
    start := time.Now()
    
    for attempt := 1; attempt <= maxRetries; attempt++ {
        req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
        client := &http.Client{Timeout: 5 * time.Second}
        
        resp, err := client.Do(req)
        if err != nil {
            if attempt == maxRetries {
                return Result{
                    URL:   url,
                    Error: err,
                    Took:  time.Since(start),
                }
            }
            time.Sleep(time.Duration(attempt) * time.Second)
            continue
        }
        defer resp.Body.Close()
        
        var data map[string]interface{}
        if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
            return Result{URL: url, Error: err, Took: time.Since(start)}
        }
        
        return Result{
            URL:   url,
            Data:  data,
            Took:  time.Since(start),
        }
    }
    
    return Result{URL: url, Error: fmt.Errorf("unknown error"), Took: time.Since(start)}
}

func main() {
    ctx := context.Background()
    
    result := fetchWithRetry(ctx, "https://api.github.com/users/golang", 3)
    
    if result.Error != nil {
        fmt.Printf("❌ %s failed: %v (took %v)\n", result.URL, result.Error, result.Took)
    } else {
        fmt.Printf("✓ %s succeeded (took %v)\n", result.URL, result.Took)
        fmt.Printf("  Name: %v\n", result.Data["name"])
    }
}
