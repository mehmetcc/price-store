package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mehmetcc/price-store/internal/config"
	"github.com/mehmetcc/price-store/internal/db"
)

type Resolver struct {
	cfg *config.Config
}

func NewAdminResolver(cfg config.Config) *Resolver {
	return &Resolver{
		cfg: &cfg,
	}
}

func (c *Resolver) AddSymbol(symbol string) error {
	url := c.cfg.PricerUrl + "/symbol"

	payload := map[string]string{
		"symbol": symbol,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	if c.cfg.ClientId != "" {
		req.Header.Set("X-Client-ID", c.cfg.ClientId)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s", string(respBody))
	}

	return nil
}

func (c *Resolver) GetSymbols() ([]string, error) {
	url := c.cfg.PricerUrl + "/symbol"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if c.cfg.ClientId != "" {
		req.Header.Set("X-Client-ID", c.cfg.ClientId)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(respBody))
	}

	var symbols []string
	if err := json.NewDecoder(resp.Body).Decode(&symbols); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return symbols, nil
}

func (c *Resolver) DeleteSymbol(symbol string) error {
	url := c.cfg.PricerUrl + "/symbol"

	payload := map[string]string{
		"symbol": symbol,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.cfg.ClientId != "" {
		req.Header.Set("X-Client-ID", c.cfg.ClientId)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s", string(respBody))
	}

	return nil
}

func (c *Resolver) GetPriceUpdatesBySymbol(symbol string, page, pageSize int) ([]db.PriceUpdate, error) {
	return db.SearchPriceUpdatesBySymbol(symbol, page, pageSize)
}

func (c *Resolver) GetPriceUpdates(page, pageSize int) ([]db.PriceUpdate, error) {
	return db.GetPriceUpdates(page, pageSize)
}

func (c *Resolver) GetTotalPriceUpdatesCount() (int64, error) {
	return db.GetTotalPriceUpdatesCount()
}

func (c *Resolver) GetFilteredPriceUpdatesCount(symbol string) (int64, error) {
	return db.GetFilteredPriceUpdatesCount(symbol)
}
