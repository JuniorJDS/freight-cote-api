package repositories

import (
	"context"
	"freight-cote-api/databases"
	r "freight-cote-api/schemas/response"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuoteRepository struct {
	quoteCollection mongo.Collection
}

func NewQuoteRepository() *QuoteRepository {
	repository := &QuoteRepository{
		quoteCollection: *databases.GetCollection("quote"),
	}
	return repository
}

func (qr *QuoteRepository) Create(quote r.QuoteResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var documents []interface{}
	for _, carrier := range quote.Carrier {
		documents = append(documents, carrier)
	}

	_, err := qr.quoteCollection.InsertMany(ctx, documents)
	if err != nil {
		log.Printf("Error to insert quote: %s\n", err.Error())
		return err
	}
	return nil
}

func (qr *QuoteRepository) GetMetrics(lastQuote int64) (*r.Metrics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	metrics := r.Metrics{}

	pipeline := []bson.M{
		{
			"$sort": bson.M{"_id": -1},
		},
	}

	if lastQuote > 0 {
		limitStage := bson.M{"$limit": lastQuote}
		pipeline = append(pipeline, limitStage)
	}

	groupStages := []bson.M{
		{
			"$group": bson.M{
				"_id":           "$name",
				"count":         bson.M{"$sum": 1},
				"total":         bson.M{"$sum": "$price"},
				"avg":           bson.M{"$avg": "$price"},
				"highest_price": bson.M{"$max": "$price"},
				"lowest_price":  bson.M{"$min": "$price"},
			},
		},
		{
			"$group": bson.M{
				"_id": nil,
				"by_carriers": bson.M{
					"$push": bson.M{
						"name":     "$_id",
						"quantity": "$count",
						"total":    "$total",
						"avg":      "$avg",
					},
				},
				"highest_price": bson.M{"$max": "$highest_price"},
				"lowest_price":  bson.M{"$min": "$lowest_price"},
			},
		},
	}
	pipeline = append(pipeline, groupStages...)

	cursor, err := qr.quoteCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Printf("Error to Aggregate Metrics: %s\n", err.Error())
		return nil, err
	}

	for cursor.Next(ctx) {
		err := cursor.Decode(&metrics)
		if err != nil {
			log.Printf("Failed to get metrics in filterCursor: %s\n", err.Error())
			return nil, err
		}
	}

	return &metrics, nil
}
