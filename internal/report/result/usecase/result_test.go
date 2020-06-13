package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
	mock_repository "github.com/abmid/icanvas-analytics/internal/report/result/repository/mock"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"gotest.tools/assert"
)

func RealSetup() *sql.DB {
	parse, err := pgx.ParseURI("postgres://abdulhamid:@localhost:5432/canvas_analytics_dev?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	db := stdlib.OpenDB(parse)
	return db
}

func TestCreateOrUpdateByCourseReportID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	mockRepo := mock_repository.NewMockReportResultRepository(ctrl)
	mockRepo.EXPECT().CreateOrUpdateByCourseReportID(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, m *entity.ReportResult) error {
		m.ID = 2
		return nil
	})
	result := entity.ReportResult{
		ReportCourseID: 1,
	}
	uc := NewReportResultUseCase(mockRepo)
	err := uc.CreateOrUpdateByCourseReportID(ctx, &result)
	assert.NilError(t, err)
	assert.Equal(t, uint32(2), result.ID)
}
