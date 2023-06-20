package categoryProductService

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
	"strings"
)

func Create(categoryName string, client *db.PrismaClient, ctx context.Context) (*db.CategoryProductModel, error) {

	product := ctx.Value(srcModel.ConfigKey("product")).(*db.ProductModel)
	category := strings.ToLower(categoryName)

	createdCategoryProduct, err := client.CategoryProduct.
		UpsertOne(
			db.CategoryProduct.CategoryNameProductID(
				db.CategoryProduct.CategoryName.Equals(category),
				db.CategoryProduct.ProductID.Equals(product.ID),
			),
		).
		Create(
			db.CategoryProduct.Category.Link(db.Category.Name.Equals(category)),
			db.CategoryProduct.Product.Link(db.Product.ID.Equals(product.ID)),
		).
		Update().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdCategoryProduct, nil

}
