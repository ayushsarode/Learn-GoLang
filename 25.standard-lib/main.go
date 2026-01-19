package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	stringExample()

	jsonExample()

	fileExamples()

	httpExamples()

	timeExamples()

	sortExamples()

	conversionExamples()
}

func stringExample() {
	fmt.Println("1. Strings Package:")
	text := "   Hello, Go World!  "

	fmt.Printf("Original: '%s'\n", text)

	fmt.Printf("Trimmed: '%s\n", strings.TrimSpace(text))

	fmt.Printf("Uppercase: '%s'\n", strings.ToUpper(text))

	fmt.Printf("Contains 'Go': %t\n", strings.Contains(text, "Go"))

	fmt.Printf("Split by comma: %v\n", strings.Split(text, ","))

	fmt.Printf("Replace 'Go' with 'Golang': %s\n", strings.ReplaceAll(text, "Go", "Golang"))

	fmt.Println()
}

func jsonExample() {
	fmt.Printf("2. JSON Package:")

	person := Person{Name: "ayush", Age: 21, City: "Nagpur"}

	jsonData, err := json.Marshal(person)

	if err != nil {
		log.Printf("Error marshaling: %v", err)
		return
	}
	fmt.Printf("JSON: %s\n", jsonData)

	jsonStr := `{"name":"Bob","age":25,"city":"San Francisco"}`
	var newPerson Person

	err = json.Unmarshal([]byte(jsonStr), &newPerson)

	if err != nil {
		log.Printf("Error unmarshaling: %v", err)
		return
	}

	fmt.Printf("Parsed Person: %+v\n", newPerson)
	fmt.Println()
}

func fileExamples() {
	fmt.Println("3. File Operations:")

	filename := "example.txt"
	content := "Hello from Go standard library! lorem ipsum this is that is my work to do in life of student in raisoni, i have presentation of RM which is research mythology, we have to analysis and work on it, its such a full of crap"

	// write to file

	err := os.WriteFile(filename, []byte(content), 0644)

	if err != nil {
		log.Printf("Error writing file: %v", err)
		return
	}
	fmt.Printf("Wrote to %s\n", filename)

	data, err := os.ReadFile(filename)

	if err != nil {
		log.Printf("Error reading file %s\n", err)
		return
	}
	fmt.Printf("Read from %s: %s\n", filename, data)

	info, err := os.Stat(filename)

	if err != nil {
		log.Printf("Error getting file info %s\n", err)
		return
	}

	infoBytes := info.Size()

	infoKB := float64(infoBytes) / 1024

	infoMB := float64(infoKB) / 1024

	fmt.Printf("File size: %.6f bytes, Modified: %v\n", infoMB, info.ModTime())

	os.Remove(filename)
	fmt.Printf("Cleaned up %s\n", filename)
	fmt.Println()
}

func httpExamples() {
	fmt.Println("4. HTTP Client:")

	// Simple GET request
	resp, err := http.Get("https://api.github.com/users/octocat")
	if err != nil {
		log.Printf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	// Read response body (limit to first 200 characters for demo)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return
	}

	preview := string(body)
	if len(preview) > 200 {
		preview = preview[:200] + "..."
	}
	fmt.Printf("Response preview: %s\n", preview)
	fmt.Println()
}

func timeExamples() {
	fmt.Println("5. Time Package:")

	now := time.Now()

	fmt.Printf("Current time: %s\n", now.Format(time.RFC3339))
	fmt.Printf("Custom format: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Unix timestamp: %d\n", now.Unix())

	// parse time

	timeStr := "2025-08-19 9:45:00"

	parsed, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		log.Printf("Error parsing time: %v", err)
		return
	}
	fmt.Printf("Parsed time: %s\n", parsed.Format(time.RFC3339))

	println()
}

func sortExamples() {
	fmt.Println("6. Sort Package:")

	fruits := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("Before sorting: %v\n", fruits)
	sort.Strings(fruits)
	fmt.Printf("After sorting: %v\n", fruits)


	// sort integers
	nums := []int{24,64,24,14,63,13,54}
	fmt.Printf("Before sorting: %v\n", nums)
	sort.Ints(nums)
	fmt.Printf("After sorting: %v\n", nums)
}

func conversionExamples() {
	fmt.Println("7. String Conversion:")

	str := "45"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error converting to int: %v", err)
		return
	}

	fmt.Printf("String '%s' to int: %d\n", str, num)

	number := 123

	converted := strconv.Itoa(number)
	fmt.Printf("Int %d to string '%s'\n", number, converted)

}
