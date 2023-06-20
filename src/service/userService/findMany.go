package userService

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
)

func FindMany(filter model.UsersFilter, client *db.PrismaClient, ctx context.Context) ([]db.UserModel, error) {

	var limit int32 = 20

	parameter := []db.UserWhereParam{}

	currentUser := ctx.Value(srcModel.ConfigKey("currentUser")).(*db.UserModel)

	if filter.Filter != nil && filter.Filter.Limit != nil {
		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	if filter.ExceptMe != nil && *filter.ExceptMe {
		parameter = append(parameter,
			db.User.ID.Not(currentUser.ID),
		)
	}

	if filter.CompanyID != nil {
		parameter = append(parameter,
			db.User.CompanyID.Equals(*filter.CompanyID))
	}

	if filter.Role != nil {
		parameter = append(parameter,
			db.User.UserRole.Equals(db.Role(*filter.Role)),
		)
	}

	FetchUser := client.User.
		FindMany(
			parameter[:]...,
		).
		Take(int(limit)).OrderBy(
		db.User.CreatedAt.Order(db.DESC),
	)

	if filter.Filter != nil {
		if filter.Filter.After != nil {
			FetchUser = client.User.
				FindMany(
					parameter[:]...,
				).
				Take(int(limit)).
				Skip(1).
				Cursor(db.User.ID.Cursor(string(*filter.Filter.After))).
				OrderBy(
					db.User.CreatedAt.Order(db.DESC),
				)
		}

		if filter.Filter.Before != nil {
			FetchUser = client.User.
				FindMany(
					parameter[:]...,
				).
				Take(int(limit)).
				Skip(1).
				Cursor(db.User.ID.Cursor(string(*filter.Filter.Before))).
				OrderBy(
					db.User.CreatedAt.Order(db.ASC),
				)
		}
	}
	users, err := FetchUser.
		Exec(ctx)

	return users, err

}
