/*
Copyright © 2020 Tomasz Martin <tomasz.martin@gmail.com>

Here is the logic of calling the API, erro handling etc.
*/
package gockify

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const BaseURL = "https://api.clockify.me/api/v1"
const ReportURL = "https://reports.api.clockify.me/v1"

type Client struct {
	BaseURL    string
	ReportURL  string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL:   BaseURL,
		ReportURL: ReportURL,
		apiKey:    apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) getResources(resp *http.Response) ([]Resource, error) {
	var decoded []Resource
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Bad Response", err)
	}
	err = json.Unmarshal(bodyBytes, &decoded)
	if err != nil {
		log.Fatal("Bad JSON", err)
	}
	return decoded, nil
}

func (c *Client) request(req *http.Request) *http.Response {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Api-Key", c.apiKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal("Bad response", err)
	}
	return resp
}

func (c *Client) apiGet(url string) ([]Resource, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Bad request", err)
	}
	resp := c.request(req)
	defer resp.Body.Close()
	return c.getResources(resp)
}

func (c *Client) apiPost(url string, values map[string]string) ([]Resource, error) {
	jsonValue, err := json.Marshal(values)
	if err != nil {
		log.Fatal("Bad JSON", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal("Bad request", err)
	}
	resp := c.request(req)
	defer resp.Body.Close()
	return c.getResources(resp)
}
