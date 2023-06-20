package tableCountService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Find(countName string, client *db.PrismaClient, ctx context.Context) (*db.TableCountModel, error) {
	tableCount, err := client.TableCount.FindUnique(
		db.TableCount.ID.Equals(countName),
	).Exec(ctx)

	return tableCount, err
}
