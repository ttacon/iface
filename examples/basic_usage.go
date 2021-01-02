package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/go-github/v33/github"
	giface "github.com/ttacon/iface/github"
)

var errNoOpenPRs = errors.New("no pull requests for repo")

func getMostRecentOpenPR(owner, repo string, prService giface.PullRequestsServiceInterface) (*github.PullRequest, error) {
	repos, _, err := prService.List(context.Background(), owner, repo, &github.PullRequestListOptions{})
	if err != nil {
		return nil, err
	} else if len(repos) == 0 {
		return nil, errNoOpenPRs
	}

	return repos[0], nil
}

func main() {
	client := github.NewClient(nil)

	repo, err := getMostRecentOpenPR("ttacon", "iface", client.PullRequests)
	if err != nil {
		fmt.Println("failed to retrieve PR: ", err)
		os.Exit(1)
	}

	fmt.Println(repo)
}
