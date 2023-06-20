package productVariationRouter

import (
	srcModel "adr/backend/src/model"
	"adr/backend/src/prisma/db"
	"adr/backend/src/service/productVariationService"
	"context"
	"log"
)

func FindTitle(ctx context.Context, id string) (*string, error) {

	client := ctx.Value(srcModel.ConfigKey("client")).(*db.PrismaClient)

	productVariationData, err := productVariationService.Find(id, client, ctx)
	if err != nil {
		log.Println("error productVariationService Find", err)
		return nil, err
	}
	if productVariationData != nil {
		return &productVariationData.Title, nil
	}
	return nil, nil
}
