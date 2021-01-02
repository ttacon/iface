package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v33/github"
	giface "github.com/ttacon/iface/github"
)

type emptyListPRHandler giface.PullRequestsService

func (lprh emptyListPRHandler) List(ctx context.Context, owner string, repo string, opts *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error) {
	return nil, nil, nil
}

func Test_getMostRecentOpenPR(t *testing.T) {
	_, err := getMostRecentOpenPR("foo", "bar", emptyListPRHandler{})
	if err != errNoOpenPRs {
		t.Error("expected errNoOpenPRs, got: ", err)
	}
}
