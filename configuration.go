package ecc

import (
	"net/http"
)

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
	AccessToken   string
	APIKey        string
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://napi.arvancloud.com/ecc/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "arvancloud/go-ecc-client",
		AccessToken:   "",
		APIKey:        "",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
