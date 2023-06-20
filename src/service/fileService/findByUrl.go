package fileService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func FindByUrl(url string, client *db.PrismaClient, ctx context.Context) ([]db.FileModel, error) {

	file, err := client.File.
		FindMany(
			db.File.URL.Equals(
				url,
			),
		).
		Exec(ctx)

	return file, err
}
