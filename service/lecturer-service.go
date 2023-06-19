package service

import (
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"
)

/*
LecturerService adalah sebuah interface yang mendefinisikan method-method untuk menangani operasi terkait dosen.
*/
type LecturerService interface {
	InsertLecturer(lecturer dto.LecturerCreateDTO) dto.LecturerCreateDTO
	UpdateLecturer(lecturer dto.LecturerUpdateDTO) dto.LecturerUpdateDTO
	GetAllLecturer() []models.Lecturers
	GetLecturerByID(id uint64) models.Lecturers
	DeleteLecturer(id uint64) error
}

//lecturerService adalah sebuah struct yang mengimplementasikan interface LecturerService.
type lecturerService struct {
	lecturerRepository repository.LecturerRepository
}

//NewLecturerService membuat sebuah instans baru dari lecturerService.
func NewLecturerService(lecturerRepo repository.LecturerRepository) LecturerService {
	return &lecturerService{
		lecturerRepository: lecturerRepo,
	}
}

//InsertLecturer memasukkan seorang dosen baru ke dalam database.
func (service *lecturerService) InsertLecturer(lecturer dto.LecturerCreateDTO) dto.LecturerCreateDTO {
	//Membuat objek model Lecturers baru dan mengisinya dengan data yang diberikan
	lecturerToInsert := models.Lecturers{}
	lecturerToInsert.Code = lecturer.Code
	lecturerToInsert.Name = lecturer.Name
	lecturerToInsert.StudyProgramID = lecturer.StudyProgramID

	//Memasukkan dosen ke dalam database menggunakan lecturerRepository
	lecturerInserted := service.lecturerRepository.InsertLecturer(lecturerToInsert)

	//Membuat objek DTO LecturerCreateDTO baru dan mengisinya dengan data dosen yang dimasukkan
	var lecturerDTO dto.LecturerCreateDTO
	lecturerDTO.Code = lecturerInserted.Code
	lecturerDTO.Name = lecturerInserted.Name
	lecturerDTO.StudyProgramID = lecturerInserted.StudyProgramID
	return lecturerDTO
}

//UpdateLecturer mengubah data seorang dosen yang ada dalam database.
func (service *lecturerService) UpdateLecturer(lecturer dto.LecturerUpdateDTO) dto.LecturerUpdateDTO {
	//Membuat objek model Lecturers baru dan mengisinya dengan data yang diberikan
	lecturerToUpdate := models.Lecturers{}
	lecturerToUpdate.ID = uint64(lecturer.ID)
	lecturerToUpdate.Code = lecturer.Code
	lecturerToUpdate.Name = lecturer.Name
	lecturerToUpdate.StudyProgramID = lecturer.StudyProgramID

	//Mengubah data dosen dalam database menggunakan lecturerRepository
	lecturerUpdated := service.lecturerRepository.UpdateLecturer(lecturerToUpdate)

	//Membuat objek DTO LecturerUpdateDTO baru dan mengisinya dengan data dosen yang diperbarui
	var lecturerDTO dto.LecturerUpdateDTO
	lecturerDTO.ID = lecturerUpdated.ID
	lecturerDTO.Code = lecturerUpdated.Code
	lecturerDTO.Name = lecturerUpdated.Name
	lecturerDTO.StudyProgramID = lecturerUpdated.StudyProgramID
	return lecturerDTO
}

//GetAllLecturer mengambil semua data dosen dari database.
func (service *lecturerService) GetAllLecturer() []models.Lecturers {
	allLecturer := service.lecturerRepository.GetAllLecturer()
	return allLecturer
}

//GetLecturerByID mengambil data seorang dosen berdasarkan ID dari database.
func (service *lecturerService) GetLecturerByID(id uint64) models.Lecturers {
	lecturer := service.lecturerRepository.GetLecturerByID(id)
	return lecturer
}

//DeleteLecturer menghapus data seorang dosen berdasarkan ID dari database.
func (service *lecturerService) DeleteLecturer(id uint64) error {
	err := service.lecturerRepository.DeleteLecturer(id)
	if err != nil {
		return err
	}
	return nil
}
