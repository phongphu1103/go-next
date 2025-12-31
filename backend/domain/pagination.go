package domain

type Pagination struct {
	Page     int
	Limit    int 
	// Optional filters
	Title    *string
	Category *string
}