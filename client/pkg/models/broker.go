package models

import "context"

type BrokerInterface interface {
	GetStatusInfo(ctx context.Context) (*BrokerStatusResponse, error)
	PostDeal(ctx context.Context, brokerCancel BrokerCancelRequest) (*DealResponse, error)
	PostCancel(ctx context.Context, id int) (*BrokerCancelResponse, error)
	GetHistory(ctx context.Context, ticker string) (*BrokerHistoryResponse, error)
}

//Status Broker request

type BrokerStatusResponse struct {
	Body struct {
		Balance    int          `json:"balance"`
		Positions  []Position   `json:"positions"`
		OpenOrders []OpenOrders `json:"open_orders"`
	} `json:"body"`
}

type Position struct {
	Ticker string `json:"ticker,omitempty"`
	Per    string `json:"per,omitempty"`
	Date   string `json:"date,omitempty"`
	Time   string `json:"time,omitempty"`
	Last   string `json:"last,omitempty"`
	Vol    int    `json:"vol,omitempty"`
}

type OpenOrders struct {
	ID     int    `json:"ID,omitempty"`
	Ticker string `json:"ticker"`
	Per    string `json:"per,omitempty"`
	Date   string `json:"date,omitempty"`
	Time   string `json:"time,omitempty"`
	Last   string `json:"last,omitempty"`
	Vol    int    `json:"vol,omitempty"`
	Type   string `json:"type,omitempty"`
	Amount int    `json:"amount"`
	Price  int    `json:"price"`
}

//History Broker request

type BrokerHistoryResponse struct {
	Body struct {
		Ticker string   `json:"ticker"`
		Prices []Prices `json:"prices"`
	} `json:"body"`
}

type Prices struct {
	ID       int   `json:"id"`
	Time     int64 `json:"time"`
	Interval int   `json:"interval"`
	Open     int   `json:"open"`
	High     int   `json:"high"`
	Low      int   `json:"low"`
	Close    int   `json:"close"`
	Volume   int   `json:"vol"`
}

//Deal Broker

type DealRequest struct {
	Deal DealBody `json:"deal"`
}

type DealBody struct {
	Ticker string `json:"ticker"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`
	Price  int    `json:"price"`
}

type DealResponse struct {
	Body struct {
		ID string `json:"id"`
	} `json:"body"`
}

//Cancel Broker

type BrokerCancelRequest struct {
	ID int `json:"ID"`
}

type BrokerCancelResponse struct {
	Body struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	} `json:"body"`
}
