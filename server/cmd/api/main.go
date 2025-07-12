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

func returnTwoTimes(number1 int, number2 int) (int, int) {
	return number1 / number2, number1 % number2
}

var result, remainder int = returnTwoTimes(10, 2)

func main() {
	const salve string = "salve"
	fmt.Println("Server Running on 3333!")
}
