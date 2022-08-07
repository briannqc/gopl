package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"
)

var (
	//go:embed ex4.14_template_issues.html
	templateIssues []byte

	//go:embed ex4.14_template_contributors.html
	templateContributors []byte

	//go:embed ex4.14_template_milestones.html
	templateMilestones []byte
)

func main() {
	http.HandleFunc("/issues", listIssueHandler)
	http.HandleFunc("/contributors", listContributorHandler)
	http.HandleFunc("/milestones", listMilestoneHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func listIssueHandler(w http.ResponseWriter, r *http.Request) {
	owner, repo := r.URL.Query().Get("owner"), r.URL.Query().Get("repo")
	if err := RenderHTMLGitHubIssues(w, owner, repo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
}

func listContributorHandler(w http.ResponseWriter, r *http.Request) {
	owner, repo := r.URL.Query().Get("owner"), r.URL.Query().Get("repo")
	if err := RenderHTMLGitHubContributors(w, owner, repo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
}

func listMilestoneHandler(w http.ResponseWriter, r *http.Request) {
	owner, repo := r.URL.Query().Get("owner"), r.URL.Query().Get("repo")
	if err := RenderHTMLGitHubMilestone(w, owner, repo); err != nil {
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

func RenderHTMLGitHubContributors(w io.Writer, owner, repo string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/contributors", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("list contributors got: %v", resp.Status)
	}

	var contributors []Contributor
	if err := json.NewDecoder(resp.Body).Decode(&contributors); err != nil {
		return err
	}

	contributorList, err := template.New("contributorList").Parse(string(templateContributors))
	if err != nil {
		return err
	}

	result := struct {
		Repo  string
		Total int
		Items []Contributor
	}{
		Repo:  fmt.Sprintf("%v/%v", owner, repo),
		Total: len(contributors),
		Items: contributors,
	}

	return contributorList.Execute(w, result)
}

func RenderHTMLGitHubMilestone(w io.Writer, owner, repo string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/milestones", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("list milestone got: %v", resp.Status)
	}

	var milestones []Milestone
	if err := json.NewDecoder(resp.Body).Decode(&milestones); err != nil {
		return err
	}

	milestoneList, err := template.New("milestoneList").Parse(string(templateMilestones))
	if err != nil {
		return err
	}

	result := struct {
		Repo  string
		Total int
		Items []Milestone
	}{
		Repo:  fmt.Sprintf("%v/%v", owner, repo),
		Total: len(milestones),
		Items: milestones,
	}

	return milestoneList.Execute(w, result)
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

type Contributor struct {
	User
	Contributions int
}

type Milestone struct {
	HTMLURL      string `json:"html_url"`
	Number       int
	Title        string
	CreatedAt    time.Time `json:"created_at"`
	Description  string
	Creator      User
	State        string
	OpenIssues   int `json:"open_issues"`
	ClosedIssues int `json:"closed_issues"`
}
