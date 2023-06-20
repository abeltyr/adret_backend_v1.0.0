package fileRelationService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func FindMany(id string, client *db.PrismaClient, ctx context.Context) ([]db.FileRelationModel, error) {

	summary, err := client.FileRelation.
		FindMany(
			db.FileRelation.Table.Equals(
				id,
			),
		).
		With(db.FileRelation.File.Fetch()).
		Exec(ctx)

	return summary, err
}
