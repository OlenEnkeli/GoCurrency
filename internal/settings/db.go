package settings

import (
	"fmt"
)

type DBConfig interface {
	URI() string
}

type PostgresSettings struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Database string `mapstructure:"POSTGRES_DATABASE"`
	SSLMode  string `mapstructure:"POSTGRES_SSL_MODE"`
}

func (s PostgresSettings) URI() string {
	result := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s",
		s.Host,
		s.Port,
		s.User,
		s.Database,
		s.SSLMode,
	)
	if s.Password != "" {
		result = fmt.Sprintf("%s password=%s", result, s.Password)
	}

	return result
}
