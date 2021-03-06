package conn

import (
	"context"
	"database/sql"
	"errors"
	"github.com/syke99/dynaQ/internal"
	"github.com/syke99/dynaQ/pkg/resources/models"
)

// Connection provides a dynamic querier on a single database connection
type Connection struct {
	service
}

type service interface {
	QueryWithDefaultContext(conn *sql.Conn, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error)
	QueryWithContext(conn *sql.Conn, ctx context.Context, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error)
}

// NewConnectionService creates a new Connection with its associated service method(s)
func NewConnectionService() Connection {
	return Connection{}
}

// QueryWithDefaultContext is a Connection service method to execute a dynamic query, with the default context of context.Background(), on a single database connection
func (c Connection) QueryWithDefaultContext(conn *sql.Conn, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error) {
	if conn == nil {
		var dummyResults []models.Row

		return dummyResults, errors.New("no database connection provided")
	}

	return internal.ConnectionQueryWithContext(conn, context.Background(), query, timeFormat, queryParams)
}

// QueryWithContext is a Connection service method to execute a dynamic query, with a context, on a single database connection
func (c Connection) QueryWithContext(conn *sql.Conn, ctx context.Context, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error) {
	if conn == nil {
		var dummyResults []models.Row

		return dummyResults, errors.New("no database connection provided")
	}

	return internal.ConnectionQueryWithContext(conn, ctx, query, timeFormat, queryParams)
}
