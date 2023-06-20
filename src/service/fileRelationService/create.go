package fileRelationService

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
)

func Create(uploadedDataRelation model.UploadedDataRelation, client *db.PrismaClient, ctx context.Context) (*db.FileRelationModel, error) {

	createdCategoryCompany, err := client.FileRelation.
		UpsertOne(
			db.FileRelation.TableOrderValue(
				db.FileRelation.Table.Equals(uploadedDataRelation.Table),
				db.FileRelation.Order.Equals(uploadedDataRelation.Order),
				db.FileRelation.Value.Equals(uploadedDataRelation.Value),
			),
		).
		Create(
			db.FileRelation.File.Link(db.File.ID.Equals(uploadedDataRelation.File)),
			db.FileRelation.Table.Set(uploadedDataRelation.Table),
			db.FileRelation.Value.Set(uploadedDataRelation.Value),
			db.FileRelation.Order.Set(uploadedDataRelation.Order),
		).
		Update(
			db.FileRelation.File.Link(db.File.ID.Equals(uploadedDataRelation.File)),
		).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdCategoryCompany, nil

}
