package main

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"hakaton2024/client/configs"
	"hakaton2024/client/internal/broker"
	"hakaton2024/client/internal/handlers"
	"hakaton2024/client/internal/services"
	"log"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	html, err := os.ReadFile("../static/html1.html")
	if err != nil {
		// Обработка ошибки, если файл не найден
		http.Error(w, "File not found", 404)
		return
	}

	// Установка Content-Type
	w.Header().Set("Content-Type", "text/html")

	// Отправка содержимого файла
	_, err = w.Write(html)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	config, err := configs.LoadConfig("../config.yml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Printf("Error making new logger: %v", err)
	}
	defer func() {
		if err = zapLogger.Sync(); err != nil {
			log.Fatalf("Failed to sync logger: %v", err)
		}
	}()
	logger := zapLogger.Sugar()

	brokerService := broker.NewBroker(config.User.UserID, config.User.Balance, config.User.URL, config.User.Login)
	clientService := services.NewClientService(*brokerService)
	clientHandler := handlers.NewClientHandler(logger, clientService, brokerService)

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static/"))))

	r.HandleFunc("/", homeHandler).Methods(http.MethodGet)
	r.PathPrefix("/submit").HandlerFunc(clientHandler.PostDeal).Methods(http.MethodPost)
	r.PathPrefix("/requests").HandlerFunc(clientHandler.GetOpenRequests).Methods(http.MethodGet)
	r.PathPrefix("/positions").HandlerFunc(clientHandler.GetOpenPositionsAndBalance).Methods(http.MethodGet)
	r.PathPrefix("/data/{TICKER}").HandlerFunc(clientHandler.GetBiddingHistory).Methods(http.MethodGet)
	r.PathPrefix("/cancel/{ID}").HandlerFunc(clientHandler.CancelRequest).Methods(http.MethodPost)

	log.Println("Starting server at :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		println(err)
	}
}
