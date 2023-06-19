package service

import (
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"
)

type CourseService interface {
	InsertCourse(course dto.CourseCreateDTO) dto.CourseCreateDTO
	UpdateCourse(course dto.CourseUpdateDTO) dto.CourseUpdateDTO
	GetAllCourse() []models.Courses
	GetCourseByID(id uint64) models.Courses
	DeleteCourse(id uint64) error
}

type courseService struct {
	courseRepository repository.CourseRepository
}

//NewCourseService membuat instance baru dari CourseService.
func NewCourseService(courseRepo repository.CourseRepository) CourseService {
	return &courseService{
		courseRepository: courseRepo,
	}
}

//InsertCourse digunakan untuk menyisipkan data course baru ke dalam database.
func (service *courseService) InsertCourse(course dto.CourseCreateDTO) dto.CourseCreateDTO {
	courseToInsert := models.Courses{}
	courseToInsert.Code = course.Code
	courseToInsert.Name = course.Name
	courseInserted := service.courseRepository.InsertCourse(courseToInsert)

	var courseDTO dto.CourseCreateDTO
	courseDTO.Code = courseInserted.Code
	courseDTO.Name = courseInserted.Name
	return courseDTO
}

//UpdateCourse digunakan untuk memperbarui data course yang ada dalam database.
func (service *courseService) UpdateCourse(course dto.CourseUpdateDTO) dto.CourseUpdateDTO {
	courseToUpdate := models.Courses{}
	courseToUpdate.ID = uint64(course.ID)
	courseToUpdate.Code = course.Code
	courseToUpdate.Name = course.Name
	courseUpdated := service.courseRepository.UpdateCourse(courseToUpdate)

	var courseDTO dto.CourseUpdateDTO
	courseDTO.ID = courseUpdated.ID
	courseDTO.Code = courseUpdated.Code
	courseDTO.Name = courseUpdated.Name
	return courseDTO
}

//GetAllCourse digunakan untuk mendapatkan semua data course dari database.
func (service *courseService) GetAllCourse() []models.Courses {
	allCourse := service.courseRepository.GetAllCourse()
	return allCourse
}

//GetCourseByID digunakan untuk mendapatkan data course berdasarkan ID dari database.
func (service *courseService) GetCourseByID(id uint64) models.Courses {
	course := service.courseRepository.GetCourseByID(id)
	return course
}

//DeleteCourse digunakan untuk menghapus data course berdasarkan ID dari database.
func (service *courseService) DeleteCourse(id uint64) error {
	err := service.courseRepository.DeleteCourse(id)
	if err != nil {
		return err
	}
	return nil
}
