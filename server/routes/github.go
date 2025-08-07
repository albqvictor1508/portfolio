package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

const githubBaseURL string = "https://api.github.com/user"

func GetGithubData(ctx *gin.Context, r *http.Request) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, tokenSource)
	client := github.NewClient(tc)

	sinceStr := r.URL.Query().Get("since")
	untilStr := r.URL.Query().Get("until")

	if untilStr == "" && sinceStr != "" {
		since, err := time.Parse("2006-01-02", sinceStr)
		if err == nil {
			until := since.AddDate(0, 3, 0)
			untilStr = until.Format("2006-01-02")
		}
	}

	response := http.Post(
		fmt.Sprintf(githubBaseURL, "/user"),
		"salve",
		"salve",
	)
}
