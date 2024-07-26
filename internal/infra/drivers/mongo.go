package drivers

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type MongoDriver[T Entity] interface {
	FindAll() []T
	FindByNear(cordinates [2]int, maxDistanceInMeters int) []T
	FindByID(id string) (T, error)
	InsertOne(entity T) T
	UpdateOne(id string, updatedEntity T) (T, error)
	DeleteOne(id string) error
}

type Entity interface {
	GetID() string
	SetID(string)
}

type mongoDriver[T Entity] struct {
	entities map[string]T
	mu       sync.Mutex
}

func NewMongoDriver[T Entity]() MongoDriver[T] {
	return &mongoDriver[T]{
		entities: make(map[string]T),
	}
}

func (m *mongoDriver[T]) FindAll() []T {
	m.mu.Lock()
	defer m.mu.Unlock()

	result := make([]T, 0, len(m.entities))
	for _, entity := range m.entities {
		result = append(result, entity)
	}
	return result
}

func (m *mongoDriver[T]) FindByNear(cordinates [2]int, maxDistanceInMeters int) []T {
	m.mu.Lock()
	defer m.mu.Unlock()

	result := make([]T, 0, len(m.entities))
	for _, entity := range m.entities {
		result = append(result, entity)
	}

	return result
}

func (m *mongoDriver[T]) FindByID(id string) (T, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	entity, exists := m.entities[id]
	if !exists {
		var empty T
		return empty, errors.New("entity not found")
	}

	return entity, nil
}

func (m *mongoDriver[T]) InsertOne(entity T) T {
	m.mu.Lock()
	defer m.mu.Unlock()

	id := uuid.New().String()
	entity.SetID(id)
	m.entities[id] = entity
	return entity
}

func (m *mongoDriver[T]) UpdateOne(id string, updatedEntity T) (T, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exists := m.entities[id]
	if !exists {
		return updatedEntity, errors.New("entity not found")
	}

	updatedEntity.SetID(id)
	m.entities[id] = updatedEntity
	return updatedEntity, nil
}

func (m *mongoDriver[T]) DeleteOne(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exists := m.entities[id]
	if !exists {
		return errors.New("entity not found")
	}

	delete(m.entities, id)
	return nil
}
