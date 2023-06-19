package service

import (
	"fmt"
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"

	gt "github.com/bas24/googletranslatefree"
	"github.com/jonreiter/govader"
)

//CommentService adalah interface yang mendefinisikan operasi-operasi yang terkait dengan komentar.
type CommentService interface {
	InsertComment(comment dto.CommentCreateDTO) dto.CommentCreateDTO
	UpdateComment(comment dto.CommentUpdateDTO) dto.CommentUpdateDTO
	GetAllComment() []dto.CommentDTO
	GetCommentByID(id uint64) dto.CommentDTO
	DeleteComment(id uint64) error
}

//commentService adalah struct yang mengimplementasikan interface CommentService.
type commentService struct {
	commentRepository  repository.CommentRepository
	lecturerRepository repository.LecturerRepository
}

//NewCommentServiceWithLecturer digunakan untuk membuat instance baru dari CommentService yang memiliki repository komentar dan repository dosen.
func NewCommentServiceWithLecturer(commentRepo repository.CommentRepository, lecturerRepo repository.LecturerRepository) CommentService {
	return &commentService{
		commentRepository:  commentRepo,
		lecturerRepository: lecturerRepo,
	}
}

//InsertComment digunakan untuk menyimpan komentar baru ke dalam database.
func (service *commentService) InsertComment(comment dto.CommentCreateDTO) dto.CommentCreateDTO {
	commentToInsert := models.Comments{}
	commentToInsert.Comment = comment.Comment
	commentToInsert.CourseID = uint64(comment.CourseID)
	commentToInsert.LecturerID = uint64(comment.LecturerID)
	commentToInsert.SchoolYear = comment.SchoolYear
	commentToInsert.Semester = comment.Semester
	commentInserted := service.commentRepository.InsertComment(commentToInsert)

	// insert to sentiment analysis
	sentimentAnalysisToInsert := models.SentimentAnalysis{}
	result, err := gt.Translate(comment.Comment, "id", "en")
	if err != nil {
		fmt.Println(err)
	}
	analyzer := govader.NewSentimentIntensityAnalyzer()
	sentiment := analyzer.PolarityScores(result)

	if sentiment.Compound > 0.05 {
		sentimentAnalysisToInsert.SentimentType = "Positive"
	} else if sentiment.Compound < -0.05 {
		sentimentAnalysisToInsert.SentimentType = "Negative"
	} else {
		sentimentAnalysisToInsert.SentimentType = "Neutral"
	}
	sentimentAnalysisToInsert.CommentID = uint(commentInserted.ID)
	sentimentAnalysisToInsert.ConfidenceLevel = sentiment.Compound
	service.commentRepository.InsertToSentimentAnalysis(sentimentAnalysisToInsert)

	var commentDTO dto.CommentCreateDTO
	commentDTO.SentimentType = sentimentAnalysisToInsert.SentimentType
	commentDTO.ConfidenceLevel = sentimentAnalysisToInsert.ConfidenceLevel
	commentDTO.Comment = commentInserted.Comment
	commentDTO.CourseID = commentInserted.CourseID
	commentDTO.LecturerID = commentInserted.LecturerID
	commentDTO.SchoolYear = commentInserted.SchoolYear
	commentDTO.Semester = commentInserted.Semester
	return commentDTO
}

//UpdateComment digunakan untuk mengupdate komentar yang sudah ada di database.
func (service *commentService) UpdateComment(comment dto.CommentUpdateDTO) dto.CommentUpdateDTO {
	commentToUpdate := models.Comments{}
	commentToUpdate.ID = uint64(comment.ID)
	commentToUpdate.Comment = comment.Comment
	commentToUpdate.CourseID = uint64(comment.CourseID)
	commentToUpdate.LecturerID = uint64(comment.LecturerID)
	commentToUpdate.SchoolYear = comment.SchoolYear
	commentToUpdate.Semester = comment.Semester
	commentUpdated := service.commentRepository.UpdateComment(commentToUpdate)

	// insert to sentiment analysis
	sentimentAnalysisToInsert := models.SentimentAnalysis{}
	result, err := gt.Translate(comment.Comment, "id", "en")
	if err != nil {
		fmt.Println(err)
	}
	analyzer := govader.NewSentimentIntensityAnalyzer()
	sentiment := analyzer.PolarityScores(result)

	if sentiment.Compound > 0.05 {
		sentimentAnalysisToInsert.SentimentType = "Positive"
	} else if sentiment.Compound < -0.05 {
		sentimentAnalysisToInsert.SentimentType = "Negative"
	} else {
		sentimentAnalysisToInsert.SentimentType = "Neutral"
	}
	sentimentAnalysisToInsert.CommentID = uint(commentToUpdate.ID)
	sentimentAnalysisToInsert.ConfidenceLevel = sentiment.Compound
	sentimentInsert := service.commentRepository.InsertToSentimentAnalysis(sentimentAnalysisToInsert)

	var commentDTO dto.CommentUpdateDTO
	commentDTO.ID = commentUpdated.ID
	commentDTO.SentimentType = sentimentInsert.SentimentType
	commentDTO.ConfidenceLevel = sentimentInsert.ConfidenceLevel
	commentDTO.Comment = commentUpdated.Comment
	commentDTO.CourseID = commentUpdated.CourseID
	commentDTO.LecturerID = commentUpdated.LecturerID
	commentDTO.SchoolYear = commentUpdated.SchoolYear
	commentDTO.Semester = commentUpdated.Semester
	return commentDTO
}

//GetAllComment digunakan untuk mendapatkan semua komentar dari database.
func (service *commentService) GetAllComment() []dto.CommentDTO {
	allComment := service.commentRepository.GetAllComment()

	var commentDTO []dto.CommentDTO
	for _, value := range allComment {
		var comment dto.CommentDTO
		lecturerData := service.lecturerRepository.GetLecturerByID(value.LecturerID)
		comment.CommentID = value.ID
		comment.Comment = value.Comment
		comment.CourseID = value.CourseID
		comment.CourseCode = value.Course.Code
		comment.CourseName = value.Course.Name
		comment.LecturerID = value.LecturerID
		comment.LecturerCode = lecturerData.Code
		comment.LecturerName = lecturerData.Name
		comment.StudyProgramID = lecturerData.StudyProgramID
		comment.StudyProgramCode = lecturerData.StudyProgram.Code
		comment.StudyProgramName = lecturerData.StudyProgram.Name
		comment.SchoolYear = value.SchoolYear
		comment.Semester = value.Semester
		commentDTO = append(commentDTO, comment)
	}

	return commentDTO
}

//GetCommentByID digunakan untuk mendapatkan komentar berdasarkan ID dari database.
func (service *commentService) GetCommentByID(id uint64) dto.CommentDTO {
	comment := service.commentRepository.GetCommentByID(id)

	var commentDTO dto.CommentDTO
	lecturerData := service.lecturerRepository.GetLecturerByID(comment.LecturerID)
	commentDTO.CommentID = comment.ID
	commentDTO.Comment = comment.Comment
	commentDTO.CourseID = comment.CourseID
	commentDTO.CourseCode = comment.Course.Code
	commentDTO.CourseName = comment.Course.Name
	commentDTO.LecturerID = comment.LecturerID
	commentDTO.LecturerCode = lecturerData.Code
	commentDTO.LecturerName = lecturerData.Name
	commentDTO.StudyProgramID = lecturerData.StudyProgramID
	commentDTO.StudyProgramCode = lecturerData.StudyProgram.Code
	commentDTO.StudyProgramName = lecturerData.StudyProgram.Name
	commentDTO.SchoolYear = comment.SchoolYear
	commentDTO.Semester = comment.Semester

	return commentDTO
}

//DeleteComment digunakan untuk menghapus komentar dari database berdasarkan ID.
func (service *commentService) DeleteComment(id uint64) error {
	err := service.commentRepository.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}
