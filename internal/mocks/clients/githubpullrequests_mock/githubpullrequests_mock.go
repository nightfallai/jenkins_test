// Code generated by MockGen. DO NOT EDIT.
// Source: ../githubintf/github_pullrequests.go

// Package githubpullrequests_mock is a generated GoMock package.
package githubpullrequests_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v33/github"
	reflect "reflect"
)

// GithubPullRequests is a mock of GithubPullRequests interface
type GithubPullRequests struct {
	ctrl     *gomock.Controller
	recorder *GithubPullRequestsMockRecorder
}

// GithubPullRequestsMockRecorder is the mock recorder for GithubPullRequests
type GithubPullRequestsMockRecorder struct {
	mock *GithubPullRequests
}

// NewGithubPullRequests creates a new mock instance
func NewGithubPullRequests(ctrl *gomock.Controller) *GithubPullRequests {
	mock := &GithubPullRequests{ctrl: ctrl}
	mock.recorder = &GithubPullRequestsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *GithubPullRequests) EXPECT() *GithubPullRequestsMockRecorder {
	return m.recorder
}

// CreateComment mocks base method
func (m *GithubPullRequests) CreateComment(ctx context.Context, owner, repo string, number int, comment *github.PullRequestComment) (*github.PullRequestComment, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, owner, repo, number, comment)
	ret0, _ := ret[0].(*github.PullRequestComment)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateComment indicates an expected call of CreateComment
func (mr *GithubPullRequestsMockRecorder) CreateComment(ctx, owner, repo, number, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*GithubPullRequests)(nil).CreateComment), ctx, owner, repo, number, comment)
}

// ListComments mocks base method
func (m *GithubPullRequests) ListComments(ctx context.Context, owner, repo string, number int, opts *github.PullRequestListCommentsOptions) ([]*github.PullRequestComment, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComments", ctx, owner, repo, number, opts)
	ret0, _ := ret[0].([]*github.PullRequestComment)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListComments indicates an expected call of ListComments
func (mr *GithubPullRequestsMockRecorder) ListComments(ctx, owner, repo, number, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComments", reflect.TypeOf((*GithubPullRequests)(nil).ListComments), ctx, owner, repo, number, opts)
}
