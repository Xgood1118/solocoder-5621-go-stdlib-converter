package internal

import (
	"encoding/json"
	"math/big"
	"os"
)

type Config struct {
	Precision     int               `json:"precision"`
	CurrencyRates map[string]float64 `json:"currency_rates"`
	Locale        string            `json:"locale"`
	DecimalSep    string            `json:"decimal_separator"`
	ThousandSep   string            `json:"thousand_separator"`
}

var cfg *Config = DefaultConfig()

func DefaultConfig() *Config {
	return &Config{
		Precision:     6,
		Locale:        "en",
		DecimalSep:    ".",
		ThousandSep:   ",",
		CurrencyRates: nil,
	}
}

func GetConfig() *Config {
	if cfg == nil {
		cfg = DefaultConfig()
	}
	return cfg
}

func SetConfig(c *Config) {
	cfg = c
	if c.CurrencyRates != nil {
		reg := DefaultRegistry()
		for key, rate := range c.CurrencyRates {
			if u, ok := reg.units[key]; ok {
				u.Factor = big.NewFloat(rate)
			}
		}
	}
}

func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}
	SetConfig(&c)
	return nil
}

func SaveConfig(path string) error {
	data, err := json.MarshalIndent(GetConfig(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
