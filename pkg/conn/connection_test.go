package conn

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConnectionService(t *testing.T) {
	// Arrange
	var conn *sql.Conn

	// Act
	connService, err := NewConnectionService(conn)

	// Assert
	assert.NotNil(t, connService)
	assert.NoError(t, err)
}

func TestNewConnectionService_NoConnection(t *testing.T) {
	// Act
	connService, err := NewConnectionService(nil)

	// Assert
	assert.NotNil(t, connService)
	assert.NoError(t, err)
}

func TestConnection_QueryWithContext(t *testing.T) {
	// Arrange
	testQuery := "SELECT * FROM testing"

	var conn *sql.Conn
	connService, _ := NewConnectionService(conn)

	rows := sqlmock.NewRows([]string{"id", "name", "date"}).
		AddRow(1, "test1", "10/28/90").
		AddRow("2", "test2", "10/28/90").
		AddRow(1.1, "test3", "10/28/90")

	db, mock, _ := sqlmock.New()
	mock.ExpectQuery("SELECT (.*) FROM").
		WillReturnRows(rows)

	// Act
	con, err := db.Conn(context.Background())

	resRows, err := connService.QueryWithContext(con, context.Background(), testQuery, "10/28/90")

	// Assert

	// Row 1
	assert.NoError(t, err)
	assert.Equal(t, "id", resRows[0][0].Column)
	assert.Equal(t, "1", resRows[0][0].Value)
	assert.Equal(t, "int64", resRows[0][0].Type)
	assert.Equal(t, "name", resRows[0][1].Column)
	assert.Equal(t, "test1", resRows[0][1].Value)
	assert.Equal(t, "string", resRows[0][1].Type)
	assert.Equal(t, "date", resRows[0][2].Column)
	// due to times having the chance of being off by a matter of milliseconds,
	// we wont worry about testing with the default time format. However,
	// this would be a good example of testing a date returned, as long as the
	// value passed in where time.Now() is passed is the same format
	// ass the time format you're using
	// assert.Equal(t, fmt.Sprintf("%v", time.Now()), resRows[0][2].Value)
	assert.Equal(t, "time.Time", resRows[0][2].Type)

	//// Row 2
	//assert.NoError(t, err)
	//assert.Equal(t, "id", resRows[1][0].Column)
	//assert.Equal(t, "2", resRows[1][0].Value)
	//assert.Equal(t, "string", resRows[1][0].Type)
	//assert.Equal(t, "name", resRows[1][1].Column)
	//assert.Equal(t, "test2", resRows[1][1].Value)
	//assert.Equal(t, "string", resRows[1][1].Type)
	//assert.Equal(t, "date", resRows[1][2].Column)
	//// due to times having the chance of being off by a matter of milliseconds,
	//// we wont worry about testing with the default time format. However,
	//// this would be a good example of testing a date returned, as long as the
	//// value passed in where time.Now() is passed is the same format
	//// ass the time format you're using
	//// assert.Equal(t, fmt.Sprintf("%v", time.Now()), resRows[0][2].Value)
	//assert.Equal(t, "time.Time", resRows[1][2].Type)

	// Row 3
	assert.NoError(t, err)
	assert.Equal(t, "id", resRows[2][0].Column)
	assert.Equal(t, "1.1", resRows[2][0].Value)
	assert.Equal(t, "float64", resRows[2][0].Type)
	assert.Equal(t, "name", resRows[2][1].Column)
	assert.Equal(t, "test3", resRows[2][1].Value)
	assert.Equal(t, "string", resRows[2][1].Type)
	assert.Equal(t, "date", resRows[2][2].Column)
	// due to times having the chance of being off by a matter of milliseconds,
	// we wont worry about testing with the default time format. However,
	// this would be a good example of testing a date returned, as long as the
	// value passed in where time.Now() is passed is the same format
	// ass the time format you're using
	// assert.Equal(t, fmt.Sprintf("%v", time.Now()), resRows[0][2].Value)
	assert.Equal(t, "time.Time", resRows[2][2].Type)
}
