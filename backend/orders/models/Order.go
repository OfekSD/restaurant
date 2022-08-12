package models

import "time"

type Order struct{
	Id string `json:"id"`
	Dishes []string `json:"dishes"`
	Orderer string `json:"orderer"`
	OrderTime time.Time `json:"order_time`
	Delivered bool	`json:"delivered"`
}