package main

import "context"

type service struct {
	store OrdersStore 
}

func newService(store OrdersStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context){
	return 
}