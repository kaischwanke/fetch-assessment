//go:build integration
// +build integration

package itest

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {

	tests := []struct {
		name      string
		inputFile string
		want      int
	}{
		{
			name:      "Exercise example one",
			inputFile: "example1.json",
			want:      28,
		},
		{
			name:      "Exercise example two",
			inputFile: "example2.json",
			want:      109,
		},
		{
			name:      "Morning Receipt",
			inputFile: "morning-receipt.json",
			want:      15,
		},
		{
			name:      "Simple Receipt",
			inputFile: "simple-receipt.json",
			want:      31,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			validate(t, "testdata/"+testCase.inputFile, testCase.want)
		})
	}

}

func validate(t *testing.T, fileName string, expectedPoint int) {

	data, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}
	in := string(data)

	uuid := callPost(t, err, in)

	points := callGet(t, err, uuid)
	assert.Equal(t, expectedPoint, points)
}

func callPost(t *testing.T, err error, in string) string {
	req, err := http.NewRequest("POST", "http://localhost:8080/receipts/process", bytes.NewBuffer([]byte(in)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v, ensure the server is running", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Define the struct to hold the response (if not already defined)
	type uuidResponse struct {
		UUID string `json:"id"`
	}

	var result uuidResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	return result.UUID
}

func callGet(t *testing.T, err error, uuid string) int {

	url := "http://localhost:8080/receipts/" + uuid + "/points"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create GET request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v, ensure the server is running", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Define struct to parse the response (if needed)
	type pointsResponse struct {
		Points int `json:"points"`
	}

	var result pointsResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	return result.Points
}
