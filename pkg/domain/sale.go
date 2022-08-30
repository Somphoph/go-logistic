package domain

import (
	"time"
)

type Sale struct {
	SaleId     int64     `json:"sale-id"`
	SaleNo     string    `json:"sale-no"`
	SaleDate   time.Time `json:"sale-date"`
	CustomerId int64     `json:"customer-id"`
	ProductId  int64     `json:"product-id"`
	Qty        int       `json:"qty"`
	Price      float32   `json:"price"`
	ReqDate    time.Time `json:"req-date"`
}
type Invoice struct {
	InvId      int64     `json:"inv-id"`
	InvNo      string    `json:"inv-no"`
	InvDate    time.Time `json:"inv-date"`
	CustomerId string    `json:"customer-id"`
	ProductId  string    `json:"product-id"`
	Qty        int       `json:"qty"`
	Price      float32   `json:"price"`
	ReceiveId  int64     `json:"receive-id"`
	SaleId     int64     `json:"sale-id"`
}

type Customer struct {
	CustomerId  int64  `json:"customer-id"`
	CusName     string `json:"cus-name"`
	CusFullName string `json:"cus-full-name"`
	Addr        string `json:"addr"`
	Addr1       string `json:"addr-1"`
}

type Product struct {
	ProductId       int64   `json:"product-id"`
	ProductName     string  `json:"product-name"`
	ProductFullName string  `json:"product-full-name"`
	Cost            float32 `json:"cost"`
	Price           float32 `json:"price"`
}

type SaleSvc interface {
	Get(saleNo string) (*Sale, error)
	List() ([]*Sale, error)
	Create(s *Sale) error
}

type SaleDB interface {
	Get(id int64) (*Sale, error)
	List(category string) ([]*Sale, error)
	Create(p *Sale) error
	Delete(id int64) error
}
