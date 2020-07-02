package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abmid/icanvas-analytics/pkg/user/entity"
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

func TestFind(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	t.Run("exist", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "deleted_at"}).AddRow(
			1,
			"test",
			"test@test.com",
			"pass",
			time.Now(),
			time.Now(),
		))

		repo := NewPG(db)
		res, err := repo.Find("test@test.com")
		assert.NilError(t, err)
		assert.Equal(t, res.Email, "test@test.com")

	})
	t.Run("not-exist", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "deleted_at"}))

		repo := NewPG(db)
		res, err := repo.Find("test@test.com")
		assert.NilError(t, err)
		assert.Equal(t, res == nil, true)
	})

}

func TestAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	t.Run("user-exist", func(t *testing.T) {
		user := entity.User{
			ID: 1,
		}
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "deleted_at"}).AddRow(
			user.ID,
			user.Name,
			user.Email,
			user.Password,
			user.CreatedAt,
			user.DeletedAt,
		))
		repo := NewPG(db)
		res, err := repo.All()
		assert.NilError(t, err)
		assert.Equal(t, len(res), 1)
	})
	t.Run("user-not-exist", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "deleted_at"}))
		repo := NewPG(db)
		res, err := repo.All()
		assert.NilError(t, err)
		assert.Equal(t, len(res), 0)
	})
	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnError(errors.New("errors"))
		repo := NewPG(db)
		res, err := repo.All()
		assert.ErrorContains(t, err, "")
		assert.Equal(t, len(res), 0)
	})
}
