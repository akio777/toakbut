package svc

import (
	// "context"
	"context"
	"database/sql"
	"errors"
	"toakbut/pkg/breaks/model"

	"github.com/uptrace/bun"
	// "github.com/uptrace/bun"
)

func (a *Breaks) Update(newData *model.Breaks) (*model.Breaks, error) {
	db := a.Db
	ctx := a.Ctx

	latest := model.Breaks{}
	err := db.NewSelect().
		Model(&latest).
		Where("user_id = ?", newData.UserID).
		Where("break_in IS NOT NULL").
		Order("break_in DESC").
		Limit(1).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("`break` not exists in today, please contact participant")
		} else {
			return nil, err
		}
	}
	if latest.BreakOut == nil {
		latest.BreakOut = newData.BreakOut
		if latest.BreakOut.Before(*latest.BreakOut) {
			return nil, errors.New("break `back` should not less than`break`")
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
		return nil, errors.New("break `back` already submitted")

	}

	return newData, nil
}
