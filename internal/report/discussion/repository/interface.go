package repository

import (
	"context"

	"github.com/abmid/icanvas-analytics/internal/report/entity"
)

type DisscussionRepositoryPG interface {
	Create(ctx context.Context, reportDiss *entity.ReportDiscussion) error
	Read(ctx context.Context) ([]entity.ReportDiscussion, error)
	Update(ctx context.Context, reportDiss *entity.ReportDiscussion) error
	Delete(ctx context.Context, reportDissID uint32) error
	FindFilter(ctx context.Context, filter entity.ReportDiscussion) ([]entity.ReportDiscussion, error)
	FindFirstByFilter(ctx context.Context, filter entity.ReportDiscussion) (entity.ReportDiscussion, error)
}
