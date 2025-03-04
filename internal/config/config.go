package config

import (
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Url struct{
    ServiceAddr string
    UrlPrefix string
}

type Config struct{
    Port string
    Host string
    Urls []Url
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