// models/ticket.go
package models

type Ticket struct {
	ID    int    `json:"id"`
	Movie string `json:"movie"`
	Price int    `json:"price"`
}
