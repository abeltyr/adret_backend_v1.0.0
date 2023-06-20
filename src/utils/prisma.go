package utils

import (
	"adr/backend/src/prisma/db"
	"context"
	"log"
)

func PrismaClient() (*db.PrismaClient, context.Context, error) {

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Println("prisma setup error:", err)
		return nil, nil, err
	}

	ctx := context.Background()

	return client, ctx, nil
}
