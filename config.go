package main

import (
	"net/http"
	"time"
)

type Config struct {
	APIURL   string
	APIToken string
}

type Client struct {
	httpClient *http.Client
	apiURL     string
	apiToken   string
}

func (c *Config) Client() (*Client, error) {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	return &Client{
		httpClient: client,
		apiURL:     c.APIURL,
		apiToken:   c.APIToken,
	}, nil
} 