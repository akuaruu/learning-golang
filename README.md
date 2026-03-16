# GitHub User Activity CLI

A simple command-line interface (CLI) tool to fetch and display the recent public activity of a GitHub user. Built with Go (Golang) using only standard libraries.

## Features
- Fetches recent events from the GitHub API.
- Displays activities cleanly in the terminal (e.g., PushEvent, IssuesEvent, WatchEvent).
- Handles errors (e.g., user not found, API failures).
- Zero external dependencies.

## Prerequisites
- [Go](https://go.dev/dl/) installed on your machine.

## How to Run

1. Clone this repository:
   ```bash/terminal
   git clone [https://github.com/akuaruu/learning-golang.git](https://github.com/akuaruu/learning-golang.git)
   cd learning-golang

2. Build the executable:
   ```bash/terminal
   go build -o github-activity main.go

3. Run the CLI tool by passing a GitHub username as an argument:
   ```bash/terminal
   ./github-activity <github_username>

4. Example :
   ```bash/terminal
   ./github-activity <github_username>

Submitted for the https://roadmap.sh/projects/github-user-activity
