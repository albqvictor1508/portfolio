package api

import (
	"net/http"
	"encoding/json"
	"time"
)

type Project struct {
	id int64
	categoryId int64
	name string
	description string
	githubUrl string
	demoUrl string
	createdAt time.Time
	updatedAt time.Time
}

type Technologies struct {
	id int64
	name string
	createdAt time.Time
	updatedAt time.Time
}

type ProjectTechnologies struct {
	id int64
	createdAt time.Time
	updatedAt time.Time
	projectId int64
}

type Categories struct {
	id int64
	name string
}

type Error struct {
	code int
	message string
}
