package fileRelationService

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
)

func Delete(uploadedDataRelation model.UploadedDataRelation, client *db.PrismaClient, ctx context.Context) (*db.FileRelationModel, error) {

	fileRelation, err := client.FileRelation.
		FindUnique(
			db.FileRelation.TableOrderValue(
				db.FileRelation.Table.Equals(uploadedDataRelation.Table),
				db.FileRelation.Order.Equals(uploadedDataRelation.Order),
				db.FileRelation.Value.Equals(uploadedDataRelation.Value),
			),
		).
		Delete().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return fileRelation, nil

}
