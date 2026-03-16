package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GitHubEvent struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`

	//adding payload
	Payload struct {
		Action string `json:"action"`
		Size   int    `json:"size"`
	} `json:"payload"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Command : go run main.go <username_github>")
		os.Exit(1)
	}

	username := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error Fetching Data: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	//Status cek,
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("User '%s' tidak ditemukan.\n", username)
		os.Exit(1)

	} else if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed retrieving data. Status HTTP : %d\n", resp.StatusCode)
		os.Exit(1)
	}

	//reading the body response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error while reading response: %v\n", err)
		os.Exit(1)
	}

	//Parsing JSON innto slice of GItHubEvent
	var events []GitHubEvent
	err = json.Unmarshal(body, &events)
	if err != nil {
		fmt.Printf("Error when Parsing JSON: %v\n", err)
		os.Exit(1)
	}

	//showing the data to terminal
	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	fmt.Printf("Recent activity for %s:\n", username)
	fmt.Println("-------------------------------------------")

	for _, event := range events {
		switch event.Type {
		case "PushEvent":

			fmt.Printf("- Pushed %d commit(s) to repository %s\n", event.Payload.Size, event.Repo.Name)

		case "IssuesEvent":
			fmt.Printf("- %s issue in %s\n", event.Payload.Action, event.Repo.Name)

		case "WatchEvent":
			fmt.Printf("- starred on %s\n", event.Repo.Name)

		case "CreateEvent":
			fmt.Printf("- Creating repo or branch %s\n", event.Repo.Name)

		default:
			fmt.Printf("- %s in %s\n", event.Type, event.Repo.Name)
		}

	}
}
