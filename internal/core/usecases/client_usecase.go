package usecases

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/gateways"
)

type ClientUseCase interface {
	GetClient(email string) (entities.Client, error)
	Createclient(client entities.Client) error
	UpdateClient(client entities.Client) error
	DeleteClient(id string) error
}

type clientUserCase struct {
	clientRepository gateways.ClientRepository
}

func NewClientUseCase(clientRepository gateways.ClientRepository) ClientUseCase {
	return &clientUserCase{
		clientRepository: clientRepository,
	}
}

func (u *clientUserCase) GetClient(email string) (entities.Client, error) {
	return u.clientRepository.FindClientByID(email)
}

func (u *clientUserCase) Createclient(client entities.Client) error {
	return u.clientRepository.InsertClient(client)
}

func (u *clientUserCase) UpdateClient(client entities.Client) error {
	_, err := u.clientRepository.UpdateClient(client)
	return err
}

func (u *clientUserCase) DeleteClient(id string) error {
	return u.clientRepository.DeleteClient(id)
}
