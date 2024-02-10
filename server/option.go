package server

type Config struct {
	Version string
}

const (
	DevelopmentVersion = "Development"
)

func (c *Config) isDev() bool {
	if c.Version == DevelopmentVersion {
		return true
	}
	return false
}

type Option func(*Config) error

func Version(v string) Option {
	return func(c *Config) error {
		if v == "" {
			c.Version = DevelopmentVersion
		} else {
			c.Version = v
		}
		return nil
	}
}
