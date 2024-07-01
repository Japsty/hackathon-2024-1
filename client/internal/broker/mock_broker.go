package broker

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"
)

func MockServer() string {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseBody interface{}
		switch r.URL.Path {
		case "/api/v1/history":
			ticker := r.URL.Query().Get("ticker")
			switch ticker {
			case "RTS":
				responseBody = map[string]interface{}{
					"body": map[string]interface{}{
						"ticker": "RTS",
						"prices": []map[string]interface{}{
							{
								"id":       1,
								"time":     time.Now().UnixNano() - 360,
								"interval": 1,
								"open":     5,
								"close":    3,
								"high":     5,
								"low":      99,
								"vol":      10,
							},
							{
								"id":       2,
								"time":     time.Now().UnixNano() - 60,
								"interval": 1,
								"open":     1,
								"close":    3,
								"high":     5,
								"low":      1,
								"vol":      10,
							},
							{
								"id":       3,
								"time":     time.Now().UnixNano(),
								"interval": 1,
								"open":     7,
								"close":    3,
								"high":     5,
								"low":      19,
								"vol":      10,
							},
						},
					},
				}
			case "IMOEX":
				responseBody = map[string]interface{}{
					"body": map[string]interface{}{
						"ticker": "IMOEX",
						"prices": []map[string]interface{}{
							{
								"id":       1,
								"time":     time.Now().UnixNano(),
								"interval": 1,
								"open":     100,
								"close":    110,
								"high":     120,
								"low":      90,
								"vol":      1000,
							},
						},
					},
				}
			}

		case "/api/v1/status":
			responseBody = map[string]interface{}{
				"body": map[string]interface{}{
					"balance": 1092,
					"positions": []map[string]interface{}{
						{
							"ticker": "RTS",
							"per":    "fbhdbfdn",
							"date":   "12.04.2024",
							"time":   "13:56",
							"last":   "110",
							"vol":    120,
						},
						{
							"ticker": "IMOEX",
							"per":    "fbhdbfdn",
							"date":   "12.04.2024",
							"time":   "13:56",
							"last":   "110",
							"vol":    120,
						},
					},
					"open_orders": []map[string]interface{}{
						{
							"ID":     1,
							"ticker": "dkjfgjmd",
							"per":    "dfjdnfd",
							"date":   "11.05.2024",
							"time":   "15:36",
							"last":   "120",
							"vol":    1000,
						},
					},
				},
			}
		default:
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		// Преобразуем данные в JSON и отправляем
		responseJSON, err := json.Marshal(responseBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
	}))

	return server.URL

}
