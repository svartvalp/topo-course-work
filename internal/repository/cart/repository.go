package cart

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/samonzeweb/godb"
	"github.com/svartvalp/topo-course-work/internal/pkg/cart"
	"github.com/svartvalp/topo-course-work/internal/pkg/errors"
)

type Repository struct {
	db *godb.DB
}

func (r *Repository) DeleteCartBooks(ctx context.Context, id int64) error {
	_, err := r.db.
		DeleteFrom("cart_book").
		Where("cart_id = ?", id).Do()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetBookInCart(ctx context.Context, cartID int64, bookID int64) (cart.BookInCart, error) {
	var bookInCart BookInCart
	err := r.db.
		SelectFrom("cart_book").
		Where("cart_id = ? and book_id = ?", cartID, bookID).
		Do(&bookInCart)
	if err != nil {
		if err == sql.ErrNoRows {
			return cart.BookInCart{}, errors.NewStatusError(404, fmt.Sprintf("not found book in cart by cart id %v and book id %v", cartID, bookID))
		}
		return cart.BookInCart{}, err
	}
	return cart.BookInCart{
		CartID: bookInCart.CartID,
		BookID: bookInCart.BookID,
		Count:  bookInCart.Count,
	}, nil
}

func (r *Repository) DeleteBookFromCart(ctx context.Context, cartID int64, bookID int64) error {
	_, err := r.db.
		DeleteFrom("cart_book").
		Where("cart_id = ? and book_id = ?", cartID, bookID).Do()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateOrUpdateBookInCart(ctx context.Context, book cart.BookInCart) error {
	_, err := r.db.
		InsertInto("cart_book").
		Columns("cart_id", "book_id", "count").
		Values(book.CartID, book.BookID, book.Count).
		Suffix("on conflict on constraint cart_book_pkey do update set count = excluded.count").Do()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(ctx context.Context, c cart.Cart) (cart.Cart, error) {
	var repoCart Cart
	_, err := r.db.
		InsertInto("cart").
		Columns("session_id", "created_time").
		Values(c.SessionID, c.CreatedTime).
		Returning("*").
		DoWithReturning(&repoCart)
	if err != nil {
		return cart.Cart{}, err
	}
	return cart.Cart{
		ID:          repoCart.ID,
		SessionID:   repoCart.SessionID,
		CreatedTime: repoCart.CreatedTime,
	}, nil
}

func (r *Repository) GetBySession(ctx context.Context, session string) (cart.Cart, error) {
	var c Cart
	err := r.db.
		SelectFrom("cart").
		Where("session_id = ?", session).Do(&c)
	if err != nil {
		if err == sql.ErrNoRows {
			return cart.Cart{}, errors.NewStatusError(404, fmt.Sprintf("not found cart by session %v", session))
		}
		return cart.Cart{}, err
	}
	return cart.Cart{
		ID:          c.ID,
		SessionID:   c.SessionID,
		CreatedTime: c.CreatedTime,
	}, nil
}

func (r *Repository) GetCartBooks(ctx context.Context, id int64) ([]cart.BookWithCount, error) {
	var repoBooks []BookWithCount
	err := r.db.
		SelectFrom("cart_book cb").
		InnerJoin("book", "b", godb.Q("cb.book_id = b.id")).Where("cb.cart_id = ?", id).
		Columns("b.id as id", "cb.count as count", "b.name as name", "b.image as image", "b.price as price").
		Do(&repoBooks)
	if err != nil {
		return nil, err
	}
	books := make([]cart.BookWithCount, 0, len(repoBooks))
	for _, b := range repoBooks {
		books = append(books, cart.BookWithCount{
			ID:    b.ID,
			Count: b.Count,
			Name:  b.Name,
			Image: b.Image,
			Price: b.Price,
		})
	}
	return books, nil
}

func New(db *godb.DB) *Repository {
	return &Repository{db: db}
}
