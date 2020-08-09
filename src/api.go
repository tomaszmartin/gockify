/*
Copyright © 2020 Tomasz Martin <tomasz.martin@gmail.com>

Here we store all the possible methods that can be called
from the API. The logic of calling them, error handling etc.
is in client.go.
*/
package gockify

import (
	"fmt"
)

func (c *Client) GetWorkspaces() ([]Resource, error) {
	url := fmt.Sprintf("%s/workspaces", c.BaseURL)
	return c.apiGet(url)
}

func (c *Client) GetProjects(workspaceId string) ([]Resource, error) {
	url := fmt.Sprintf("%s/workspaces/%s/projects", c.BaseURL, workspaceId)
	return c.apiGet(url)
}

func (c *Client) GetTasks(workspaceId string, projectId string) ([]Resource, error) {
	url := fmt.Sprintf("%s/workspaces/%s/projects/%s/tasks", c.BaseURL, workspaceId, projectId)
	return c.apiGet(url)
}
