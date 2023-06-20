package inventoryService

import (
	"adr/backend/src/prisma/db"
	"context"
	"fmt"
)

func Count(where string, client *db.PrismaClient, ctx context.Context) (int, error) {

	query := fmt.Sprintf(`SELECT COUNT(*) FROM "Inventory" %v`, where)

	var counts interface{}
	err := client.Prisma.QueryRaw(query).Exec(ctx, &counts)
	if err != nil {
		return 0, err
	}

	count := counts.([]interface{})[0].(map[string]interface{})["count"].(float64)
	if err != nil {
		return 0, err
	}

	return int(count), nil

}
