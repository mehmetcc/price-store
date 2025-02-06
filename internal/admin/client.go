package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mehmetcc/price-store/internal/config"
	"github.com/mehmetcc/price-store/internal/db"
	"github.com/mehmetcc/price-store/internal/websocket"
)

type Client struct {
	cfg      *config.Config
	wsClient *websocket.Client
}

func NewAdminClient(cfg config.Config, wsClient websocket.Client) *Client {
	return &Client{
		cfg:      &cfg,
		wsClient: &wsClient,
	}
}

func (c *Client) AddSymbol(symbol string) error {
	return c.wsClient.SendSymbol(symbol)
}

func (c *Client) GetSymbols() ([]string, error) {
	url := fmt.Sprintf("%s/symbol?client_id=%s", c.cfg.PricerUrl, c.cfg.ClientId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get symbols: %s", resp.Status)
	}

	var symbols []string
	if err := json.NewDecoder(resp.Body).Decode(&symbols); err != nil {
		return nil, err
	}

	return symbols, nil
}

func (c *Client) GetPriceUpdates(page, pageSize int) ([]db.PriceUpdate, error) {
	return db.GetPriceUpdates(page, pageSize)
}

func (c *Client) GetTotalPriceUpdatesCount() (int64, error) {
	return db.GetTotalPriceUpdatesCount()
}
