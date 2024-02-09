package server

type Config struct {
	Version string
}

type Option func(*Config) error

func Version(v string) Option {
	return func(c *Config) error {
		if v == "" {
			c.Version = "Development"
		} else {
			c.Version = v
		}
		return nil
	}
}
