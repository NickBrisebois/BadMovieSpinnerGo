package web

type WebConfig struct {
	WebHost string `env:"WEB_HOST" default:"http://localhost"`
	WebPort string `env:"WEB_PORT" default:"8000"`
}
