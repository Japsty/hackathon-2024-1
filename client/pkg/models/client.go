package models

import (
	"context"
)

type Client struct {
	Login    string `json:"username"`
	ID       string `json:"id"`
	Password string `json:"-"`
}

type ClientRepo interface {
	ExecuteDeal(ctx context.Context) error
	GetStatus(ctx context.Context) error
	BidCancel(ctx context.Context) error
	TicketData(ctx context.Context, ticket string) error
}
