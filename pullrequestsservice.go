package github

import (
	"context"

	"github.com/google/go-github/v33/github"
)

type PullRequestsServiceInterface interface {
	Create(ctx context.Context, owner string, repo string, pull *github.NewPullRequest) (*github.PullRequest, *github.Response, error)
	CreateComment(ctx context.Context, owner string, repo string, number int, comment *github.PullRequestComment) (*github.PullRequestComment, *github.Response, error)
	CreateCommentInReplyTo(ctx context.Context, owner string, repo string, number int, body string, commentID int64) (*github.PullRequestComment, *github.Response, error)
	CreateReview(ctx context.Context, owner, repo string, number int, review *github.PullRequestReviewRequest) (*github.PullRequestReview, *github.Response, error)
	DeleteComment(ctx context.Context, owner string, repo string, commentID int64) (*github.Response, error)
	DeletePendingReview(ctx context.Context, owner, repo string, number int, reviewID int64) (*github.PullRequestReview, *github.Response, error)
	DismissReview(ctx context.Context, owner, repo string, number int, reviewID int64, review *github.PullRequestReviewDismissalRequest) (*github.PullRequestReview, *github.Response, error)
	Edit(ctx context.Context, owner string, repo string, number int, pull *github.PullRequest) (*github.PullRequest, *github.Response, error)
	EditComment(ctx context.Context, owner string, repo string, commentID int64, comment *github.PullRequestComment) (*github.PullRequestComment, *github.Response, error)
	Get(ctx context.Context, owner string, repo string, number int) (*github.PullRequest, *github.Response, error)
	GetComment(ctx context.Context, owner string, repo string, commentID int64) (*github.PullRequestComment, *github.Response, error)
	GetRaw(ctx context.Context, owner string, repo string, number int, opts github.RawOptions) (string, *github.Response, error)
	GetReview(ctx context.Context, owner, repo string, number int, reviewID int64) (*github.PullRequestReview, *github.Response, error)
	IsMerged(ctx context.Context, owner string, repo string, number int) (bool, *github.Response, error)
	List(ctx context.Context, owner string, repo string, opts *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error)
	ListComments(ctx context.Context, owner string, repo string, number int, opts *github.PullRequestListCommentsOptions) ([]*github.PullRequestComment, *github.Response, error)
	ListCommits(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.RepositoryCommit, *github.Response, error)
	ListFiles(ctx context.Context, owner string, repo string, number int, opts *github.ListOptions) ([]*github.CommitFile, *github.Response, error)
	ListPullRequestsWithCommit(ctx context.Context, owner, repo, sha string, opts *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error)
	ListReviewComments(ctx context.Context, owner, repo string, number int, reviewID int64, opts *github.ListOptions) ([]*github.PullRequestComment, *github.Response, error)
	ListReviewers(ctx context.Context, owner, repo string, number int, opts *github.ListOptions) (*github.Reviewers, *github.Response, error)
	ListReviews(ctx context.Context, owner, repo string, number int, opts *github.ListOptions) ([]*github.PullRequestReview, *github.Response, error)
	Merge(ctx context.Context, owner string, repo string, number int, commitMessage string, options *github.PullRequestOptions) (*github.PullRequestMergeResult, *github.Response, error)
	RemoveReviewers(ctx context.Context, owner, repo string, number int, reviewers github.ReviewersRequest) (*github.Response, error)
	RequestReviewers(ctx context.Context, owner, repo string, number int, reviewers github.ReviewersRequest) (*github.PullRequest, *github.Response, error)
	SubmitReview(ctx context.Context, owner, repo string, number int, reviewID int64, review *github.PullRequestReviewRequest) (*github.PullRequestReview, *github.Response, error)
	UpdateBranch(ctx context.Context, owner, repo string, number int, opts *github.PullRequestBranchUpdateOptions) (*github.PullRequestBranchUpdateResponse, *github.Response, error)
	UpdateReview(ctx context.Context, owner, repo string, number int, reviewID int64, body string) (*github.PullRequestReview, *github.Response, error)
}

type PullRequestsService struct {
	PullRequestsServiceInterface
}
