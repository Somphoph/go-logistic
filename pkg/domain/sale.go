package domain

import (
	"time"
)

type Sale struct {
	SaleId     int64     `json:"saleId"`
	SaleNo     string    `json:"saleNo"`
	SaleDate   time.Time `json:"saleDate"`
	CustomerId int64     `json:"customerId"`
	ProductId  int64     `json:"productId"`
	Qty        int       `json:"qty"`
	Price      float32   `json:"price"`
	ReqDate    time.Time `json:"reqDate"`
}
type Invoice struct {
	InvId      int64     `json:"invId"`
	InvNo      string    `json:"invNo"`
	InvDate    time.Time `json:"invDate"`
	CustomerId string    `json:"customerId"`
	ProductId  string    `json:"productId"`
	Qty        int       `json:"qty"`
	Price      float32   `json:"price"`
	ReceiveId  int64     `json:"receiveId"`
	SaleId     int64     `json:"saleId"`
}

type Customer struct {
	CustomerId  int64  `json:"customerId"`
	CusName     string `json:"cusName"`
	CusFullName string `json:"cusFullName"`
	Addr        string `json:"addr"`
	Addr1       string `json:"addr1"`
}

type Product struct {
	ProductId       int64   `json:"productId"`
	ProductName     string  `json:"productName"`
	ProductFullName string  `json:"productFullName"`
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
