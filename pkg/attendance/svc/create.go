package svc

import (
	"context"
	"database/sql"
	"errors"

	"toakbut/pkg/attendance/model"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

type Attendance struct {
	Ctx context.Context
	Db  *bun.DB
	Log *logrus.Logger
}

func (a *Attendance) Create(newData *model.Attendance) (*model.Attendance, error) {
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
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	if latest.CheckIn != nil && newData.CheckIn.Day() == latest.CheckIn.Day() {
		return nil, errors.New("attendance `in` record for the same day already exists")
	}
	txFunc := func(context context.Context, tx bun.Tx) error {
		_, err := db.NewInsert().
			Model(newData).
			Returning("*").
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	if err := db.RunInTx(ctx, nil, txFunc); err != nil {
		return nil, err
	}
	return newData, nil
}
