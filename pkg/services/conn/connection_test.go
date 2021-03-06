package conn

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/syke99/dynaQ/internal"
)

func TestNewConnectionService(t *testing.T) {
	// Act
	connService := NewConnectionService()

	// Assert
	assert.NotNil(t, connService)
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

	dummyQueryArgs := internal.QueryArgs{}

	resRows, err := connService.QueryWithContext(con, context.Background(), testQuery, "2006-01-02 15:04", dummyQueryArgs)

	// Assert

	// Row 1
	assert.NoError(t, err)
	assert.Equal(t, "id", resRows[0].Columns[0].
		Name)
	assert.Equal(t, "1", resRows[0].Columns[0].
		Value)
	assert.Equal(t, "int64", resRows[0].Columns[0].
		Type)
	assert.Equal(t, "name", resRows[0].Columns[1].Name)
	assert.Equal(t, "test1", resRows[0].Columns[1].Value)
	assert.Equal(t, "string", resRows[0].Columns[1].Type)
	assert.Equal(t, "date", resRows[0].Columns[2].Name)
	// due to times having the chance of being off by a matter of milliseconds,
	// we wont worry about testing with the default time format. However,
	// this would be a good example of testing a date returned, as long as the
	// value passed in where time.Now() is passed is the same format
	// ass the time format you're using
	// assert.Equal(t, fmt.Sprintf("%v", time.Now()), resRows[0].Columns[2].Value)
	assert.Equal(t, "time.Time", resRows[0].Columns[2].Type)

	// Row 2
	assert.NoError(t, err)
	assert.Equal(t, "id", resRows[1].Columns[0].Name)
	assert.Equal(t, "2", resRows[1].Columns[0].Value)
	assert.Equal(t, "string", resRows[1].Columns[0].Type)
	assert.Equal(t, "name", resRows[1].Columns[1].Name)
	assert.Equal(t, "test2", resRows[1].Columns[1].Value)
	assert.Equal(t, "string", resRows[1].Columns[1].Type)
	assert.Equal(t, "date", resRows[1].Columns[2].Name)
	// due to times having the chance of being off by a matter of milliseconds,
	// we wont worry about testing with the default time format. However,
	// this would be a good example of testing a date returned, as long as the
	// value passed in where time.Now() is passed is the same format
	// ass the time format you're using
	// assert.Equal(t, fmt.Sprintf("%v", time.Now()), resRows[0].Columns[2].Value)
	assert.Equal(t, "time.Time", resRows[1].Columns[2].Type)

	// Row 3
	assert.NoError(t, err)
	assert.Equal(t, "id", resRows[2].Columns[0].Name)
	assert.Equal(t, "1.1", resRows[2].Columns[0].Value)
	assert.Equal(t, "float64", resRows[2].Columns[0].Type)
	assert.Equal(t, "name", resRows[2].Columns[1].Name)
	assert.Equal(t, "test3", resRows[2].Columns[1].Value)
	assert.Equal(t, "string", resRows[2].Columns[1].Type)
	assert.Equal(t, "date", resRows[2].Columns[2].Name)
	// due to times having the chance of being off by a matter of milliseconds,
	// we wont worry about testing with the default time format. However,
	// this would be a good example of testing a date returned, as long as the
	// value passed in where time.Now() is passed is the same format
	// ass the time format you're using
	// assert.Equal(t, fmt.Sprintf("%v", time.Now()), resRows[0].Columns[2].Value)
	assert.Equal(t, "time.Time", resRows[2].Columns[2].Type)
}

func TestConnection_QueryWithContext_NoContext(t *testing.T) {
	// Arrange
	testQuery := "SELECT * FROM testing"

	connService := NewConnectionService()

	rows := sqlmock.NewRows([]string{"id", "name", "date"}).
		AddRow(1, "test1", "2018-01-20 04:35").
		AddRow("2", "test2", "2018-01-20 04:35").
		AddRow(1.1, "test3", "2018-01-20 04:35")

	_, mock, _ := sqlmock.New()
	mock.ExpectQuery("SELECT (.*) FROM").
		WillReturnRows(rows)

	// Act
	dummyQueryArgs := internal.QueryArgs{}

	_, err := connService.QueryWithContext(nil, context.Background(), testQuery, "2006-01-02 15:04", dummyQueryArgs)

	// Assert
	assert.Error(t, err)
}

func TestConnection_QueryWithContext_BadQuery(t *testing.T) {
	// Arrange

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

	dummyQueryArgs := internal.QueryArgs{}

	_, err = connService.QueryWithContext(con, context.Background(), "", "2006-01-02 15:04", dummyQueryArgs)

	// Assert
	assert.Error(t, err)
}
