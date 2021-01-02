github.com/ttacon/iface/github
===

When developing code using the `github.com/google/go-github` module, you have to make do with
testing calls against mocked web servers. Personally, I find it easier, and more delightful to
write code using interfaces and so I in order to test my code, I can provide the correctly mocked
versions, instead of needing to write any extra code. This library makes that possible, without
making the experience too onerous. Let's look at an example of how we'd mock a call to list all
pull requests for a given GitHub repo.

## Using this interface

Let's say that we wanted to return the [most recent open pull request](./examples/basic_usage.go) for this repo:

```go
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
```

Then to test it, since we used the `iface` type, [we can do](./examples/basic_usage_test.go):

```go
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
```

Which is much simpler than testing it would have been otherwise.