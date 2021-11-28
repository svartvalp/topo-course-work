package cart

import (
	"time"
)

type Cart struct {
	ID          int64     `db:"id"`
	SessionID   string    `db:"session_id"`
	CreatedTime time.Time `db:"created_time"`
}

type BookWithCount struct {
	ID    int64  `db:"id"`
	Count int    `db:"count"`
	Name  string `db:"name"`
	Image []byte `db:"image"`
	Price int    `db:"price"`
}

type BookInCart struct {
	CartID int64 `db:"cart_id"`
	BookID int64 `db:"book_id"`
	Count  int   `db:"count"`
}

func (b BookInCart) TableName() string {
	return "cart_book"
}
