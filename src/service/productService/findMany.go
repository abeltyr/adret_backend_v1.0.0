package productService

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/utils"
	"context"
	"strings"
)

func FindMany(filter model.ProductsFilter, client *db.PrismaClient, ctx context.Context) ([]db.ProductModel, error) {

	var limit int32 = 20

	parameter := []db.ProductWhereParam{}
	cursorParameter := db.Product.CreatedAt.Order(db.DESC)

	if filter.Filter != nil && filter.Filter.Limit != nil {
		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	if filter.Category != nil {
		parameter = append(parameter,
			db.Product.CategoryProduct.Every(
				db.CategoryProduct.CategoryName.Equals(
					strings.ToLower(*filter.Category),
				),
			),
		)
	}

	if filter.CreatorID != nil {
		parameter = append(parameter,
			db.Product.CreatorID.Equals(*filter.CreatorID),
		)
	}

	if filter.Code != nil && filter.Title != nil {
		parameter = append(parameter,
			db.Product.Or(
				db.Product.ProductCode.Contains(
					*filter.Code,
				),
				db.Product.Title.Contains(
					*filter.Title,
				),
				db.Product.Title.Mode(db.QueryModeInsensitive),
				db.Product.ProductCode.Mode(db.QueryModeInsensitive),
			),
		)
	}

	if filter.Title != nil && filter.Code == nil {
		parameter = append(parameter,
			db.Product.Title.Contains(
				*filter.Title,
			),
		)
	}

	if filter.Code != nil && filter.Title == nil {
		parameter = append(parameter,
			db.Product.ProductCode.Contains(
				*filter.Code,
			),
		)
	}

	if filter.TopSelling != nil && *(filter.TopSelling) {
		parameter = append(parameter,
			db.Product.Inventory.Some(
				db.Inventory.SalesAmount.GTE(2),
			),
		)
		cursorParameter = db.Product.Inventory.Every(db.Inventory.SalesAmount.Order(db.ASC))
	}

	currentCompany := ctx.Value(srcModel.ConfigKey("currentCompany"))
	if currentCompany != nil {
		parameter = append(parameter,
			db.Product.CompanyID.Equals(currentCompany.(*db.CompanyModel).ID),
		)
	}

	FetchProduct := client.Product.
		FindMany(
			parameter[:]...,
		).
		Take(int(limit)).OrderBy(
		db.Product.CreatedAt.Order(db.DESC),
	)

	if filter.Filter != nil {
		if filter.Filter.After != nil {

			if filter.TopSelling == nil || (filter.TopSelling != nil && !(*filter.TopSelling)) {
				cursorParameter = db.Product.CreatedAt.Order(db.DESC)
			}
			FetchProduct = client.Product.
				FindMany(
					parameter[:]...,
				).
				Take(int(limit)).
				Skip(1).
				Cursor(db.Product.ID.Cursor(string(*filter.Filter.After))).
				OrderBy(
					cursorParameter,
				)
		}

		if filter.Filter.Before != nil {
			if filter.TopSelling == nil || (filter.TopSelling != nil && !(*filter.TopSelling)) {
				cursorParameter = db.Product.CreatedAt.Order(db.ASC)
			}
			FetchProduct = client.Product.
				FindMany(
					parameter[:]...,
				).
				Take(int(limit)).
				Skip(1).
				Cursor(db.Product.ID.Cursor(string(*filter.Filter.Before))).
				OrderBy(
					cursorParameter,
				)
		}
	}

	products, err := FetchProduct.
		Exec(ctx)

	return products, err

}
