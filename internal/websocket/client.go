package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mehmetcc/price-store/internal/config"
	"github.com/mehmetcc/price-store/internal/db"
)

type Client struct {
	conn *websocket.Conn
	url  string
}

func NewClient(cfg *config.Config) (*Client, error) {
	headers := make(http.Header)
	headers.Add("X-Client-ID", cfg.ClientId)

	c, _, err := websocket.DefaultDialer.Dial(cfg.WsUrl, headers)
	if err != nil {
		log.Fatalf("failed to connect to websocket: %v", err)
	}

	return &Client{
		conn: c,
		url:  cfg.WsUrl,
	}, nil
}

func (c *Client) SendSymbol(symbol string) error {
	return c.conn.WriteMessage(websocket.TextMessage, []byte(symbol))
}

func (c *Client) Listen() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("websocket read error:", err)
			// TODO implement reconnect logic
			time.Sleep(2 * time.Second) // Fetch every 2 seconds
			continue
		}

		data, err := c.unmarshall(message)
		if err != nil {
			continue
		}
		c.persist(data)
	}
}

func (c *Client) unmarshall(message []byte) (map[string]float64, error) {
	var data map[string]float64
	log.Printf("Received message: %s", string(message))
	if err := json.Unmarshal(message, &data); err != nil {
		log.Println("error unmarshalling websocket message:", err)
		return nil, err
	}
	return data, nil
}

func (c *Client) persist(data map[string]float64) {
	// TODO possible N + 1 here
	// TODO implement batch saves in the future
	var wg sync.WaitGroup

	for symbol, price := range data {
		pu := db.NewPriceUpdate(symbol, price)
		wg.Add(1)

		go func(pu *db.PriceUpdate) {
			defer wg.Done()
			if err := db.Create(pu); err != nil {
				log.Println("error saving price update:", err)
			}
		}(pu)
	}

	wg.Wait()
}

func (c *Client) Close() {
	_ = c.conn.Close()
}
