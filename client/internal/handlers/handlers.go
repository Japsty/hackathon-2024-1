package handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"hakaton2024/client/internal/broker"
	"hakaton2024/client/internal/services"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type ClientHandler struct {
	Tmpl    *template.Template
	Logger  *zap.SugaredLogger
	Service *services.ClientService
	Broker  *broker.Broker
}

func NewClientHandler(logger *zap.SugaredLogger, service *services.ClientService, broker *broker.Broker) *ClientHandler {
	return &ClientHandler{
		Logger:  logger,
		Service: service,
		Broker:  broker,
	}
}

var TimeoutTime = 500000 * time.Millisecond

func (c *ClientHandler) PostDeal(w http.ResponseWriter, r *http.Request) {
	ctxWthTimeout, cancel := context.WithTimeout(r.Context(), TimeoutTime)
	defer cancel()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Cannot parse form", http.StatusBadRequest)
		return
	}
	amountString := r.FormValue("amount")
	amount, err := strconv.Atoi(amountString)
	if err != nil {
		c.Logger.Errorf("Buy strconv.Atoi amount Error: %s", err)
		http.Error(w, "Cannot parse form", http.StatusBadRequest)
		return
	}
	priceString := r.FormValue("price")
	price, err := strconv.Atoi(priceString)
	if err != nil {
		c.Logger.Errorf("Buy strconv.Atoi price Error: %s", err)
		http.Error(w, "Cannot parse form", http.StatusBadRequest)
		return
	}
	instrument := r.FormValue("instrument")

	actionString := r.FormValue("action")

	if actionString != "SELL" && actionString != "BUY" {
		c.Logger.Errorf("Invalid action: %s", actionString)
		http.Error(w, "Invalid action", http.StatusBadRequest)
		return
	}

	if err = c.Service.ExecuteDeal(ctxWthTimeout, instrument, actionString, amount, price); err != nil {
		c.Logger.Errorw("Failed to buy", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	c.Logger.Infow("Purchase made", "instrument", instrument, "action", actionString, "amount", amount, "price", price)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *ClientHandler) GetOpenRequests(w http.ResponseWriter, r *http.Request) {
	ctxWthTimeout, cancel := context.WithTimeout(r.Context(), TimeoutTime)
	defer cancel()

	requests, err := c.Service.GetStatusOpenOrders(ctxWthTimeout)
	if err != nil {
		c.Logger.Errorw("GetOpenRequests Failed to get open requests", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(requests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Logger.Infoln("Posts marshaled")

	_, err = w.Write(resp)
	if err != nil {
		c.Logger.Errorln(err.Error())
	}
	_, err = w.Write([]byte("\n\n"))
	if err != nil {
		c.Logger.Errorln(err.Error())
	}
}

func (c *ClientHandler) GetOpenPositionsAndBalance(w http.ResponseWriter, r *http.Request) {
	ctxWthTimeout, cancel := context.WithTimeout(r.Context(), TimeoutTime)
	defer cancel()

	positions, err := c.Service.GetStatusPositions(ctxWthTimeout)
	if err != nil {
		c.Logger.Errorw("GetOpenPositionsAndBalance Failed to get open positions", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(positions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Logger.Infoln("Posts marshaled")

	_, err = w.Write(resp)
	if err != nil {
		c.Logger.Errorln(err.Error())
	}
	_, err = w.Write([]byte("\n\n"))
	if err != nil {
		c.Logger.Errorln(err.Error())
	}
}

func (c *ClientHandler) CancelRequest(w http.ResponseWriter, r *http.Request) {
	ctxWthTimeout, cancel := context.WithTimeout(r.Context(), TimeoutTime)
	defer cancel()

	id := mux.Vars(r)["ID"]
	if _, err := c.Service.BidCancel(ctxWthTimeout, id); err != nil {
		c.Logger.Errorw("CancelRequest Failed to cancel request", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *ClientHandler) GetBiddingHistory(w http.ResponseWriter, r *http.Request) {
	ctxWthTimeout, cancel := context.WithTimeout(r.Context(), TimeoutTime)
	defer cancel()

	ticker := mux.Vars(r)["TICKER"]
	history, err := c.Service.TickerData(ctxWthTimeout, ticker)
	if err != nil {
		c.Logger.Errorw("Failed to get bidding history", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(history)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Logger.Infoln("Posts marshaled")

	_, err = w.Write(resp)
	if err != nil {
		c.Logger.Errorln(err.Error())
	}
	_, err = w.Write([]byte("\n\n"))
	if err != nil {
		c.Logger.Errorln(err.Error())
	}
}
