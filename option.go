package prokishi

type Config struct {
	Host     string
	Port     int
	Code     string
	EngineId string
	Version  string
}

var gConf *Config

func getConfig() *Config {
	return gConf
}

func setConfig(c *Config) {
	gConf = c
}

func getDefaultConfig(host string, port int) *Config {
	var conf Config
	conf.Host = host
	conf.Port = port
	conf.Code = ""
	conf.EngineId = ""
	return &conf
}

type Option func(*Config) error

func Code(code string) Option {
	return func(c *Config) error {
		c.Code = code
		return nil
	}
}

func Engine(id string) Option {
	return func(c *Config) error {
		c.EngineId = id
		return nil
	}
}

func Version(v string) Option {
	return func(c *Config) error {
		if v == "" {
			v = "Development"
		}
		c.Version = v
		return nil
	}
}
