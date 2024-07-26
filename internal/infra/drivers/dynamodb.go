package drivers

import (
	"errors"
	"sync"
)

type DynamoDBDriver[T DynamoDBEntity] interface {
	FindOne(partitionKey, sortKey string) (T, error)
	FindMany(partitionKey string) ([]T, error)
	InsertOne(entity T) T
	UpdateOne(entity T) (T, error)
	DeleteOne(id string) error
}

type DynamoDBEntity interface {
	GetHashKey() string
	GetSortKey() string
}

type dynamodbDriver[T DynamoDBEntity] struct {
	items             map[string]T
	collectionOfItems map[string][]T
	mu                sync.Mutex
}

func NewDynamoDBDriver[T DynamoDBEntity]() DynamoDBDriver[T] {
	return &dynamodbDriver[T]{
		items: make(map[string]T),
	}
}

func (d *dynamodbDriver[T]) FindOne(partitionKey, sortKey string) (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	entity, exists := d.items[partitionKey+"-"+sortKey]
	if !exists {
		var empty T
		return empty, errors.New("entity not found")
	}

	return entity, nil
}

func (d *dynamodbDriver[T]) FindMany(partitionKey string) ([]T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	items, exists := d.collectionOfItems[partitionKey]
	if !exists {
		return nil, errors.New("entity not found")
	}

	return items, nil
}

func (d *dynamodbDriver[T]) InsertOne(entity T) T {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.items[entity.GetHashKey()+""+entity.GetSortKey()] = entity
	d.collectionOfItems[entity.GetHashKey()] = append(d.collectionOfItems[entity.GetHashKey()], entity)
	return entity
}

func (d *dynamodbDriver[T]) UpdateOne(updatedEntity T) (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	_, exists := d.items[updatedEntity.GetHashKey()+"-"+updatedEntity.GetSortKey()]
	if !exists {
		return updatedEntity, errors.New("entity not found")
	}

	d.items[updatedEntity.GetHashKey()+"-"+updatedEntity.GetSortKey()] = updatedEntity
	return updatedEntity, nil
}

func (m *dynamodbDriver[T]) DeleteOne(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exists := m.items[id]
	if !exists {
		return errors.New("entity not found")
	}

	delete(m.items, id)
	return nil
}
