package main

import (
	"fmt"
	"encoding/json"
)

type Project struct {
	ID int64 `json:"id"`
	GithubUrl string `json:"github_url"`
	DemoUrl string `json:"demo_url"`
	Description string `json:"description"`
}

type Technology struct {
	ID int64 `json:"id"`
	ProjectId int64 `json:"project_id"`
	Name string `json:"name"`
}

func main() {
}
