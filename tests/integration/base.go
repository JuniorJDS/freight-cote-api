package integration_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	api "freight-cote-api/api/app"
	"freight-cote-api/databases"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseTest struct {
	suite.Suite
	database *mongo.Database
}

func NewBaseTest() *BaseTest {
	return &BaseTest{
		database: databases.GetDB(),
	}
}

func (b *BaseTest) appClient(
	verb string, pathEndpoint string, body io.Reader,
) (*http.Response, error) {
	baseURL := os.Getenv("APP_BASE_URL")
	url := baseURL + pathEndpoint

	req, err := http.NewRequest(verb, url, body)
	if err != nil {
		fmt.Println("Erro to make request: ", err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	app := api.MakeApp()
	resp, err := app.Test(req, -1)

	return resp, err
}

func (b *BaseTest) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collections, err := b.database.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		fmt.Println("Error to list collections in TearDown: ", err.Error())
	}

	for i := range collections {
		collectionName := collections[i]
		collection := b.database.Collection(collectionName)

		_, err = collection.DeleteMany(ctx, bson.M{})
		if err != nil {
			fmt.Println("Error to delete collection ", collectionName, " documents: ", err.Error())
		}

		_, err = collection.Indexes().DropAll(ctx)
		if err != nil {
			fmt.Println("Error to delete collection ", collectionName, " indexes: ", err.Error())
		}
	}
}
