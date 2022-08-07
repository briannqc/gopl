package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main_411() {
	ops := flag.String("ops", "", "Available: read")
	repo := flag.String("repo", "", "Repository where the issue was/will be created in")
	number := flag.Int("number", 0, "Issue number")
	flag.Parse()

	if *ops == "read" {
		if description, err := GetIssue(*repo, *number); err != nil {
			fmt.Fprintf(os.Stderr, "Read issue: %v on repo: %v failed, err: %v", *number, *repo, err)
		} else {
			printIssueInEditor(description)
		}
	}
}

func printIssueInEditor(s string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}
	tmpFile, err := ioutil.TempFile("", "issue")
	if err != nil {
		log.Fatal(err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	buf := bufio.NewWriter(tmpFile)
	buf.WriteString(s)
	buf.Flush()

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tmpFile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err := cmd.Run(); err != nil {
		log.Fatal()
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
