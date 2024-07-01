package services

import (
	"context"
	"hakaton2024/client/internal/broker"
	"hakaton2024/client/pkg/models"
	"strconv"
)

type ClientService struct {
	Broker *broker.Broker
}

func NewClientService(broker broker.Broker) *ClientService {
	return &ClientService{&broker}
}

func (cs *ClientService) ExecuteDeal(ctx context.Context, ticker string, dealType string, amount int, price int) error {
	deal := models.DealBody{
		Ticker: ticker,
		Type:   dealType,
		Amount: amount,
		Price:  price,
	}

	dealRequest := models.DealRequest{Deal: deal}

	_, err := cs.Broker.PostDeal(ctx, dealRequest)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ClientService) GetStatusOpenOrders(ctx context.Context) ([]models.OpenOrders, error) {
	statusInfo, err := cs.Broker.GetStatusInfo(ctx)
	if err != nil {
		return nil, err
	}
	return statusInfo.Body.OpenOrders, nil
}

func (cs *ClientService) GetStatusPositions(ctx context.Context) ([]models.Position, error) {
	statusInfo, err := cs.Broker.GetStatusInfo(ctx)
	if err != nil {
		return nil, err
	}
	for idx, val := range statusInfo.Body.Positions {
		val.Time = strconv.Itoa(idx)
	}
	return statusInfo.Body.Positions, nil
}

func (cs *ClientService) BidCancel(ctx context.Context, requestID string) (*models.BrokerCancelResponse, error) {
	reqID, err := strconv.Atoi(requestID)
	if err != nil {
		return nil, err
	}
	resp, err := cs.Broker.PostCancel(ctx, reqID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cs *ClientService) TickerData(ctx context.Context, ticker string) (*models.BrokerHistoryResponse, error) {
	history, err := cs.Broker.GetHistory(ctx, ticker)
	if err != nil {
		return history, err
	}

	return history, nil
}
