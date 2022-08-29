package domain

import "time"

type Purchase struct {
	PurchaseId   int64     `json:"purchase-id"`
	PurchaseNo   string    `json:"purchase-no"`
	PurchaseDate time.Time `json:"purchase-date"`
	VendorId     int64     `json:"vendor-id"`
	ProductId    int64     `json:"product-id"`
	Qty          int       `json:"qty"`
	Cost         float32   `json:"cost"`
	ReqDate      time.Time `json:"req-date"`
}

type Receive struct {
	ReceiveId    int64     `json:"receive-id"`
	ReceiveDocNo string    `json:"receive-doc-no"`
	ReceiveDate  time.Time `json:"receive-date"`
	ProductId    int64     `json:"product-id"`
	Qty          int       `json:"qty"`
	Cost         float32   `json:"cost"`
	PurchaseId   int64     `json:"purchase-id"`
}

type Vendor struct {
	VendorId       int64  `json:"vendor-id"`
	VendorName     string `json:"vendor-name"`
	VendorFullName string `json:"vendor-full-name"`
	Addr           string `json:"addr"`
	Addr1          string `json:"addr-1"`
}
