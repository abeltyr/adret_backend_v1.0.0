package fileService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Find(id string, client *db.PrismaClient, ctx context.Context) (*db.FileModel, error) {

	file, err := client.File.
		FindUnique(
			db.File.ID.Equals(
				id,
			),
		).
		Exec(ctx)

	return file, err
}
