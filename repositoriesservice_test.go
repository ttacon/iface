package github

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-github/v33/github"
)

var ErrExpectedForTesting = errors.New("expected error for testing")

type ListAll RepositoriesService

func (la ListAll) ListAll(ctx context.Context, opts *github.RepositoryListAllOptions) ([]*github.Repository, *github.Response, error) {
	return nil, nil, ErrExpectedForTesting
}

func Test_RepositoriesService_ListAll(t *testing.T) {
	listService := ListAll{}
	if _, _, err := listService.ListAll(context.Background(), nil); err != ErrExpectedForTesting {
		t.Error("received unexpected error: ", err)
	}
}
