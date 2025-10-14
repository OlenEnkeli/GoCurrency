package settings

import "time"

type APISettings struct {
	Host           string        `default:"127.0.0.1" mapstructure:"API_HOST"`
	Port           string        `default:"8000"      mapstructure:"API_PORT"`
	Version        string        `default:"0.0.1"     mapstructure:"API_VERSION"`
	MaxHeaderBytes int           `default:"1024"      mapstructure:"API_MAX_HEADER_BYTES"`
	ReadTimeout    time.Duration `default:"10s"       mapstructure:"API_READ_TIMEOUT"`
	WriteTimeout   time.Duration `default:"10s"       mapstructure:"API_WRITE_TIMEOUT"`
}
