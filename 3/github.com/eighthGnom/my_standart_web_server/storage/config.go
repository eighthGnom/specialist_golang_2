package storage

type Config struct {
	StorageURI string `toml:"storage_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
