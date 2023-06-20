package productRouter

import (
	"adr/backend/src/graphql/model"
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/fileRelationService"
	"context"
	"log"
)

func Media(ctx context.Context, product *model.Product) ([]*string, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	fileRelation, err := fileRelationService.FindMany(product.ID, client, ctx)
	if err != nil {
		log.Println("error fileRelationService FindMany", err)
		return nil, err
	}
	mediaData := []*string{}

	for _, data := range fileRelation {

		file := data.File()
		mediaData = append(mediaData,
			&file.URL,
		)
	}

	return mediaData, nil
}
