package gateways

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/drivers"
)

type ClientRepository interface {
	FindClientByID(id string) (entities.Client, error)
	InsertClient(client entities.Client) error
	UpdateClient(client entities.Client) (entities.Client, error)
	DeleteClient(id string) error
}

type clientRepository struct {
	dynamodbDriver drivers.DynamoDBDriver[entities.Client]
}

func NewClientRepository(dynamodbDriver drivers.DynamoDBDriver[entities.Client]) ClientRepository {
	return &clientRepository{
		dynamodbDriver: dynamodbDriver,
	}
}

func (r *clientRepository) FindClientByID(id string) (entities.Client, error) {
	return r.dynamodbDriver.FindOne(id, "")
}

func (r *clientRepository) InsertClient(client entities.Client) error {
	_ = r.dynamodbDriver.InsertOne(client)
	return nil
}

func (r *clientRepository) UpdateClient(client entities.Client) (entities.Client, error) {
	return r.dynamodbDriver.UpdateOne(client)
}

func (r *clientRepository) DeleteClient(id string) error {
	return r.dynamodbDriver.DeleteOne(id)
}
