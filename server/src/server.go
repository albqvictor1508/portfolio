package main

import "fmt"

type Project struct {
	ID long `json:"id"` 
	GithubUrl string `json:"github_url"`
	//CreatedAt date `json:"created_at"`
}

func main() {
	server := "fake server my guy"
	fmt.Println(server)
}
