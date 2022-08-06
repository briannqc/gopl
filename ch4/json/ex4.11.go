package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	ops := flag.String("ops", "", "Available: create, read, update, and close")
	repo := flag.String("repo", "", "Repository where the issue was/will be created in")
	number := flag.Int("number", 0, "Issue number")
	flag.Parse()

	if *ops == "read" {
		if description, err := GetIssue(*repo, *number); err != nil {
			fmt.Fprintf(os.Stderr, "Read issue: %v on repo: %v failed, err: %v", *number, *repo, err)
		} else {
			fmt.Println(description)
		}
	}
}

func GetIssue(repo string, number int) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%v/issues/%v", repo, number)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Got non OK response: %v", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return "", err
	}

	sb := &strings.Builder{}
	fmt.Fprintf(sb, "#%v - %v (%s)\n", issue.Number, issue.Title, issue.State)
	fmt.Fprintln(sb, issue.Body)

	return sb.String(), nil
}
