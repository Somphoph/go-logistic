package domain

import "time"

type Sale struct {
	SaleNo   string    `json:"sale-no"`
	SaleDate time.Time `json:"sale-date"`
	Customer string    `json:"customer"`
	Product  string    `json:"product"`
	Qty      int       `json:"qty"`
	Price    float32   `json:"price"`
	ReqDate  time.Time `json:"req-date"`
}
type Invoice struct {
	InvNo     string    `json:"inv-no"`
	InvDate   time.Time `json:"inv-date"`
	Customer  string    `json:"customer"`
	Product   string    `json:"product"`
	Qty       int       `json:"qty"`
	Price     float32   `json:"price"`
	ReceiveNo string    `json:"receive-no"`
	SaleNo    string    `json:"sale-no"`
}
type SaleSvc interface {
	Get(saleNo string) (*Sale, error)
	List() ([]*Sale, error)
	Create(s *Sale) error
}
