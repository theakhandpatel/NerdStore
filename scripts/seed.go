package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

const (
	apiURL = "http://localhost:8080/v1/resources"
)

var (
	titles = []string{"golang", "python", "javascript", "css", "html", "react", "vue", "angular", "node.js", "express"}
	links  = []string{"www.google.com", "www.github.com", "www.stackoverflow.com", "www.medium.com", "www.dev.to"}
	tags   = []string{"frontend", "backend", "fullstack", "devops", "mobile", "web", "database", "api", "cloud", "security"}
)

type Resource struct {
	Title string   `json:"title"`
	Link  string   `json:"link"`
	Tags  []string `json:"tags,omitempty"`
}

func generateResource() Resource {
	resource := Resource{
		Title: titles[rand.Intn(len(titles))],
		Link:  links[rand.Intn(len(links))],
	}

	// Randomly decide whether to include tags
	if rand.Float32() < 0.9 {
		numTags := rand.Intn(3) + 1
		selectedIndices := rand.Perm(len(tags))[:numTags] // Get unique indices for tags
		resource.Tags = make([]string, numTags)
		for i, idx := range selectedIndices {
			resource.Tags[i] = tags[idx]
		}
	}

	return resource
}

func createResource(resource Resource) error {
	jsonData, err := json.Marshal(resource)
	if err != nil {
		return err
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	return nil
}

func main() {

	for i := 0; i < 100; i++ {
		resource := generateResource()
		err := createResource(resource)
		if err != nil {
			fmt.Printf("Error creating resource %d: %v  %+v\n", i+1, err, resource)
		} else {
			fmt.Printf("Created resource %d: %+v\n", i+1, resource)
		}
	}

	fmt.Println("Database seeded successfully with 100 dummy resources via API.")
}
