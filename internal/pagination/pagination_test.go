package pagination

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"gotest.tools/assert"
)

func TestBuildPagination(t *testing.T) {
	// Create Mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"total_count"}).AddRow(30))
	// Init
	pag := New(db)
	// Create Test Query
	var query sq.SelectBuilder
	query = squirrel.Select().From("TEST")
	res, _ := pag.BuildPagination(query, uint64(10), uint64(1))
	assert.Equal(t, res.PerPage, uint32(10))
	assert.Equal(t, res.LastPage, uint32(3))
}
