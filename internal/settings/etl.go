package settings

type CurrencyAPISettings struct {
	ApiKey  string `mapstructure:"CURRENCY_API_KEY"`
	BaseUrl string `mapstructure:"CURRENCY_BASE_URL"`
}
