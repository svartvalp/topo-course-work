package cart

import (
	"time"
)

type Info struct {
	SessionID string
	Books     []BookWithCount
}

type Cart struct {
	ID          int64
	SessionID   string
	CreatedTime time.Time
}

type BookWithCount struct {
	ID    int64
	Count int
	Name  string
	Price int
	Image []byte
}

type BookInCart struct {
	CartID int64
	BookID int64
	Count  int
}
