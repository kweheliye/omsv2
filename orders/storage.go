package main

import "context"

type store struct {
	//add mongo db
}

func NewStore() *store {
	return &store{}
}

func (s *store) CreateOrder(ctx context.Context) error {
	return nil
}
