package main

import (
	"fmt"
	"errors"
	//"encoding/json"	
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
	const salve string = "salve"
	fmt.Println("Server Running on 3333!")

	var result, remainder, error int = returnTwoTimes(10, 2)
	fmt.Printf("The result: %v, and the remainder: %v", result, remainder)
}
