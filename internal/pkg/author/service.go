package author

import (
	"context"
)

type Service struct {
	authorRepo authorAdapter
}

type authorAdapter interface {
	GetAuthor(ctx context.Context, id int64) (Author, error)
}

func New(authorRepo authorAdapter) *Service {
	return &Service{authorRepo: authorRepo}
}

func (s *Service) GetAuthor(ctx context.Context, id int64) (Author, error) {
	return s.authorRepo.GetAuthor(ctx, id)
}
