package config

type Url struct{
    ServiceAddr string
    UrlPrefix string
}

type Config struct{
    port string
    host string
    Urls []Url
}