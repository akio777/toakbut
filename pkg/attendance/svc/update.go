package svc

import (
	"context"
	"database/sql"
	"errors"
	"toakbut/pkg/attendance/model"

	"github.com/uptrace/bun"
)

func (a *Attendance) Update(newData *model.Attendance) (*model.Attendance, error) {
	db := a.Db
	ctx := a.Ctx

	latest := model.Attendance{}
	err := db.NewSelect().
		Model(&latest).
		Where("user_id = ?", newData.UserID).
		Order("check_in DESC").
		Limit(1).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("attendance `in` not exists in today, please contact participant")
		} else {
			return nil, err
		}
	}
	if latest.CheckOut == nil {
		latest.CheckOut = newData.CheckOut
		if latest.CheckOut.Before(*latest.CheckIn) {
			return nil, errors.New("attendance `out` should not less than `in`")
		}
		txFunc := func(context context.Context, tx bun.Tx) error {
			_, err := db.NewUpdate().
				Model(&latest).
				Where("id = ?", latest.ID).
				Where("user_id = ?", newData.UserID).
				Exec(ctx)
			if err != nil {
				return err
			}
			return nil
		}
		if err := db.RunInTx(ctx, nil, txFunc); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("attendance `out` already submitted")

	}

	return newData, nil
}
