package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

var (
	//go:embed ex4.14_template_issues.html
	templateIssues []byte
)

func main() {
	http.HandleFunc("/issues", listIssueHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func listIssueHandler(w http.ResponseWriter, r *http.Request) {
	owner, repo := r.URL.Query().Get("owner"), r.URL.Query().Get("repo")
	if err := RenderHTMLGitHubIssues(w, owner, repo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
}

func RenderHTMLGitHubIssues(w io.Writer, owner, repo string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/issues", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("list issues got: %v", resp.Status)
	}

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return err
	}

	issueList, err := template.New("issueList").Parse(string(templateIssues))
	if err != nil {
		return err
	}

	result := struct {
		Repo        string
		TotalIssues int
		Items       []Issue
	}{
		Repo:        fmt.Sprintf("%v/%v", owner, repo),
		TotalIssues: len(issues),
		Items:       issues,
	}

	return issueList.Execute(w, result)
}

type Issue struct {
	Number  int
	State   string
	User    User
	Title   string
	HTMLURL string `json:"html_url"`
}

type User struct {
	Login     string
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
}
