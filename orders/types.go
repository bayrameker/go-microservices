package main

type OrderService interface {
	CreateOrder(context.Contect) error
}

type OrdersStore interface {
	Create(context.Context) error
}