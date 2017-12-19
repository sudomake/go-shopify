package goshopify

import "fmt"

// TransactionService is an interface for interfacing with the transactions endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/transaction
type TransactionService interface {
	List(int, interface{}) ([]Transaction, error)
	Count(int, interface{}) (int, error)
}

// TransactionServiceOp handles communication with the transaction related methods of the
// Shopify API.
type TransactionServiceOp struct {
	client *Client
}

// TransactionsResource represents the result from the transactions.json endpoint
type TransactionsResource struct {
	Transactions []Transaction `json:"transactions"`
}

// List transactions
func (s *TransactionServiceOp) List(orderID int, options interface{}) ([]Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions.json", ordersBasePath, orderID)
	resource := new(TransactionsResource)
	err := s.client.Get(path, resource, options)
	return resource.Transactions, err
}

// Count transactions
func (s *TransactionServiceOp) Count(orderID int, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%d/transactions/count.json", ordersBasePath, orderID)
	return s.client.Count(path, options)
}
