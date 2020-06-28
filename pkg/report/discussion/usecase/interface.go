/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"context"

	"github.com/abmid/icanvas-analytics/pkg/report/entity"
)

type ReportDiscussionUseCase interface {
	Create(ctx context.Context, reportDiss *entity.ReportDiscussion) error
	Read(ctx context.Context) ([]entity.ReportDiscussion, error)
	Update(ctx context.Context, reportDiss *entity.ReportDiscussion) error
	Delete(ctx context.Context, reportDissID uint32) error
	CreateOrUpdateByFilter(ctx context.Context, filter entity.ReportDiscussion, assigment *entity.ReportDiscussion) error
	FindFilter(ctx context.Context, filter entity.ReportDiscussion) ([]entity.ReportDiscussion, error)
}
