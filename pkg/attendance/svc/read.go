package svc

import (
	"context"
	"toakbut/pkg/attendance/model"

	"github.com/uptrace/bun"
)

func (a *Attendance) Read(newData *model.Attendance) (*model.Attendance, error) {
	db := a.Db
	ctx := a.Ctx

	txFunc := func(context context.Context, tx bun.Tx) error {
		_, err := db.NewUpdate().
			Model(newData).
			WherePK().
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
