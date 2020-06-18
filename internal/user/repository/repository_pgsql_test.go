package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abmid/icanvas-analytics/internal/user/entity"
	"gotest.tools/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	repo := NewPG(db)
	user := entity.User{
		Name:     "Test",
		Email:    "test@test.com",
		Password: "pass",
	}
	err = repo.Create(&user)

	assert.NilError(t, err)
	assert.Equal(t, user.ID, uint32(1))
}
