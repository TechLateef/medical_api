// repository/medical_record_repository.go
package repository

import (
	model "medical_api/model"

	"gorm.io/gorm"
)

type MedicalRecordRepository interface {
	CreateMedicalRecord(record model.MedicalRecord) model.MedicalRecord
	UpdateMedicalRecord(record model.MedicalRecord, recordID uint64) model.MedicalRecord
	DeleteMedicalRecord(recordID uint64) model.MedicalRecord
	FindMedicalRecordByID(recordID uint64) model.MedicalRecord
	GetAllMedicalRecords() []model.MedicalRecord
}

type medicalRecordRepository struct {
	connection *gorm.DB
}

func NewMedicalRecordRepository(db *gorm.DB) MedicalRecordRepository {
	return &medicalRecordRepository{
		connection: db,
	}
}

func (repo *medicalRecordRepository) CreateMedicalRecord(record model.MedicalRecord) model.MedicalRecord {
	repo.connection.Save(&record)
	return record
}

func (repo *medicalRecordRepository) UpdateMedicalRecord(record model.MedicalRecord, recordID uint64) model.MedicalRecord {
	repo.connection.Model(&model.MedicalRecord{}).Where("id = ?", recordID).Updates(record)
	return record
}

func (repo *medicalRecordRepository) DeleteMedicalRecord(recordID uint64) model.MedicalRecord {
	var record model.MedicalRecord
	repo.connection.Delete(&record, recordID)
	return record
}

func (repo *medicalRecordRepository) FindMedicalRecordByID(recordID uint64) model.MedicalRecord {
	var record model.MedicalRecord
	repo.connection.Preload("Patient").Preload("Doctor").First(&record, recordID)
	return record
}

func (repo *medicalRecordRepository) GetAllMedicalRecords() []model.MedicalRecord {
	var records []model.MedicalRecord
	repo.connection.Preload("Patient").Preload("Doctor").Find(&records)
	return records
}
