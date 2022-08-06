package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSearchGithubIssues(t *testing.T) {
	terms := []string{"repo:golang/go", "json", "decoder"}
	result, err := SearchGithubIssues(terms)
	if !assert.NoError(t, err) {
		return
	}

	var lessThanAMonthIssues []Issue
	var lessThanAYearIssues []Issue
	var moreThanAYearIssues []Issue

	lastMonth := time.Now().Add(-30 * 24 * time.Hour)
	lastYear := time.Now().Add(-365 * 24 * time.Hour)
	for _, issue := range result.Items {
		if issue.CreatedAt.After(lastMonth) {
			lessThanAMonthIssues = append(lessThanAMonthIssues, issue)
		} else if issue.CreatedAt.After(lastYear) {
			lessThanAYearIssues = append(lessThanAYearIssues, issue)
		} else {
			moreThanAYearIssues = append(moreThanAYearIssues, issue)
		}
	}

	fmt.Printf("There are %v issues less than a month old\n", len(lessThanAMonthIssues))
	for _, item := range lessThanAMonthIssues {
		fmt.Printf("#%-5d %9.9s %.55s (%v)\n",
			item.Number, item.User.Login, item.Title, item.State)
	}

	fmt.Printf("There are %v issues less than a year old\n", len(lessThanAYearIssues))
	for _, item := range lessThanAYearIssues {
		fmt.Printf("#%-5d %9.9s %.55s (%v)\n",
			item.Number, item.User.Login, item.Title, item.State)
	}

	fmt.Printf("There are %v issues more than a year old.\n", len(lessThanAYearIssues))
	for _, item := range lessThanAYearIssues {
		fmt.Printf("#%-5d %9.9s %.55s (%v)\n",
			item.Number, item.User.Login, item.Title, item.State)
	}
}
