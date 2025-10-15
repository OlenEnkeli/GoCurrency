package repositories

import (
	"github.com/OlenEnkeli/GoCurrency/internal/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func NewPostgresConnection() *sqlx.DB {
	db, err := sqlx.Open("postgres", settings.Settings.Postgres.URI())
	if err != nil {
		logrus.Fatalf("Can`t connect to PostgreSQL DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatalf("Can`t connect to PostgreSQL DB: %v", err)
	}

	logrus.Info(settings.Settings.Postgres.URI())

	var dbName string

	err = db.Get(&dbName, "SELECT current_database();")
	if err != nil {
		logrus.Fatalf("Can`t connect to PostgreSQL DB: %v", err)
	}

	logrus.Infof("Connected to PostgreSQL database: %s", dbName)

	return db
}
