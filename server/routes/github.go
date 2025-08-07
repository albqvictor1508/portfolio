package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func GetGithubData(ctx *gin.Context) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if accessToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "GITHUB_ACCESS_TOKEN not set"})
		return
	}

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := github.NewClient(oauthClient)

	sinceStr := ctx.Query("since")
	var sinceTime time.Time
	var err error

	if sinceStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "'since' parameter is required"})
		return
	}

	sinceTime, err = time.Parse("2006-01-02", sinceStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'since' date format. Use YYYY-MM-DD."})
		return
	}

	user, _, err := client.Users.Get(context.Background(), "albqvictor1508")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get GitHub user", "details": err.Error()})
		return
	}
	username := *user.Login

	repos, _, err := client.Repositories.List(context.Background(), "", nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list repositories", "details": err.Error()})
		return
	}

	commitsByDay := make(map[string]int)

	for _, repo := range repos {
		commitOpts := &github.CommitsListOptions{
			Author:      username,
			Since:       sinceTime,
			ListOptions: github.ListOptions{PerPage: 100},
		}
		for {
			commits, resp, err := client.Repositories.ListCommits(context.Background(), username, *repo.Name, commitOpts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not fetch commits for repo %s: %v\n", *repo.Name, err)
				break
			}

			for _, commit := range commits {
				commitDate := commit.GetCommit().GetAuthor().GetDate()
				dateStr := commitDate.UTC().Format("2006-01-02")
				commitsByDay[dateStr]++
			}

			if resp.NextPage == 0 {
				break
			}
			commitOpts.Page = resp.NextPage
		}
	}

	ctx.JSON(http.StatusOK, commitsByDay)
}
