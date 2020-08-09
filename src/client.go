package gockify

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (c *Client) printJSON(resp *http.Response) {
	var prettyJSON bytes.Buffer
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Indent(&prettyJSON, bodyBytes, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prettyJSON.Bytes())
}

func (c *Client) getResources(resp *http.Response, err error) ([]Resource, error) {
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Bad response", err)
	}
	var decoded []Resource
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &decoded)
	if err != nil {
		log.Fatal("Bad JSON", err)
	}
	return decoded, nil
}
