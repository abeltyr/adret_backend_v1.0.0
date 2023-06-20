package productVariationService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"strings"
)

func Create(data string, index int, product db.ProductModel, client *db.PrismaClient, ctx context.Context) (*db.ProductVariationModel, error) {
	title := strings.ToLower(data)

	company := ctx.Value(srcModel.ConfigKey("currentCompany")).(*db.CompanyModel)
	createdCategoryCompany, err := client.ProductVariation.
		UpsertOne(
			db.ProductVariation.TitleProductID(
				db.ProductVariation.Title.Equals(title),
				db.ProductVariation.ProductID.Equals(product.ID),
			),
		).
		Create(
			db.ProductVariation.Title.Set(title),
			db.ProductVariation.Product.Link(db.Product.ID.Equals(product.ID)),
			db.ProductVariation.Company.Link(db.Company.ID.Equals(company.ID)),
			db.ProductVariation.Order.Set(index),
		).
		Update(
			db.ProductVariation.Title.Set(title),
			db.ProductVariation.Product.Link(db.Product.ID.Equals(product.ID)),
			db.ProductVariation.Company.Link(db.Company.ID.Equals(company.ID)),
			db.ProductVariation.Order.Set(index),
		).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdCategoryCompany, nil

}
