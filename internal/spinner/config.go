package spinner

import "net/url"

type SpinnerConfig struct {
	ServerHost string `env:"SERVER_HOST"`
	ServerPort string `env:"SERVER_PORT"`
}

func (c *SpinnerConfig) ServerURL() (string, error) {
	if c.ServerHost == "" {
		c.ServerHost = "http://localhost"
	}
	if c.ServerPort == "" {
		c.ServerPort = "8080"
	}
	serverURL, err := url.Parse(c.ServerHost + ":" + c.ServerPort)
	if err != nil {
		return "", err
	}
	return serverURL.String(), nil
}
