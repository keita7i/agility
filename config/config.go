package config

type Config struct {
}

func FromEnv() (Config, error) {
	return Config{}, nil
}
