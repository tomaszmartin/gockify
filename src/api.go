package gockify

import (
	"fmt"
	"log"
	"net/http"
)

func (c *Client) request(req *http.Request, err error) (*http.Response, error) {
	if err != nil {
		log.Fatal("Bad request", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Api-Key", c.apiKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal("Bad response", err)
	}
	return resp, nil
}

func (c *Client) GetWorkspaces() ([]Resource, error) {
	url := fmt.Sprintf("%s/workspaces", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	resp, err := c.request(req, err)
	return c.getResources(resp, err)
}

func (c *Client) GetProjects(workspaceId string) ([]Resource, error) {
	url := fmt.Sprintf("%s/workspaces/%s/projects", c.BaseURL, workspaceId)
	req, err := http.NewRequest("GET", url, nil)
	resp, err := c.request(req, err)
	return c.getResources(resp, err)
}

func (c *Client) GetTasks(workspaceId string, projectId string) ([]Resource, error) {
	url := fmt.Sprintf("%s/workspaces/%s/projects/%s/tasks", c.BaseURL, workspaceId, projectId)
	req, err := http.NewRequest("GET", url, nil)
	resp, err := c.request(req, err)
	return c.getResources(resp, err)
}
