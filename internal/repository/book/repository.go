package book

import (
	"context"

	"github.com/samonzeweb/godb"
	"github.com/svartvalp/topo-course-work/internal/pkg/book"
)

type Repository struct {
	db *godb.DB
}

func (r *Repository) GetBookAuthors(ctx context.Context, id int64) ([]book.Author, error) {
	var repoAuthors []Author
	err := r.db.
		SelectFrom("book_author ba").
		InnerJoin("author", "au", godb.Q("ba.author_id = au.id")).
		Where("ba.book_id = ?", id).Do(&repoAuthors)
	if err != nil {
		return nil, err
	}
	authors := make([]book.Author, 0, len(repoAuthors))
	for _, auth := range repoAuthors {
		authors = append(authors, book.Author{
			ID:      auth.ID,
			Name:    auth.Name,
			Surname: auth.Surname,
		})
	}
	return authors, nil
}

func (r *Repository) ListBooks(ctx context.Context, query book.ListBooksQuery) ([]book.Book, error) {
	var repoBooks []Book
	sel := r.db.
		Select(&repoBooks).
		Limit(query.Limit).
		Offset(query.Offset)
	if query.Filter != nil {
		filter := query.Filter
		if filter.Name != "" {
			sel.Where("name = ?", filter.Name)
		}
	}
	err := sel.Do()
	if err != nil {
		return nil, err
	}
	books := make([]book.Book, 0, len(repoBooks))
	for _, b := range repoBooks {
		books = append(books, book.Book{
			ID:            b.ID,
			Name:          b.Name,
			PublishedYear: int(b.PublishedYear.Int32),
			Description:   b.Description,
			Price:         b.Price,
			Image:         b.Image,
		})
	}
	return books, nil
}

func NewRepository(db *godb.DB) *Repository {
	return &Repository{db: db}
}
