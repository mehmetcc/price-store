package routes

import (
	"encoding/json"
	"log"
	"net/http"

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
}
