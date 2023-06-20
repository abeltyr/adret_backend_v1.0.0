package tableCountService

import (
	"adr/backend/src/prisma/db"
	"context"
)

func Create(countName string, increment *int, client *db.PrismaClient, ctx context.Context) (*db.TableCountModel, error) {

	incrementData := 1

	if increment != nil {
		incrementData = *increment
	}

	tableCount, err := client.TableCount.UpsertOne(
		db.TableCount.ID.Equals(countName),
	).Create(
		db.TableCount.ID.Set(countName),
		db.TableCount.Count.Set(incrementData),
	).Update(
		db.TableCount.Count.Increment(incrementData),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return tableCount, nil

}
