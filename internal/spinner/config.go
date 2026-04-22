package spinner

import (
	"fmt"
	"net/url"
	"strings"
)

type SpinnerConfig struct {
	ServerHost string `env:"SERVER_HOST" default:"http://localhost"`
	ServerPort string `env:"SERVER_PORT" default:"8080"`
}

func (c *SpinnerConfig) ServerURL() (string, error) {
	urlStr := fmt.Sprintf("%s:%s", c.ServerHost, c.ServerPort)
	if !strings.HasPrefix(urlStr, "http") {
		urlStr = "http://" + urlStr
	}
	serverURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	return serverURL.String(), nil
}
