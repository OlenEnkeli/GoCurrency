package currency_freak

import (
	"strings"

	"github.com/OlenEnkeli/GoCurrency/internal/settings"
)

type CurrencyAPI struct {
	ApiKey  string
	BaseUrl string
}

func NewCurrencyAPI() *CurrencyAPI {
	return &CurrencyAPI{
		ApiKey:  settings.Settings.CurrencyAPI.ApiKey,
		BaseUrl: settings.Settings.CurrencyAPI.BaseUrl,
	}
}

func (api *CurrencyAPI) GetLatestPairs(pairs []string) map[string]string {
	parsedPairs := strings.Join(pairs, ",")
	parsedUrl :=
}
