package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/watchtowerai/nightfall_dlp/internal/clients/diffreviewer"
	"github.com/watchtowerai/nightfall_dlp/internal/clients/diffreviewer/github"
	"github.com/watchtowerai/nightfall_dlp/internal/clients/nightfall"
)

const (
	nightfallConfigFileName = ".nightfalldlp/config.json"
	githubActionsEnvVar     = "GITHUB_ACTIONS"
	githubTokenEnvVar       = "NIGHTFALL_GITHUB_TOKEN"
)

// main starts the service process.
func main() {
	fmt.Println("Running NightfallDLP Action")
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "nightfalldlp: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done NightfallDLP Action")
}

func run() error {
	ctx := context.Background()
	diffReviewClient, err := CreateDiffReviewerClient(&http.Client{})
	if err != nil {
		return err
	}

	nightfallConfig, err := diffReviewClient.LoadConfig(nightfallConfigFileName)
	if err != nil {
		return err
	}
	nightfallClient := nightfall.NewClient(*nightfallConfig)

	fileDiffs, err := diffReviewClient.GetDiff()
	if err != nil {
		return err
	}

	comments, err := nightfallClient.ReviewDiff(ctx, diffReviewClient.GetLogger(), fileDiffs)
	if err != nil {
		return err
	}

	return diffReviewClient.WriteComments(comments)
}

// usingGithubAction determine if nightfalldlp is being run by
// Github Actions
func usingGithubAction() bool {
	val, ok := os.LookupEnv(githubActionsEnvVar)
	if !ok {
		return false
	}
	return val == "true"
}

// CreateDiffReviewerClient determines the current environment that is running nightfalldlp
// and returns the corresponding DiffReviewer client
func CreateDiffReviewerClient(httpClient *http.Client) (diffreviewer.DiffReviewer, error) {
	switch {
	case usingGithubAction():
		githubToken, ok := os.LookupEnv(githubTokenEnvVar)
		if !ok {
			return nil, errors.New("missing github token in env")
		}
		return github.NewAuthenticatedGithubService(githubToken), nil
	default:
		return nil, errors.New("current environment unknown")
	}
}
