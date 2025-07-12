package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
)
/*
var role {} = {
	BACK: "back-end"
	FRONT: "front-end"
}
*/

type Project struct {
	ID int64 `json:"id"` 
	GithubUrl string `json:"github_url"`
	DemoUrl string `json:"demo_url"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
}

type Technology struct {
	ID int64 `json:"id"`
	ProjectId long `json:"project_id"`
}

func main() {
	server := "fake server my guy"
	fmt.Println(server)

	log.Fatal(http.ListenAndServe(":3333", nil))
}
