package main

import "context"

type OrdersService interface {
	CreateOrder(ctx context.Context) error
}

type OrdersStore interface {
	CreateOrder(ctx context.Context) error
}
