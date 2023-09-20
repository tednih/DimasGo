// models/ticket.go
package models

type Ticket struct {
	ID      int    `json:"id"`
	Match   string `json:"match"`
	Price   int    `json:"price"`
	Tickets int    `json:"tickets"`
}
