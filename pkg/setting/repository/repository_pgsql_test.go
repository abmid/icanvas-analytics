package repository

import (
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abmid/icanvas-analytics/pkg/setting/entity"
	"gotest.tools/assert"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	mock.ExpectQuery("INSERT INTO " + DBNAME).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("2"))

	repo := NewRepositoryPG(db)
	setting := entity.Setting{
		Name: "Test",
	}
	err = repo.Create(&setting)
	assert.NilError(t, err)
	assert.Equal(t, setting.ID, uint32(2))
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	mock.ExpectExec("UPDATE settings").WillReturnResult(sqlmock.NewResult(2, 1))
	repo := NewRepositoryPG(db)
	err = repo.Update(2, entity.Setting{})
	assert.NilError(t, err)
}

func TestFindByFilter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	t.Run("exist", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "name", "category", "value", "created_at", "updated_at",
		}).AddRow(1, "name", "category", "value", time.Now(), time.Now())

		mock.ExpectQuery("SELECT").WithArgs("Test").WillReturnRows(rows)

		repo := NewRepositoryPG(db)
		filter := entity.Setting{
			Category: "Test",
		}
		res, err := repo.FindByFilter(filter)
		assert.NilError(t, err)
		assert.Equal(t, len(res), 1)
	})

	t.Run("not-exists", func(t *testing.T) {

		mock.ExpectQuery("SELECT").WithArgs("Test").WillReturnError(sql.ErrNoRows)

		repo := NewRepositoryPG(db)
		filter := entity.Setting{
			Category: "Test",
		}
		res, err := repo.FindByFilter(filter)
		assert.NilError(t, err)
		assert.Equal(t, len(res), 0)
	})

}
