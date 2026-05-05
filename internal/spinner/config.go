package spinner

import (
	"fmt"
	"net/url"
	"strings"
)

type SpinnerConfig struct {
	APIHost string `env:"API_HOST" default:"http://localhost"`
	APIPort string `env:"API_PORT" default:"8080"`
}

func (c *SpinnerConfig) ServerURL() (string, error) {
	urlStr := fmt.Sprintf("%s:%s", c.APIHost, c.APIPort)
	if !strings.HasPrefix(urlStr, "http") {
		urlStr = "http://" + urlStr
	}
	serverURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	return serverURL.String(), nil
}
