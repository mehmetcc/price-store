package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/mehmetcc/price-store/internal/admin"
)

func SetupRoutes(mux *http.ServeMux, resolver *admin.Resolver) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	mux.HandleFunc("/symbol", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			return
		}

		switch r.Method {
		case http.MethodPost:
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
			if err := resolver.AddSymbol(req.Symbol); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "tracking initiated",
				"symbol": req.Symbol,
			})
		case http.MethodGet:
			symbols, err := resolver.GetSymbols()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"symbols": symbols,
			})
		case http.MethodDelete:
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
			if err := resolver.DeleteSymbol(req.Symbol); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "tracking removed",
				"symbol": req.Symbol,
			})
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/price", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			return
		}

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

		priceUpdates, err := resolver.GetPriceUpdates(page, pageSize)
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		count, err := resolver.GetTotalPriceUpdatesCount()
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
