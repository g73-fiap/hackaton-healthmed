package gateways

import (
	"errors"
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/drivers"
	"mime/multipart"
)

type ClientRepository interface {
	FindClientByID(id string) (entities.Client, error)
	InsertClient(client entities.Client) error
	UpdateClient(client entities.Client) (entities.Client, error)
	DeleteClient(id string) error
	InsertMedicalReport(name, fileName string, file *multipart.FileHeader) error
	DeleteMedicalReport(name, fileName string) error
	FindMedicalReport(name, fileName, requester string) (string, error)
}

type clientRepository struct {
	dynamodbDriver drivers.DynamoDBDriver[entities.Client]
	s3Driver       drivers.S3Driver
}

func NewClientRepository(dynamodbDriver drivers.DynamoDBDriver[entities.Client], s3Driver drivers.S3Driver) ClientRepository {
	return &clientRepository{
		dynamodbDriver: dynamodbDriver,
		s3Driver:       s3Driver,
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

func (r *clientRepository) InsertMedicalReport(name, fileName string, file *multipart.FileHeader) error {
	location, err := r.s3Driver.PutObject(name, file)
	if err != nil {
		return err
	}

	client, err := r.FindClientByID(name)
	if err != nil {
		return err
	}

	medicalReport := entities.MedicalReport{
		SharedWith:   []string{},
		FileName:     fileName,
		FileLocation: location,
	}

	client.MedicalReports = append(client.MedicalReports, medicalReport)

	client, err = r.UpdateClient(client)
	if err != nil {
		return err
	}

	return nil
}
func (r *clientRepository) DeleteMedicalReport(email, fileName string) error {
	client, err := r.FindClientByID(email)
	if err != nil {
		return err
	}

	var location string
	var index int
	for i, report := range client.MedicalReports {
		if report.FileName == fileName {
			location = report.FileLocation
			index = i
			break
		}
	}

	if location == "" {
		return errors.New("file not found")
	}

	newReports := removeElementByIndex(client.MedicalReports, index)
	client.MedicalReports = newReports

	_, err = r.UpdateClient(client)
	if err != nil {
		return err
	}

	err = r.s3Driver.RemoveObject(location)
	if err != nil {
		return err
	}

	return nil
}
func (r *clientRepository) FindMedicalReport(email, fileName, requester string) (string, error) {
	client, err := r.FindClientByID(email)
	if err != nil {
		return "", err
	}

	var medicalReport entities.MedicalReport
	for _, report := range client.MedicalReports {
		if report.FileName == fileName {
			medicalReport = report
			break
		}
	}

	if medicalReport.FileLocation == "" {
		return "", errors.New("file not found")
	}

	for _, doctor := range medicalReport.SharedWith {
		if doctor == requester {
			return medicalReport.FileLocation, nil
		}
	}

	return "", errors.New("access not authorized")
}

func removeElementByIndex[T any](slice []T, index int) []T {
	sliceLen := len(slice)
	sliceLastIndex := sliceLen - 1

	if index != sliceLastIndex {
		slice[index] = slice[sliceLastIndex]
	}

	return slice[:sliceLastIndex]
}
