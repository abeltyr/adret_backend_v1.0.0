package inventoryRouter

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/fileRelationService"
	"context"
	"log"
)

func Media(ctx context.Context, id string) (*string, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	fileRelation, err := fileRelationService.FindMany(id, client, ctx)
	if err != nil {
		log.Println("error fileRelationService FindMany", err)
		return nil, err
	}

	if len(fileRelation) > 0 {
		return &fileRelation[0].File().URL, nil
	} else {
		return nil, nil

	}
}
