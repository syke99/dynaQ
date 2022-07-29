package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/syke99/dynaQ/internal"
	"github.com/syke99/dynaQ/pkg/resources/models"
)

// DataBase provides a dynamic querier on an instance of a database
type DataBase struct {
	service
}

type service interface {
	Query(db *sql.DB, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error)
	QueryWithContext(db *sql.DB, ctx context.Context, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error)
}

// NewDbService creates a new DataBase with its associated service method(s)
func NewDbService() DataBase {
	return DataBase{}
}

// Query is a DataBase service method to execute a dynamic query on an instance of a database
func (db DataBase) Query(dBase *sql.DB, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error) {
	return internal.DatabaseQuery(dBase, query, timeFormat, queryParams)
}

// QueryWithContext is a DataBase service method to execute a dynamic query, with a context, on an instance of a database
func (db DataBase) QueryWithContext(dBase *sql.DB, ctx context.Context, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error) {
	if dBase == nil {
		var dummyResults []models.Row

		return dummyResults, errors.New("no database instance provided")
	}

	return internal.DatabaseQueryWithContext(dBase, ctx, query, timeFormat, queryParams)
}
