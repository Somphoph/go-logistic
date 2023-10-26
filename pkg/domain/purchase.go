package domain

import "time"

type Purchase struct {
	PurchaseId   int64     `json:"purchaseId"`
	PurchaseNo   string    `json:"purchaseNo"`
	PurchaseDate time.Time `json:"purchaseDate"`
	VendorId     int64     `json:"vendorId"`
	ProductId    int64     `json:"productId"`
	Qty          int       `json:"qty"`
	Cost         float32   `json:"cost"`
	ReqDate      time.Time `json:"reqDate"`
}

type Receive struct {
	ReceiveId    int64     `json:"receiveId"`
	ReceiveDocNo string    `json:"receiveDocNo"`
	ReceiveDate  time.Time `json:"receiveDate"`
	ProductId    int64     `json:"productId"`
	Qty          int       `json:"qty"`
	Cost         float32   `json:"cost"`
	PurchaseId   int64     `json:"purchaseId"`
}

type Vendor struct {
	VendorId       int64  `json:"vendorId"`
	VendorName     string `json:"vendorName"`
	VendorFullName string `json:"vendorFullName"`
	Addr           string `json:"addr"`
	Addr1          string `json:"addr1"`
}
