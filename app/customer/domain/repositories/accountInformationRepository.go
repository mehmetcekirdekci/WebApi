package repositories

import (
	"context"
	"errors"

	"github.com/mehmetcekirdekci/WebApi/app/customer/domain/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountInformationRepository interface {
	InsertAccountInformation(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
}

type (
	accountInformationRepository struct {
		accountInformationsCollection *mongo.Collection
	}
)

func NewAccountInformationRepository(database *mongo.Database) AccountInformationRepository {
	return &accountInformationRepository{
		accountInformationsCollection: database.Collection(types.AccountInformationsTable),
	}
}

func (receiver *accountInformationRepository) InsertAccountInformation(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	result, err := receiver.accountInformationsCollection.InsertOne(ctx, document)
	if err != nil {
		return result, errors.New("Account information can not be inserted.")
	}
	return result, nil
}