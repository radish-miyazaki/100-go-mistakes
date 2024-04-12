package client

import "github.com/radish-miyazaki/100-go-mistakes/02-code-project-organization/06-interface-producer/store"

type customersGetter interface {
	GetAllCustomers() ([]store.Customer, error)
}
