package config

import (
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Url struct{
    ServiceAddr string `toml:"service_addr"`
    UrlPrefix string `toml:"url_prefix"`
}

type Config struct{
    Port string `toml:"port"`
    Host string `toml:"host"`
    Urls []Url `toml:"urls"`
}

func New(path string) (*Config, error) {
    var cfg Config
    
    file, err := os.Open(path)
    if err != nil{
        return nil, err
    }

    body, err := io.ReadAll(file)
    if err != nil{
        return nil, err
    }

    err = toml.Unmarshal(body, &cfg)
    return &cfg, err
}