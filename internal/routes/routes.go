package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/mehmetcc/price-store/internal/admin"
)

func SetupRoutes(mux *http.ServeMux, client *admin.Client) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	mux.HandleFunc("/symbol", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			type request struct {
				Symbol string `json:"symbol"`
			}

			var req request
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "cannot parse request body", http.StatusBadRequest)
				return
			}
			if req.Symbol == "" {
				http.Error(w, "symbol is required", http.StatusBadRequest)
				return
			}

			if err := client.AddSymbol(req.Symbol); err != nil {
				http.Error(w, "failed to send symbol to websocket", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "tracking initiated",
				"symbol": req.Symbol,
			})
		} else if r.Method == http.MethodGet {
			symbols, err := client.GetSymbols()
			if err != nil {
				http.Error(w, "failed to retrieve symbols", http.StatusInternalServerError)
				log.Println(err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"symbols": symbols,
			})
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/price", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		page := 1
		pageSize := 100
		if r.URL.Query().Get("page") != "" {
			var err error
			page, err = strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				http.Error(w, "Invalid page parameter", http.StatusBadRequest)
				return
			}
		}
		if r.URL.Query().Get("pageSize") != "" {
			var err error
			pageSize, err = strconv.Atoi(r.URL.Query().Get("pageSize"))
			if err != nil {
				http.Error(w, "Invalid pageSize parameter", http.StatusBadRequest)
				return
			}
		}

		priceUpdates, err := client.GetPriceUpdates(page, pageSize)
		if err != nil {
			http.Error(w, "failed to retrieve price updates", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"price_updates": priceUpdates,
		})
	})

	mux.HandleFunc("/price/count", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		count, err := client.GetTotalPriceUpdatesCount()
		if err != nil {
			http.Error(w, "failed to retrieve price updates count", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"count": count,
		})
	})
}
