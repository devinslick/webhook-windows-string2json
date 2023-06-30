package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	// Parse command line arguments
	urlFlag := flag.String("url", "", "Webhook destination URL")
	string2jsonFlag := flag.String("string2json", "", "Key-value pairs to convert to JSON")
	flag.Parse()

	// Validate URL flag
	if *urlFlag == "" {
		log.Fatal("Missing URL flag. Please provide the destination URL.")
	}

	// Validate string2json flag
	if *string2jsonFlag == "" {
		log.Fatal("Missing string2json flag. Please provide key-value pairs in the format 'key1 value1 key2 value2'.")
	}

	// Split the string2json flag value into key-value pairs
	pairs := splitString2JSONFlag(*string2jsonFlag)

	// Ensure an even number of key-value pairs
	if len(pairs)%2 != 0 {
		log.Fatal("Invalid number of key-value pairs.")
	}

	// Create a map to hold the key-value pairs
	jsonData := make(map[string]string)

	// Convert key-value pairs to JSON
	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]
		if strings.Contains(value, "%RANDOMNUMBER%") {
			value = strings.ReplaceAll(value, "%RANDOMNUMBER%", generateRandomNumber())
		}
		jsonData[key] = value
	}

	// Convert JSON to byte array
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("Failed to convert JSON: %v", err)
	}

	// Create HTTP POST request
	req, err := http.NewRequest("POST", *urlFlag, strings.NewReader(string(jsonBytes)))
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK response: %v", resp.Status)
	}

	// Print success message
	fmt.Println("Webhook sent successfully.")
}

func splitString2JSONFlag(s string) []string {
	// Manually split the string2json flag value into key-value pairs
	// by splitting on spaces and removing empty strings
	var pairs []string
	for _, p := range strings.Split(s, " ") {
		if p != "" {
			pairs = append(pairs, p)
		}
	}
	return pairs
}

func generateRandomNumber() string {
	// Generate a random number between 1 and 1,000,000
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000000) + 1
	return fmt.Sprintf("%d", randomNumber)
}
