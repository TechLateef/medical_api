package services

import (
	"log"
	"medical_api/dto"
	"medical_api/model"
	"medical_api/repository"

	"github.com/mashingan/smapping"
)

type MedicalRecordService interface {
	CreateMedicalRecord(record dto.CreateMedicalRecordDTO) model.MedicalRecord
	UpdateMedicalRecord(record dto.UpdateMedicalRecordDTO, recordID uint64) model.MedicalRecord
	DeleteMedicalRecord(recordID uint64) model.MedicalRecord
	FindMedicalRecordByID(recordID uint64) model.MedicalRecord
	GetAllMedicalRecords() []model.MedicalRecord
}

type medicalRecordService struct {
	medicalRecordRepository repository.MedicalRecordRepository
}

func NewMedicalRecordService(repo repository.MedicalRecordRepository) MedicalRecordService {
	return &medicalRecordService{
		medicalRecordRepository: repo,
	}
}

func (service *medicalRecordService) CreateMedicalRecord(record dto.CreateMedicalRecordDTO) model.MedicalRecord {
	medicalRecord := model.MedicalRecord{}
	err := smapping.FillStruct(&medicalRecord, smapping.MapFields(&record))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	return service.medicalRecordRepository.CreateMedicalRecord(medicalRecord)
}

func (service *medicalRecordService) UpdateMedicalRecord(record dto.UpdateMedicalRecordDTO, recordID uint64) model.MedicalRecord {
	medicalRecord := model.MedicalRecord{}
	err := smapping.FillStruct(&medicalRecord, smapping.MapFields(&record))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	return service.medicalRecordRepository.UpdateMedicalRecord(medicalRecord, recordID)
}

func (service *medicalRecordService) DeleteMedicalRecord(recordID uint64) model.MedicalRecord {
	return service.medicalRecordRepository.DeleteMedicalRecord(recordID)
}

func (service *medicalRecordService) FindMedicalRecordByID(recordID uint64) model.MedicalRecord {
	return service.medicalRecordRepository.FindMedicalRecordByID(recordID)
}

func (service *medicalRecordService) GetAllMedicalRecords() []model.MedicalRecord {
	return service.medicalRecordRepository.GetAllMedicalRecords()
}
