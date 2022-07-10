package conn

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewConnectionService(t *testing.T) {
	// Act
	connService := NewConnectionService()

	// Assert
	assert.NotNil(t, connService)
}

func TestConnection_QueryRowWithContext(t *testing.T) {
	// Arrange
	testQuery := "SELECT * FROM testing"

	connService := NewConnectionService()

	rows := sqlmock.NewRows([]string{"id", "name", "date"}).
		AddRow(1, "test1", "2018-01-20 04:35")

	db, mock, _ := sqlmock.New()
	mock.ExpectQuery("SELECT (.*) FROM").
		WillReturnRows(rows)

	// Act
	con, err := db.Conn(context.Background())

	resRow, err := connService.QueryRowWithContext(con, context.Background(), testQuery, "2006-01-02 15:04")

	// Assert

	// Row 1
	assert.NoError(t, err)
	assert.Equal(t, "id", resRow[0].Name)
	assert.Equal(t, "1", resRow[0].Value)
	assert.Equal(t, "int64", resRow[0].Type)
	assert.Equal(t, "name", resRow[1].Name)
	assert.Equal(t, "test1", resRow[1].Value)
	assert.Equal(t, "string", resRow[1].Type)
	assert.Equal(t, "date", resRow[2].Name)
	// due to times having the chance of being off by a matter of milliseconds,
	// we wont worry about testing with the default time format. However,
	// this would be a good example of testing a date returned, as long as the
	// value passed in where time.Now() is passed is the same format
	// ass the time format you're using
	// assert.Equal(t, fmt.Sprintf("%v", time.Now()), resRows[0][2].Value)
	assert.Equal(t, "time.Time", resRow[2].Type)
}

func TestConnection_QueryWithContext(t *testing.T) {
	// Arrange
	testQuery := "SELECT * FROM testing"

	connService := NewConnectionService()

	rows := sqlmock.NewRows([]string{"id", "name", "date"}).
		AddRow(1, "test1", "2018-01-20 04:35").
		AddRow("2", "test2", "2018-01-20 04:35").
		AddRow(1.1, "test3", "2018-01-20 04:35")

	db, mock, _ := sqlmock.New()
	mock.ExpectQuery("SELECT (.*) FROM").
		WillReturnRows(rows)

	// Act
	con, err := db.Conn(context.Background())

	resRows, err := connService.QueryWithContext(con, context.Background(), testQuery, "2006-01-02 15:04")

	// Assert

	// Row 1
	assert.NoError(t, err)
	assert.Equal(t, "id", resRows[0][0].Name)
	assert.Equal(t, "1", resRows[0][0].Value)
	assert.Equal(t, "int64", resRows[0][0].Type)
	assert.Equal(t, "name", resRows[0][1].Name)
	assert.Equal(t, "test1", resRows[0][1].Value)
	assert.Equal(t, "string", resRows[0][1].Type)
	assert.Equal(t, "date", resRows[0][2].Name)
	// due to times having the chance of being off by a matter of milliseconds,
	// we wont worry about testing with the default time format. However,
	// this would be a good example of testing a date returned, as long as the
	// value passed in where time.Now() is passed is the same format
	// ass the time format you're using
	// assert.Equal(t, fmt.Sprintf("%v", time.Now()), resRows[0][2].Value)
	assert.Equal(t, "time.Time", resRows[0][2].Type)

	// Row 2
	assert.NoError(t, err)
	assert.Equal(t, "id", resRows[1][0].Name)
	assert.Equal(t, "2", resRows[1][0].Value)
	assert.Equal(t, "string", resRows[1][0].Type)
	assert.Equal(t, "name", resRows[1][1].Name)
	assert.Equal(t, "test2", resRows[1][1].Value)
	assert.Equal(t, "string", resRows[1][1].Type)
	assert.Equal(t, "date", resRows[1][2].Name)
	// due to times having the chance of being off by a matter of milliseconds,
	// we wont worry about testing with the default time format. However,
	// this would be a good example of testing a date returned, as long as the
	// value passed in where time.Now() is passed is the same format
	// ass the time format you're using
	// assert.Equal(t, fmt.Sprintf("%v", time.Now()), resRows[0][2].Value)
	assert.Equal(t, "time.Time", resRows[1][2].Type)

	// Row 3
	assert.NoError(t, err)
	assert.Equal(t, "id", resRows[2][0].Name)
	assert.Equal(t, "1.1", resRows[2][0].Value)
	assert.Equal(t, "float64", resRows[2][0].Type)
	assert.Equal(t, "name", resRows[2][1].Name)
	assert.Equal(t, "test3", resRows[2][1].Value)
	assert.Equal(t, "string", resRows[2][1].Type)
	assert.Equal(t, "date", resRows[2][2].Name)
	// due to times having the chance of being off by a matter of milliseconds,
	// we wont worry about testing with the default time format. However,
	// this would be a good example of testing a date returned, as long as the
	// value passed in where time.Now() is passed is the same format
	// ass the time format you're using
	// assert.Equal(t, fmt.Sprintf("%v", time.Now()), resRows[0][2].Value)
	assert.Equal(t, "time.Time", resRows[2][2].Type)
}
