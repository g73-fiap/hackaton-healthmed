package usecases

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/gateways"
	"mime/multipart"
)

type ClientUseCase interface {
	GetClient(email string) (entities.Client, error)
	Createclient(client entities.Client) error
	UpdateClient(client entities.Client) error
	DeleteClient(id string) error
	SaveMedicalReport(email string, name string, file *multipart.FileHeader) error
	RemoveMedicalReport(email, fileName string) error
	GetMedicalReport(email, fileName, requester string) (string, error)
	ShareMedicalReport(email, fileName, request string) error
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

func (u *clientUserCase) SaveMedicalReport(email string, name string, file *multipart.FileHeader) error {
	return u.clientRepository.InsertMedicalReport(email, name, file)
}

func (u *clientUserCase) RemoveMedicalReport(email, fileName string) error {
	return u.clientRepository.DeleteMedicalReport(email, fileName)
}

func (u *clientUserCase) GetMedicalReport(email, fileName, requester string) (string, error) {
	return u.clientRepository.FindMedicalReport(email, fileName, requester)
}

func (u *clientUserCase) ShareMedicalReport(email, fileName, requester string) error {
	client, err := u.GetClient(email)
	if err != nil {
		return err
	}

	var newMedicalReport entities.MedicalReport
	var index int
	for i, report := range client.MedicalReports {
		if report.FileName == fileName {
			index = i
			newMedicalReport = report
			break
		}
	}

	currentReport := &client.MedicalReports[index]
	currentReport.SharedWith = append(newMedicalReport.SharedWith, requester)

	err = u.UpdateClient(client)
	if err != nil {
		return err
	}

	return nil
}
