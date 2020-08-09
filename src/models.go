/*
Copyright © 2020 Tomasz Martin <tomasz.martin@gmail.com>

Here are all the model definitions.
*/
package gockify

type ResourceType string

const (
	Workspace ResourceType = "Workspace"
	Project   ResourceType = "Project"
	Task      ResourceType = "Task"
)

type Resource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
