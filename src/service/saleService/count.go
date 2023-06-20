// package

// import (
// 	"adr/backend/src/prisma/db"
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func Count(inventoryId string, client *db.PrismaClient) (int, error) {

// 	databaseUrl := os.Getenv("DATABASE_URL")

// 	mongoDBClient, err := mongo.NewClient(options.Client().ApplyURI(databaseUrl))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

// 	err = mongoDBClient.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer mongoDBClient.Disconnect(ctx)

// 	collection := mongoDBClient.Database("adret").Collection("")

// 	fmt.Println(collection)

// 	count, err := collection.CountDocuments(ctx,
// 		bson.D{
// 			{Key: "inventoryId", Value: inventoryId},
// 		})

// 	return int(count), err

// }

package saleService

import (
	"adr/backend/src/prisma/db"
	"context"
	"fmt"
)

func Count(where string, client *db.PrismaClient, ctx context.Context) (int, error) {

	query := fmt.Sprintf(`SELECT COUNT(*) FROM "Sale" %v`, where)

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
