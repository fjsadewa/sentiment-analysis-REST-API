package service

import (
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"
)

/*
StudyProgramService adalah sebuah interface yang mendefinisikan method-method
untuk menangani operasi terkait program studi.
*/
type StudyProgramService interface {
	InsertProdi(studyProgram dto.StudyProgramCreateDTO) models.StudyPrograms
	UpdateProdi(studyProgram dto.StudyProgramUpdateDTO) models.StudyPrograms
	DeleteProdi(id uint64) error
	AllProdi() []models.StudyPrograms
	FindProdiByID(id uint64) (models.StudyPrograms, error)
	FindProdiByCode(Code string) models.StudyPrograms
}

//studyProgramService adalah sebuah struct yang mengimplementasikan interface StudyProgramService.
type studyProgramService struct {
	studyProgramRepository repository.StudyProgramRepository
}

//NewStudyProgramService membuat sebuah instance baru dari studyProgramService.
func NewStudyProgramService(studyProgramRepo repository.StudyProgramRepository) StudyProgramService {
	return &studyProgramService{
		studyProgramRepository: studyProgramRepo,
	}
}

//InsertProdi memasukkan program studi baru ke dalam database.
func (service *studyProgramService) InsertProdi(studyProgram dto.StudyProgramCreateDTO) models.StudyPrograms {
	//Membuat objek model StudyPrograms baru dan mengisinya dengan data yang diberikan
	studyProgramToInsert := models.StudyPrograms{}
	studyProgramToInsert.Code = studyProgram.Code
	studyProgramToInsert.Name = studyProgram.Name

	//Memasukkan program studi ke dalam database menggunakan studyProgramRepository
	studyProgramInserted := service.studyProgramRepository.InsertStudyProgram(studyProgramToInsert)
	return studyProgramInserted
}

//UpdateProdi mengubah data program studi yang ada dalam database.
func (service *studyProgramService) UpdateProdi(studyProgram dto.StudyProgramUpdateDTO) models.StudyPrograms {
	//Membuat objek model StudyPrograms baru dan mengisinya dengan data yang diberikan
	studyProgramToUpdate := models.StudyPrograms{}
	studyProgramToUpdate.ID = uint64(studyProgram.ID)
	studyProgramToUpdate.Code = studyProgram.Code
	studyProgramToUpdate.Name = studyProgram.Name

	//Mengubah data program studi dalam database menggunakan studyProgramRepository
	studyProgramUpdated := service.studyProgramRepository.UpdateStudyProgram(studyProgramToUpdate)
	return studyProgramUpdated
}

//DeleteProdi menghapus data program studi berdasarkan ID dari database.
func (service *studyProgramService) DeleteProdi(id uint64) error {
	// Menghapus data program studi berdasarkan ID dari database menggunakan studyProgramRepository
	err := service.studyProgramRepository.DeleteStudyProgram(id)
	if err != nil {
		return err
	}
	return nil
}

//AllProdi mengambil semua data program studi dari database.
func (service *studyProgramService) AllProdi() []models.StudyPrograms {
	allStudyProgram := service.studyProgramRepository.AllStudyProgram()
	return allStudyProgram
}

//FindProdiByID mengambil data program studi berdasarkan ID dari database.
func (s *studyProgramService) FindProdiByID(id uint64) (models.StudyPrograms, error) {
	studyProgram, err := s.studyProgramRepository.FindStudyProgramByID(id)
	if err != nil {
		return studyProgram, err
	}
	return studyProgram, nil
}

// FindProdiByCode mengambil data program studi berdasarkan kode dari database.
func (service *studyProgramService) FindProdiByCode(Code string) models.StudyPrograms {
	studyProgram := service.studyProgramRepository.FindStudyProgramByCode(Code)
	return studyProgram
}
