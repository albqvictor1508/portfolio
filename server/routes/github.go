package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

const githubBaseURL string = "https://api.github.com/user"

func GetGithubData(ctx *gin.Context) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, tokenSource)
	client := github.NewClient(tc)

	response := http.Post(
		fmt.Sprintf(githubBaseURL, "/user"),
		"salve",
		"salve",
	)
}
