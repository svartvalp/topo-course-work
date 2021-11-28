package author

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/samonzeweb/godb"
	"github.com/svartvalp/topo-course-work/internal/pkg/author"
	"github.com/svartvalp/topo-course-work/internal/pkg/errors"
)

type Repository struct {
	db *godb.DB
}

func (r *Repository) GetAuthor(ctx context.Context, id int64) (author.Author, error) {
	var auth Author
	err := r.db.SelectFrom("author").Where("id = ?", id).Do(&auth)
	if err != nil {
		if err == sql.ErrNoRows {
			return author.Author{}, errors.NewStatusError(400, fmt.Sprintf("not found author by id %v", id))
		}
		return author.Author{}, err
	}
	return author.Author{
		ID:          auth.ID,
		Name:        auth.Name,
		Surname:     auth.Surname,
		MiddleName:  auth.MiddleName,
		BirthYear:   auth.BirthYear,
		Description: auth.Description,
	}, nil
}

func New(db *godb.DB) *Repository {
	return &Repository{db: db}
}
