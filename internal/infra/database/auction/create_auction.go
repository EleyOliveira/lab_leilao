package auction

import (
	"context"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}
type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	go ar.ControlTimeUpdateStatusComplete(ctx, auctionEntity.Id)

	return nil
}

func (ar *AuctionRepository) UpdateAuctionStatusCompleted(
	ctx context.Context,
	id string) *internal_error.InternalError {

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": auction_entity.Completed}}

	_, err := ar.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return internal_error.NewInternalServerError("Error trying to update status auction to completed")
	}

	return nil

}

func (ar *AuctionRepository) ControlTimeUpdateStatusComplete(
	ctx context.Context,
	id string) {

	ticker := time.NewTicker(getAuctionInterval())
	defer ticker.Stop()

	done := make(chan bool)
	defer close(done)

	for {
		select {
		case <-ticker.C:
			err := ar.UpdateAuctionStatusCompleted(ctx, id)
			if err != nil {
				logger.Error("Error trying to update status auction to completed", err)
			}
			done <- true
		case <-done:
			return
		}
	}
}

func getAuctionInterval() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(auctionInterval)
	if err != nil {
		return time.Minute * 5
	}

	return duration
}
