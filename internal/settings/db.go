package settings

import (
	"fmt"
	"net"
)

type DBType string

const (
	Postgres DBType = "postgres"
	Mongo    DBType = "mongo"
)

type DBConfig interface {
	URI() string
}

type PostgresSettings struct {
	Host     string `mapstructure:"DB_POSTGRES_HOST"`
	Port     string `mapstructure:"DB_POSTGRES_PORT"`
	User     string `mapstructure:"DB_POSTGRES_USER"`
	Password string `mapstructure:"DB_POSTGRES_PASSWORD"`
	Database string `mapstructure:"DB_POSTGRES_DATABASE"`
	SSLMode  string `mapstructure:"DB_POSTGRES_SSL_MODE"`
}

func (s PostgresSettings) URI() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		s.Host,
		s.Port,
		s.User,
		s.Password,
		s.Database,
		s.SSLMode,
	)
}

type MongoSettings struct {
	Host     string `mapstructure:"DB_MONGO_HOST"`
	Port     string `mapstructure:"DB_MONGO_PORT"`
	User     string `mapstructure:"DB_MONGO_USER"`
	Password string `mapstructure:"DB_MONGO_PASSWORD"`
	Database string `mapstructure:"DB_MONGO_DATABASE"`
}

func (s MongoSettings) URI() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s/%s",
		s.User,
		s.Password,
		net.JoinHostPort(
			s.Host,
			s.Port,
		),
		s.Database,
	)
}

type DBSettings struct {
	DBType   DBType           `mapstructure:"DB_TYPE"`
	Mongo    MongoSettings    `mapstructure:",squash"`
	Postgres PostgresSettings `mapstructure:",squash"`
}
