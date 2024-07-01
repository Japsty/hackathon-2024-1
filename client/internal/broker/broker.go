package broker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hakaton2024/client/pkg/models"
	"io"
	"net/http"
	"strconv"
)

type Broker struct {
	UserID      int
	Login       string
	UserBalance int
	serverURL   string
}

func NewBroker(id int, userBalance int, serverURL string, login string) *Broker {
	return &Broker{UserID: id, UserBalance: userBalance, serverURL: serverURL, Login: login}
}

func (b *Broker) GetStatusInfo(ctx context.Context) (*models.BrokerStatusResponse, error) {
	url := fmt.Sprintf("%s/api/v1/status", b.serverURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	headerID := strconv.Itoa(b.UserID)
	//userBalance := strconv.Itoa(b.UserBalance)
	//
	//req.Header.Add("id", headerID)
	//req.Header.Add("balance", userBalance)
	req.Header.Add("Authorization", headerID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var status models.BrokerStatusResponse
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (b *Broker) PostDeal(ctx context.Context, dealRequest models.DealRequest) (*models.DealResponse, error) {
	marshaled, err := json.Marshal(dealRequest)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/deal", b.serverURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(marshaled))
	if err != nil {
		return nil, err
	}

	headerID := strconv.Itoa(b.UserID)
	userBalance := strconv.Itoa(b.UserBalance)

	req.Header.Add("id", headerID)
	req.Header.Add("balance", userBalance)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var deal models.DealResponse
	err = json.Unmarshal(body, &deal)
	if err != nil {
		return nil, err
	}

	return &deal, nil
}

func (b *Broker) PostCancel(ctx context.Context, id int) (*models.BrokerCancelResponse, error) {
	reqBody := models.BrokerCancelRequest{ID: id}

	marshaled, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/cancel", b.serverURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(marshaled))
	if err != nil {
		return nil, err
	}

	headerID := strconv.Itoa(b.UserID)
	userBalance := strconv.Itoa(b.UserBalance)

	req.Header.Add("id", headerID)
	req.Header.Add("balance", userBalance)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cancel models.BrokerCancelResponse
	err = json.Unmarshal(body, &cancel)
	if err != nil {
		return nil, err
	}

	return &cancel, nil
}

func (b *Broker) GetHistory(ctx context.Context, ticker string) (*models.BrokerHistoryResponse, error) {
	//url := MockServer()
	reqUrl := fmt.Sprintf("%s/api/v1/history?ticker=%s", b.serverURL, ticker)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	headerID := strconv.Itoa(b.UserID)
	userBalance := strconv.Itoa(b.UserBalance)

	req.Header.Add("id", headerID)
	req.Header.Add("balance", userBalance)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var history models.BrokerHistoryResponse
	err = json.Unmarshal(body, &history)
	if err != nil {
		return nil, err
	}

	return &history, nil
}
