package rate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RateService struct {
}

func NewRateService() *RateService {
	return &RateService{}
}

type ExchangeRate struct {
	Rate float64 `json:"rate"`
}

func (r *RateService) GetRate() (float64, error) {
	const url = "https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?valcode=USD&json"

	resp, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("bad status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}

	var exchangeRates []ExchangeRate
	err = json.Unmarshal(body, &exchangeRates)
	if err != nil {
		return 0.0, err
	}

	if len(exchangeRates) == 0 {
		return 0.0, fmt.Errorf("no exchange rates found")
	}

	return exchangeRates[0].Rate, nil
}
