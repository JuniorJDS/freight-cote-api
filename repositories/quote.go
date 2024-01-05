package repositories

import (
	"context"
	"freight-cote-api/databases"
	r "freight-cote-api/schemas/response"
	"log"
	"time"

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

	_, err := qr.quoteCollection.InsertOne(ctx, quote.Carrier[0])
	if err != nil {
		log.Printf("Error to insert quote: %s\n", err.Error())
		return err
	}
	return nil
}

func (qr *QuoteRepository) Get() {

}

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Shipping struct {
// 	Name     string  `bson:"name"`
// 	Service  string  `bson:"service"`
// 	Deadline int     `bson:"deadline"`
// 	Price    float64 `bson:"price"`
// }

// func main() {
// 	// Replace the connection string and collection name with your actual values
// 	connectionString := "your_connection_string"
// 	collectionName := "your_collection_name"

// 	// Set up MongoDB client
// 	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	// Access the collection
// 	collection := client.Database("your_database_name").Collection(collectionName)

// 	// Find documents
// 	findOptions := options.Find()
// 	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cursor.Close(ctx)

// 	// Create a slice to hold the found documents
// 	var shipments []Shipping
// 	if err := cursor.All(ctx, &shipments); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Aggregate the results in memory
// 	aggregatedResult := aggregateShipments(shipments)

// 	fmt.Printf("Aggregated Result: %+v\n", aggregatedResult)
// }

// func aggregateShipments(shipments []Shipping) map[string]interface{} {
// 	// Aggregation logic on the in-memory slice
// 	aggregatedResult := make(map[string]interface{})

// 	// Example: Count occurrences of each shipping name
// 	nameCounts := make(map[string]int)
// 	for _, shipment := range shipments {
// 		nameCounts[shipment.Name]++
// 	}
// 	aggregatedResult["nameCounts"] = nameCounts

// 	// Example: Calculate total price and average price per name
// 	nameTotalPrice := make(map[string]float64)
// 	nameCount := make(map[string]int)
// 	for _, shipment := range shipments {
// 		nameTotalPrice[shipment.Name] += shipment.Price
// 		nameCount[shipment.Name]++
// 	}
// 	aggregatedResult["nameTotalPrice"] = nameTotalPrice
// 	aggregatedResult["nameAvgPrice"] = calculateAveragePrice(nameTotalPrice, nameCount)

// 	// Example: Find cheapest and most expensive shipping overall
// 	cheapestShipping, mostExpensiveShipping := findCheapestAndMostExpensive(shipments)
// 	aggregatedResult["cheapestShipping"] = cheapestShipping
// 	aggregatedResult["mostExpensiveShipping"] = mostExpensiveShipping

// 	return aggregatedResult
// }

// func calculateAveragePrice(totalPriceMap map[string]float64, countMap map[string]int) map[string]float64 {
// 	avgPriceMap := make(map[string]float64)
// 	for name, totalPrice := range totalPriceMap {
// 		count := countMap[name]
// 		avgPriceMap[name] = totalPrice / float64(count)
// 	}
// 	return avgPriceMap
// }

// func findCheapestAndMostExpensive(shipments []Shipping) (Shipping, Shipping) {
// 	var cheapest, mostExpensive Shipping
// 	for i, shipment := range shipments {
// 		if i == 0 || shipment.Price < cheapest.Price {
// 			cheapest = shipment
// 		}
// 		if i == 0 || shipment.Price > mostExpensive.Price {
// 			mostExpensive = shipment
// 		}
// 	}
// 	return cheapest, mostExpensive
// }
