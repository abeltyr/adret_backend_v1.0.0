package fileService

import (
	"adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"context"
)

func Create(uploadedData model.UploadedData, client *db.PrismaClient, ctx context.Context) (*db.FileModel, error) {
	createFile, err := client.File.
		CreateOne(
			db.File.URL.Set(uploadedData.URL),
			db.File.ContentType.Set(uploadedData.ContentType),
			db.File.PreviewURL.Set(uploadedData.PreviewURL),
			db.File.Width.Set(uploadedData.Width),
			db.File.Name.Set(uploadedData.Name),
			db.File.Uploader.Set(uploadedData.Uploader),
			db.File.Height.Set(uploadedData.Height),
			db.File.Size.Set(uploadedData.Size),
		).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createFile, nil

}
