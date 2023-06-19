package service

import (
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"
)

/*
SentimenAnalysisService adalah sebuah interface yang mendefinisikan method-method
untuk menangani operasi terkait analisis sentimen.
*/
type SentimenAnalysisService interface {
	InsertSentimenAnalysis(sentimenAnalysis dto.SentimenAnalysisCreateDTO) dto.SentimenAnalysisCreateDTO
	UpdateSentimenAnalysis(sentimenAnalysis dto.SentimenAnalysisUpdateDTO) dto.SentimenAnalysisUpdateDTO
	GetAllSentimenAnalysis() []models.SentimentAnalysis
	GetSentimenAnalysisByID(id uint64) models.SentimentAnalysis
	DeleteSentimenAnalysis(id uint64) error
}

//sentimenAnalysisService adalah sebuah struct yang mengimplementasikan interface SentimenAnalysisService.
type sentimenAnalysisService struct {
	sentimenAnalysisRepository repository.SentimenAnalysisRepository
	commentRepository          repository.CommentRepository
}

//NewSentimenAnalysisServiceWithComment membuat sebuah instans baru dari sentimenAnalysisService.
func NewSentimenAnalysisServiceWithComment(sentimenAnalysisRepo repository.SentimenAnalysisRepository, commentRepo repository.CommentRepository) SentimenAnalysisService {
	return &sentimenAnalysisService{
		sentimenAnalysisRepository: sentimenAnalysisRepo,
		commentRepository:          commentRepo,
	}
}

//InsertSentimenAnalysis memasukkan hasil analisis sentimen baru ke dalam database.
func (service *sentimenAnalysisService) InsertSentimenAnalysis(sentimenAnalysis dto.SentimenAnalysisCreateDTO) dto.SentimenAnalysisCreateDTO {
	//Membuat objek model SentimentAnalysis baru dan mengisinya dengan data yang diberikan
	sentimenAnalysisToInsert := models.SentimentAnalysis{}
	sentimenAnalysisToInsert.SentimentType = sentimenAnalysis.SentimentType
	sentimenAnalysisToInsert.CommentID = uint(sentimenAnalysis.CommentID)
	sentimenAnalysisToInsert.ConfidenceLevel = sentimenAnalysis.ConfidenceLevel

	//Memasukkan hasil analisis sentimen ke dalam database menggunakan sentimenAnalysisRepository
	sentimenAnalysisInserted := service.sentimenAnalysisRepository.InsertSentimenAnalysis(sentimenAnalysisToInsert)

	//Membuat objek DTO SentimenAnalysisCreateDTO baru dan mengisinya dengan data hasil analisis sentimen yang dimasukkan
	var sentimenAnalysisDTO dto.SentimenAnalysisCreateDTO
	sentimenAnalysisDTO.SentimentType = sentimenAnalysisInserted.SentimentType
	sentimenAnalysisDTO.ConfidenceLevel = sentimenAnalysisInserted.ConfidenceLevel
	sentimenAnalysisDTO.CommentID = uint64(sentimenAnalysisInserted.CommentID)

	return sentimenAnalysisDTO
}

//UpdateSentimenAnalysis mengubah data hasil analisis sentimen yang ada dalam database.
func (service *sentimenAnalysisService) UpdateSentimenAnalysis(sentimenAnalysis dto.SentimenAnalysisUpdateDTO) dto.SentimenAnalysisUpdateDTO {
	//Membuat objek model SentimentAnalysis baru dan mengisinya dengan data yang diberikan
	sentimenAnalysisToUpdate := models.SentimentAnalysis{}
	sentimenAnalysisToUpdate.ID = uint64(sentimenAnalysis.ID)
	sentimenAnalysisToUpdate.SentimentType = sentimenAnalysis.SentimentType
	sentimenAnalysisToUpdate.CommentID = uint(sentimenAnalysis.CommentID)
	sentimenAnalysisToUpdate.ConfidenceLevel = sentimenAnalysis.ConfidenceLevel

	//Mengubah data hasil analisis sentimen dalam database menggunakan sentimenAnalysisRepository
	sentimenAnalysisUpdated := service.sentimenAnalysisRepository.UpdateSentimenAnalysis(sentimenAnalysisToUpdate)

	//Membuat objek DTO SentimenAnalysisUpdateDTO baru dan mengisinya dengan data hasil analisis sentimen yang diperbarui
	var sentimenAnalysisDTO dto.SentimenAnalysisUpdateDTO
	sentimenAnalysisDTO.ID = sentimenAnalysisUpdated.ID
	sentimenAnalysisDTO.SentimentType = sentimenAnalysisUpdated.SentimentType
	sentimenAnalysisDTO.CommentID = uint64(sentimenAnalysisUpdated.CommentID)
	sentimenAnalysisDTO.ConfidenceLevel = sentimenAnalysisUpdated.ConfidenceLevel
	return sentimenAnalysisDTO
}

//GetAllSentimenAnalysis mengambil semua data hasil analisis sentimen dari database.
func (service *sentimenAnalysisService) GetAllSentimenAnalysis() []models.SentimentAnalysis {
	allSentimenAnalysis := service.sentimenAnalysisRepository.GetAllSentimenAnalysis()
	return allSentimenAnalysis
}

//GetSentimenAnalysisByID mengambil data hasil analisis sentimen berdasarkan ID dari database.
func (service *sentimenAnalysisService) GetSentimenAnalysisByID(id uint64) models.SentimentAnalysis {
	sentimenAnalysis := service.sentimenAnalysisRepository.GetSentimenAnalysisByID(id)
	return sentimenAnalysis
}

//DeleteSentimenAnalysis menghapus data hasil analisis sentimen berdasarkan ID dari database.
func (service *sentimenAnalysisService) DeleteSentimenAnalysis(id uint64) error {
	err := service.sentimenAnalysisRepository.DeleteSentimenAnalysis(id)
	if err != nil {
		return err
	}
	return nil
}
