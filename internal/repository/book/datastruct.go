package book

import (
	"database/sql"
)

type Book struct {
	ID            int64         `db:"id"`
	Name          string        `db:"name"`
	PublishedYear sql.NullInt32 `db:"published_year"`
	Description   string        `db:"description"`
	Price         int           `db:"price"`
	Image         []byte        `db:"image"`
}

type Author struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Surname string `db:"surname"`
}

func (b Book) TableName() string {
	return "book"
}
