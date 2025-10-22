package settings

import (
	"log"

	"github.com/spf13/viper"
)

type AllSettings struct {
	App         AppSettings         `mapstructure:",squash"`
	API         APISettings         `mapstructure:",squash"`
	Postgres    PostgresSettings    `mapstructure:",squash"`
	CurrencyAPI CurrencyAPISettings `mapstructure:",squash"`
}

var Settings AllSettings

func Init() {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Unable to load config: %s", err)
	}

	err = viper.Unmarshal(&Settings)

	if err != nil {
		log.Fatalf("Unable to load config: %s", err)
	}
}
