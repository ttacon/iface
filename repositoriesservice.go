package github

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-github/v33/github"
)

type RepositoriesServiceInterface interface {
	AddAdminEnforcement(ctx context.Context, owner, repo, branch string) (*github.AdminEnforcement, *github.Response, error)
	AddAppRestrictions(ctx context.Context, owner, repo, branch string, slug []string) ([]*github.App, *github.Response, error)
	AddCollaborator(ctx context.Context, owner, repo, user string, opts *github.RepositoryAddCollaboratorOptions) (*github.CollaboratorInvitation, *github.Response, error)
	CompareCommits(ctx context.Context, owner, repo string, base, head string) (*github.CommitsComparison, *github.Response, error)
	CompareCommitsRaw(ctx context.Context, owner, repo, base, head string, opts github.RawOptions) (string, *github.Response, error)
	Create(ctx context.Context, org string, repo *github.Repository) (*github.Repository, *github.Response, error)
	CreateComment(ctx context.Context, owner, repo, sha string, comment *github.RepositoryComment) (*github.RepositoryComment, *github.Response, error)
	CreateDeployment(ctx context.Context, owner, repo string, request *github.DeploymentRequest) (*github.Deployment, *github.Response, error)
	CreateDeploymentStatus(ctx context.Context, owner, repo string, deployment int64, request *github.DeploymentStatusRequest) (*github.DeploymentStatus, *github.Response, error)
	CreateFile(ctx context.Context, owner, repo, path string, opts *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *github.Response, error)
	CreateFork(ctx context.Context, owner, repo string, opts *github.RepositoryCreateForkOptions) (*github.Repository, *github.Response, error)
	CreateFromTemplate(ctx context.Context, templateOwner, templateRepo string, templateRepoReq *github.TemplateRepoRequest) (*github.Repository, *github.Response, error)
	CreateHook(ctx context.Context, owner, repo string, hook *github.Hook) (*github.Hook, *github.Response, error)
	CreateKey(ctx context.Context, owner string, repo string, key *github.Key) (*github.Key, *github.Response, error)
	CreateProject(ctx context.Context, owner, repo string, opts *github.ProjectOptions) (*github.Project, *github.Response, error)
	CreateRelease(ctx context.Context, owner, repo string, release *github.RepositoryRelease) (*github.RepositoryRelease, *github.Response, error)
	CreateStatus(ctx context.Context, owner, repo, ref string, status *github.RepoStatus) (*github.RepoStatus, *github.Response, error)
	Delete(ctx context.Context, owner, repo string) (*github.Response, error)
	DeleteComment(ctx context.Context, owner, repo string, id int64) (*github.Response, error)
	DeleteDeployment(ctx context.Context, owner, repo string, deploymentID int64) (*github.Response, error)
	DeleteFile(ctx context.Context, owner, repo, path string, opts *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *github.Response, error)
	DeleteHook(ctx context.Context, owner, repo string, id int64) (*github.Response, error)
	DeleteInvitation(ctx context.Context, owner, repo string, invitationID int64) (*github.Response, error)
	DeleteKey(ctx context.Context, owner string, repo string, id int64) (*github.Response, error)
	DeletePreReceiveHook(ctx context.Context, owner, repo string, id int64) (*github.Response, error)
	DeleteRelease(ctx context.Context, owner, repo string, id int64) (*github.Response, error)
	DeleteReleaseAsset(ctx context.Context, owner, repo string, id int64) (*github.Response, error)
	DisableAutomatedSecurityFixes(ctx context.Context, owner, repository string) (*github.Response, error)
	DisableDismissalRestrictions(ctx context.Context, owner, repo, branch string) (*github.PullRequestReviewsEnforcement, *github.Response, error)
	DisablePages(ctx context.Context, owner, repo string) (*github.Response, error)
	DisableVulnerabilityAlerts(ctx context.Context, owner, repository string) (*github.Response, error)
	Dispatch(ctx context.Context, owner, repo string, opts github.DispatchRequestOptions) (*github.Repository, *github.Response, error)
	DownloadContents(ctx context.Context, owner, repo, filepath string, opts *github.RepositoryContentGetOptions) (io.ReadCloser, *github.Response, error)
	DownloadReleaseAsset(ctx context.Context, owner, repo string, id int64, followRedirectsClient *http.Client) (rc io.ReadCloser, redirectURL string, err error)
	Edit(ctx context.Context, owner, repo string, repository *github.Repository) (*github.Repository, *github.Response, error)
	EditHook(ctx context.Context, owner, repo string, id int64, hook *github.Hook) (*github.Hook, *github.Response, error)
	EditRelease(ctx context.Context, owner, repo string, id int64, release *github.RepositoryRelease) (*github.RepositoryRelease, *github.Response, error)
	EditReleaseAsset(ctx context.Context, owner, repo string, id int64, release *github.ReleaseAsset) (*github.ReleaseAsset, *github.Response, error)
	EnableAutomatedSecurityFixes(ctx context.Context, owner, repository string) (*github.Response, error)
	EnablePages(ctx context.Context, owner, repo string, pages *github.Pages) (*github.Pages, *github.Response, error)
	EnableVulnerabilityAlerts(ctx context.Context, owner, repository string) (*github.Response, error)
	Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error)
	GetAdminEnforcement(ctx context.Context, owner, repo, branch string) (*github.AdminEnforcement, *github.Response, error)
	GetArchiveLink(ctx context.Context, owner, repo string, archiveformat github.ArchiveFormat, opts *github.RepositoryContentGetOptions, followRedirects bool) (*url.URL, *github.Response, error)
	GetBranch(ctx context.Context, owner, repo, branch string) (*github.Branch, *github.Response, error)
	GetBranchProtection(ctx context.Context, owner, repo, branch string) (*github.Protection, *github.Response, error)
	GetByID(ctx context.Context, id int64) (*github.Repository, *github.Response, error)
	GetCodeOfConduct(ctx context.Context, owner, repo string) (*github.CodeOfConduct, *github.Response, error)
	GetCombinedStatus(ctx context.Context, owner, repo, ref string, opts *github.ListOptions) (*github.CombinedStatus, *github.Response, error)
	GetComment(ctx context.Context, owner, repo string, id int64) (*github.RepositoryComment, *github.Response, error)
	GetCommit(ctx context.Context, owner, repo, sha string) (*github.RepositoryCommit, *github.Response, error)
	GetCommitRaw(ctx context.Context, owner string, repo string, sha string, opts github.RawOptions) (string, *github.Response, error)
	GetCommitSHA1(ctx context.Context, owner, repo, ref, lastSHA string) (string, *github.Response, error)
	GetCommunityHealthMetrics(ctx context.Context, owner, repo string) (*github.CommunityHealthMetrics, *github.Response, error)
	GetContents(ctx context.Context, owner, repo, path string, opts *github.RepositoryContentGetOptions) (fileContent *github.RepositoryContent, directoryContent []*github.RepositoryContent, resp *github.Response, err error)
	GetDeployment(ctx context.Context, owner, repo string, deploymentID int64) (*github.Deployment, *github.Response, error)
	GetDeploymentStatus(ctx context.Context, owner, repo string, deploymentID, deploymentStatusID int64) (*github.DeploymentStatus, *github.Response, error)
	GetHook(ctx context.Context, owner, repo string, id int64) (*github.Hook, *github.Response, error)
	GetKey(ctx context.Context, owner string, repo string, id int64) (*github.Key, *github.Response, error)
	GetLatestPagesBuild(ctx context.Context, owner, repo string) (*github.PagesBuild, *github.Response, error)
	GetLatestRelease(ctx context.Context, owner, repo string) (*github.RepositoryRelease, *github.Response, error)
	GetPageBuild(ctx context.Context, owner, repo string, id int64) (*github.PagesBuild, *github.Response, error)
	GetPagesInfo(ctx context.Context, owner, repo string) (*github.Pages, *github.Response, error)
	GetPermissionLevel(ctx context.Context, owner, repo, user string) (*github.RepositoryPermissionLevel, *github.Response, error)
	GetPreReceiveHook(ctx context.Context, owner, repo string, id int64) (*github.PreReceiveHook, *github.Response, error)
	GetPullRequestReviewEnforcement(ctx context.Context, owner, repo, branch string) (*github.PullRequestReviewsEnforcement, *github.Response, error)
	GetReadme(ctx context.Context, owner, repo string, opts *github.RepositoryContentGetOptions) (*github.RepositoryContent, *github.Response, error)
	GetRelease(ctx context.Context, owner, repo string, id int64) (*github.RepositoryRelease, *github.Response, error)
	GetReleaseAsset(ctx context.Context, owner, repo string, id int64) (*github.ReleaseAsset, *github.Response, error)
	GetReleaseByTag(ctx context.Context, owner, repo, tag string) (*github.RepositoryRelease, *github.Response, error)
	GetRequiredStatusChecks(ctx context.Context, owner, repo, branch string) (*github.RequiredStatusChecks, *github.Response, error)
	GetSignaturesProtectedBranch(ctx context.Context, owner, repo, branch string) (*github.SignaturesProtectedBranch, *github.Response, error)
	GetVulnerabilityAlerts(ctx context.Context, owner, repository string) (bool, *github.Response, error)
	IsCollaborator(ctx context.Context, owner, repo, user string) (bool, *github.Response, error)
	License(ctx context.Context, owner, repo string) (*github.RepositoryLicense, *github.Response, error)
	List(ctx context.Context, user string, opts *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
	ListAll(ctx context.Context, opts *github.RepositoryListAllOptions) ([]*github.Repository, *github.Response, error)
	ListAllTopics(ctx context.Context, owner, repo string) ([]string, *github.Response, error)
	ListApps(ctx context.Context, owner, repo, branch string) ([]*github.App, *github.Response, error)
	ListBranches(ctx context.Context, owner string, repo string, opts *github.BranchListOptions) ([]*github.Branch, *github.Response, error)
	ListBranchesHeadCommit(ctx context.Context, owner, repo, sha string) ([]*github.BranchCommit, *github.Response, error)
	ListByOrg(ctx context.Context, org string, opts *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error)
	ListCodeFrequency(ctx context.Context, owner, repo string) ([]*github.WeeklyStats, *github.Response, error)
	ListCollaborators(ctx context.Context, owner, repo string, opts *github.ListCollaboratorsOptions) ([]*github.User, *github.Response, error)
	ListComments(ctx context.Context, owner, repo string, opts *github.ListOptions) ([]*github.RepositoryComment, *github.Response, error)
	ListCommitActivity(ctx context.Context, owner, repo string) ([]*github.WeeklyCommitActivity, *github.Response, error)
	ListCommitComments(ctx context.Context, owner, repo, sha string, opts *github.ListOptions) ([]*github.RepositoryComment, *github.Response, error)
	ListCommits(ctx context.Context, owner, repo string, opts *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error)
	ListContributors(ctx context.Context, owner string, repository string, opts *github.ListContributorsOptions) ([]*github.Contributor, *github.Response, error)
	ListContributorsStats(ctx context.Context, owner, repo string) ([]*github.ContributorStats, *github.Response, error)
	ListDeploymentStatuses(ctx context.Context, owner, repo string, deployment int64, opts *github.ListOptions) ([]*github.DeploymentStatus, *github.Response, error)
	ListDeployments(ctx context.Context, owner, repo string, opts *github.DeploymentsListOptions) ([]*github.Deployment, *github.Response, error)
	ListForks(ctx context.Context, owner, repo string, opts *github.RepositoryListForksOptions) ([]*github.Repository, *github.Response, error)
	ListHooks(ctx context.Context, owner, repo string, opts *github.ListOptions) ([]*github.Hook, *github.Response, error)
	ListInvitations(ctx context.Context, owner, repo string, opts *github.ListOptions) ([]*github.RepositoryInvitation, *github.Response, error)
	ListKeys(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Key, *github.Response, error)
	ListLanguages(ctx context.Context, owner string, repo string) (map[string]int, *github.Response, error)
	ListPagesBuilds(ctx context.Context, owner, repo string, opts *github.ListOptions) ([]*github.PagesBuild, *github.Response, error)
	ListParticipation(ctx context.Context, owner, repo string) (*github.RepositoryParticipation, *github.Response, error)
	ListPreReceiveHooks(ctx context.Context, owner, repo string, opts *github.ListOptions) ([]*github.PreReceiveHook, *github.Response, error)
	ListProjects(ctx context.Context, owner, repo string, opts *github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListPunchCard(ctx context.Context, owner, repo string) ([]*github.PunchCard, *github.Response, error)
	ListReleaseAssets(ctx context.Context, owner, repo string, id int64, opts *github.ListOptions) ([]*github.ReleaseAsset, *github.Response, error)
	ListReleases(ctx context.Context, owner, repo string, opts *github.ListOptions) ([]*github.RepositoryRelease, *github.Response, error)
	ListRequiredStatusChecksContexts(ctx context.Context, owner, repo, branch string) (contexts []string, resp *github.Response, err error)
	ListStatuses(ctx context.Context, owner, repo, ref string, opts *github.ListOptions) ([]*github.RepoStatus, *github.Response, error)
	ListTags(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.RepositoryTag, *github.Response, error)
	ListTeams(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Team, *github.Response, error)
	ListTrafficClones(ctx context.Context, owner, repo string, opts *github.TrafficBreakdownOptions) (*github.TrafficClones, *github.Response, error)
	ListTrafficPaths(ctx context.Context, owner, repo string) ([]*github.TrafficPath, *github.Response, error)
	ListTrafficReferrers(ctx context.Context, owner, repo string) ([]*github.TrafficReferrer, *github.Response, error)
	ListTrafficViews(ctx context.Context, owner, repo string, opts *github.TrafficBreakdownOptions) (*github.TrafficViews, *github.Response, error)
	Merge(ctx context.Context, owner, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error)
	OptionalSignaturesOnProtectedBranch(ctx context.Context, owner, repo, branch string) (*github.Response, error)
	PingHook(ctx context.Context, owner, repo string, id int64) (*github.Response, error)
	RemoveAdminEnforcement(ctx context.Context, owner, repo, branch string) (*github.Response, error)
	RemoveAppRestrictions(ctx context.Context, owner, repo, branch string, slug []string) ([]*github.App, *github.Response, error)
	RemoveBranchProtection(ctx context.Context, owner, repo, branch string) (*github.Response, error)
	RemoveCollaborator(ctx context.Context, owner, repo, user string) (*github.Response, error)
	RemovePullRequestReviewEnforcement(ctx context.Context, owner, repo, branch string) (*github.Response, error)
	ReplaceAllTopics(ctx context.Context, owner, repo string, topics []string) ([]string, *github.Response, error)
	ReplaceAppRestrictions(ctx context.Context, owner, repo, branch string, slug []string) ([]*github.App, *github.Response, error)
	RequestPageBuild(ctx context.Context, owner, repo string) (*github.PagesBuild, *github.Response, error)
	RequireSignaturesOnProtectedBranch(ctx context.Context, owner, repo, branch string) (*github.SignaturesProtectedBranch, *github.Response, error)
	TestHook(ctx context.Context, owner, repo string, id int64) (*github.Response, error)
	Transfer(ctx context.Context, owner, repo string, transfer github.TransferRequest) (*github.Repository, *github.Response, error)
	UpdateBranchProtection(ctx context.Context, owner, repo, branch string, preq *github.ProtectionRequest) (*github.Protection, *github.Response, error)
	UpdateComment(ctx context.Context, owner, repo string, id int64, comment *github.RepositoryComment) (*github.RepositoryComment, *github.Response, error)
	UpdateFile(ctx context.Context, owner, repo, path string, opts *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *github.Response, error)
	UpdateInvitation(ctx context.Context, owner, repo string, invitationID int64, permissions string) (*github.RepositoryInvitation, *github.Response, error)
	UpdatePages(ctx context.Context, owner, repo string, opts *github.PagesUpdate) (*github.Response, error)
	UpdatePreReceiveHook(ctx context.Context, owner, repo string, id int64, hook *github.PreReceiveHook) (*github.PreReceiveHook, *github.Response, error)
	UpdatePullRequestReviewEnforcement(ctx context.Context, owner, repo, branch string, patch *github.PullRequestReviewsEnforcementUpdate) (*github.PullRequestReviewsEnforcement, *github.Response, error)
	UpdateRequiredStatusChecks(ctx context.Context, owner, repo, branch string, sreq *github.RequiredStatusChecksRequest) (*github.RequiredStatusChecks, *github.Response, error)
	UploadReleaseAsset(ctx context.Context, owner, repo string, id int64, opts *github.UploadOptions, file *os.File) (*github.ReleaseAsset, *github.Response, error)
}

type RepositoriesService struct {
	RepositoriesServiceInterface
}