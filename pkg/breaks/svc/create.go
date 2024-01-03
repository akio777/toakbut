package svc

import (
	"context"
	"database/sql"
	"errors"

	"toakbut/pkg/breaks/model"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

type Breaks struct {
	Ctx context.Context
	Db  *bun.DB
	Log *logrus.Logger
}

func (a *Breaks) Create(newData *model.Breaks) (*model.Breaks, error) {
	db := a.Db
	ctx := a.Ctx
	latest := model.Breaks{}
	err := db.NewSelect().
		Model(&latest).
		Where("user_id = ?", newData.UserID).
		Where("break_in IS NOT NULL").
		Where("break_out IS NULL").
		Order("break_in DESC").
		Limit(1).
		Scan(ctx)

	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	if latest.BreakIn == nil {
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
	} else {
		return nil, errors.New("you're in break")
	}
	return newData, nil
}
