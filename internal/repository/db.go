package repository

import (
	"context"

	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/postgresql"
	"github.com/svartvalp/topo-course-work/internal/config"
)

func NewDB(ctx context.Context) (*godb.DB, error) {
	db, err := godb.Open(postgresql.Adapter, config.GetConfig().DataSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}
