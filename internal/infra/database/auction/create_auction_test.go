package auction

import (
	"context"
	"fullcycle-auction_go/configuration/database/mongodb"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUpdateAuctionStatusCompleted(t *testing.T) {

	ctx := context.Background()

	if err := godotenv.Load("../../../../cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx, "MONGODB_URL")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	auctionRepositoryTeste := NewAuctionRepository(databaseConnection)

	idTeste := uuid.New().String()

	auctionRepositoryTeste.CreateAuction(ctx, &auction_entity.Auction{
		Id:          idTeste,
		ProductName: "Produto Teste",
		Category:    "Categoria Teste",
		Description: "Descrição Teste",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	})

	auctionInserido, _ := auctionRepositoryTeste.FindAuctionById(ctx, idTeste)
	assert.Equal(t, auction_entity.Active, auctionInserido.Status)

	time.Sleep(auctionRepositoryTeste.AuctionDuration)

	auctionAtualizado, _ := auctionRepositoryTeste.FindAuctionById(ctx, idTeste)
	assert.Equal(t, auction_entity.Active, auctionAtualizado.Status)

}
