package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Result holds the outcome of processing a single image.
type Result struct {
	ImagePath   string
	Description string
	Error       error
}

// DashScopeRequest defines the structure for the API request.
type DashScopeRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Message defines the structure for a message in the request.
type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

// Content defines the structure for the content of a message.
type Content struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

// ImageURL defines the structure for the image URL.
type ImageURL struct {
	URL string `json:"url"`
}

// DashScopeResponse defines the structure for the API response.
type DashScopeResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// generateSemantic calls the DashScope API to get a description for an image.
func generateSemantic(imagePath string) (string, error) {
	fmt.Printf("Processing image: %s\n", imagePath)

	apiKey := os.Getenv("DASHSCOPE_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("DASHSCOPE_API_KEY environment variable not set")
	}

	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)
	// Read and Base64 encode the image
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to read image file: %w", err)
	}
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	// Determine mime type from extension
	ext := filepath.Ext(imagePath)
	mimeType := ""
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		mimeType = "image/jpeg"
	case ".png":
		mimeType = "image/png"
	case ".gif":
		mimeType = "image/gif"
	default:
		return "", fmt.Errorf("unsupported image type: %s", ext)
	}
	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Image)

	// Prepare the request payload
	reqPayload := DashScopeRequest{
		Model: "qwen-vl-plus",
		Messages: []Message{
			{
				Role: "user",
				Content: []Content{
					{
						Type: "image_url",
						ImageURL: &ImageURL{
							URL: dataURL,
						},
					},
					{
						Type: "text",
						Text: "Describe this image.",
					},
				},
			},
		},
	}

	payloadBytes, err := json.Marshal(reqPayload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request payload: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create http request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to dashscope: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var apiResp DashScopeResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if apiResp.Error.Message != "" {
		return "", fmt.Errorf("API returned an error: %s", apiResp.Error.Message)
	}

	if len(apiResp.Choices) == 0 {
		return "", fmt.Errorf("no description returned from API")
	}

	description := apiResp.Choices[0].Message.Content
	fmt.Printf("Finished processing image: %s\n", imagePath)
	return description, nil
}

// summarize creates a final summary from all individual descriptions.
func summarize(results []Result) string {
	var summary strings.Builder
	summary.WriteString("Image Processing Workflow Summary:\n")
	summary.WriteString("==================================\n\n")

	successful_count := 0
	for _, result := range results {
		if result.Error != nil {
			summary.WriteString(fmt.Sprintf("Failed to process %s: %v\n", result.ImagePath, result.Error))
		} else {
			summary.WriteString(fmt.Sprintf("- Image: %s\n", result.ImagePath))
			summary.WriteString(fmt.Sprintf("  Description: %s\n", result.Description))
			summary.WriteString("\n")
			successful_count++
		}
	}

	summary.WriteString("----------------------------------\n")
	summary.WriteString(fmt.Sprintf("Summary: Processed %d images, with %d successful and %d failures.\n", len(results), successful_count, len(results)-successful_count))

	return summary.String()
}

func main() {
	// --- Workflow Configuration ---
	// TODO: Add the paths to your images here.
	imagePaths := []string{
		"hello/image-workflow-demo/images/sample1.jpg",
		//		"images/sample2.png",
		//		"images/sample3.gif",
	}

	// --- Workflow Execution ---
	var wg sync.WaitGroup
	resultsChan := make(chan Result, len(imagePaths))

	fmt.Println("Starting image processing workflow...")

	for _, path := range imagePaths {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			description, err := generateSemantic(p)
			resultsChan <- Result{ImagePath: p, Description: description, Error: err}
		}(path)
	}

	// Wait for all goroutines to finish, then close the channel.
	wg.Wait()
	close(resultsChan)

	// --- Result Aggregation ---
	var allResults []Result
	for result := range resultsChan {
		allResults = append(allResults, result)
	}

	// --- Summarization ---
	finalSummary := summarize(allResults)
	fmt.Println("\n" + finalSummary)
}
