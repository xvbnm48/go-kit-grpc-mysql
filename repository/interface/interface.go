package _interface

import (
	"context"
	"io"

	model "github.com/xvbnm48/go-kit-grpc-mysql/model"
)

type ReadWriter interface {
	io.Closer
	CreateCustomer(ctx context.Context, c model.Customer) error
	GetCustomerById(ctx context.Context, id int32) (model.Customer, error)
	UpdateCustomer(ctx context.Context, c model.Customer) error
	DeleteCustomer(ctx context.Context, id DeleteCustomerRequest) error
}
