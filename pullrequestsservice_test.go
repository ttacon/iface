package github

import (
	"context"
	"testing"

	"github.com/google/go-github/v33/github"
)

type ListPullRequests PullRequestsService

func (lpr ListPullRequests) List(ctx context.Context, owner string, repo string, opts *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error) {
	return nil, nil, ErrExpectedForTesting
}

func Test_PullRequestsService_List(t *testing.T) {
	lpr := ListPullRequests{}
	if _, _, err := lpr.List(nil, "", "", nil); err != ErrExpectedForTesting {
		t.Error("received unexpected error: ", err)
	}
}
