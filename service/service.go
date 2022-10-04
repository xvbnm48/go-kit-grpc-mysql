package service

import (
	"context"

	"github.com/go-kit/kit/log"
	model "github.com/xvbnm48/go-kit-grpc-mysql/model"
)

type service struct {
	Logger log.Logger
}

type Service interface {
	CreateCustomer(ctx context.Context, c model.Customer) error
	GetCustomerById(ctx context.Context, id int32) (model.Customer, error)
	UpdateCustomer(ctx context.Context, c model.Customer) error
	DeleteCustomer(ctx context.Context, id DeleteCustomerRequest) error
}

type DeleteCustomerRequest struct {
	ID int32
}
