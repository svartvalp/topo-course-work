package book

import (
	"context"

	"github.com/svartvalp/topo-course-work/internal/pkg/cart"
)

type Service struct {
	bookRepo    bookAdapter
	cartService *cart.Service
}

type bookAdapter interface {
	ListBooks(ctx context.Context, query ListBooksQuery) ([]Book, error)
	GetBookAuthors(ctx context.Context, id int64) ([]Author, error)
}

func (s *Service) ListBooks(ctx context.Context, req ListBooksRequest) ([]Enriched, error) {
	query := ListBooksQuery{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
		Filter: req.Filter,
	}

	inCartMap := make(map[int64]int, 0)
	if req.Filter != nil && req.Filter.Session != "" {
		cartInfo, err := s.cartService.GetCart(ctx, req.Filter.Session)
		if err != nil {
			return nil, err
		}
		for _, book := range cartInfo.Books {
			inCartMap[book.ID] = book.Count
		}
	}

	books, err := s.bookRepo.ListBooks(ctx, query)
	enriched := make([]Enriched, 0, len(books))
	for _, b := range books {
		authors, err := s.bookRepo.GetBookAuthors(ctx, b.ID)
		if err != nil {
			return nil, err
		}
		for i := range authors {
			fullName := authors[i].Surname
			if authors[i].Name != "" {
				fullName += " " + authors[i].Name
			}
			authors[i].FullName = fullName
		}
		count := 0
		if val, ok := inCartMap[b.ID]; ok {
			count = val
		}
		enriched = append(enriched, Enriched{
			Book:        b,
			InCartCount: count,
			Authors:     authors,
		})
	}
	if err != nil {
		return nil, err
	}
	return enriched, nil
}

func New(bookRepo bookAdapter, cartService *cart.Service) *Service {
	return &Service{bookRepo: bookRepo, cartService: cartService}
}
