package gockify

import (
	"fmt"
	"net/http"
)

func (c *Client) request(req *http.Request) {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Api-Key", c.apiKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	c.printResponse(resp)
}

func (c *Client) GetWorkspaces() {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/workspaces", c.BaseURL), nil)
	if err != nil {
		fmt.Println(err)
	}
	c.request(req)
}
