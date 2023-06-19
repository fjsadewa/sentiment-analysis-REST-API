package controllers

import (
	"net/http"
	"sentiment/dto"
	"sentiment/helper"
	"sentiment/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
SentimenAnalysisController adalah interface yang menyediakan operasi-operasi untuk mengelola data analisis sentimen.
*/
type SentimenAnalysisController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAllData(ctx *gin.Context)
	GetDataByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

//sentimenAnalysisController adalah implementasi dari SentimenAnalysisController.
type sentimenAnalysisController struct {
	sentimenAnalysisService service.SentimenAnalysisService
	jwtService              service.JWTService
}

/*
NewSentimenAnalysisController digunakan untuk membuat instance baru dari sentimenAnalysisController.
Fungsi ini menerima sentimenAnalysisService dan jwtService yang digunakan oleh controller.
Fungsi ini mengembalikan instance baru dari SentimenAnalysisController.
*/
func NewSentimenAnalysisController(sentimenAnalysisService service.SentimenAnalysisService, jwtService service.JWTService) SentimenAnalysisController {
	return &sentimenAnalysisController{
		sentimenAnalysisService: sentimenAnalysisService,
		jwtService:              jwtService,
	}
}

//Create digunakan untuk membuat data analisis sentimen baru.
func (c *sentimenAnalysisController) Create(ctx *gin.Context) {
	var sentimenAnalysisCreate dto.SentimenAnalysisCreateDTO
	errDTO := ctx.ShouldBind(&sentimenAnalysisCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	sentimenAnalysis := c.sentimenAnalysisService.InsertSentimenAnalysis(sentimenAnalysisCreate)
	res := helper.BuildResponse(true, "OK!", sentimenAnalysis)
	ctx.JSON(http.StatusOK, res)
}

//Update digunakan untuk memperbarui data analisis sentimen yang sudah ada.
func (c *sentimenAnalysisController) Update(ctx *gin.Context) {
	var sentimenAnalysisUpdateDTO dto.SentimenAnalysisUpdateDTO
	errDTO := ctx.ShouldBind(&sentimenAnalysisUpdateDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	sentimenAnalysis := c.sentimenAnalysisService.UpdateSentimenAnalysis(sentimenAnalysisUpdateDTO)
	res := helper.BuildResponse(true, "OK!", sentimenAnalysis)
	ctx.JSON(http.StatusOK, res)
}

//GetAllData digunakan untuk mengambil semua data analisis sentimen.
func (c *sentimenAnalysisController) GetAllData(ctx *gin.Context) {
	sentimenAnalysiss := c.sentimenAnalysisService.GetAllSentimenAnalysis()
	res := helper.BuildResponse(true, "OK!", sentimenAnalysiss)
	ctx.JSON(http.StatusOK, res)
}

//GetDataByID digunakan untuk mengambil data analisis sentimen berdasarkan ID.
func (c *sentimenAnalysisController) GetDataByID(ctx *gin.Context) {
	id := ctx.Param("id")
	sentimenAnalysisID, _ := strconv.ParseUint(id, 0, 0)
	sentimenAnalysis := c.sentimenAnalysisService.GetSentimenAnalysisByID(sentimenAnalysisID)
	if sentimenAnalysis.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "Data not found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", sentimenAnalysis)
	ctx.JSON(http.StatusOK, res)
}

//Delete digunakan untuk menghapus data analisis sentimen berdasarkan ID.
func (c *sentimenAnalysisController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	sentimenAnalysisID, _ := strconv.ParseUint(id, 0, 0)
	err := c.sentimenAnalysisService.DeleteSentimenAnalysis(sentimenAnalysisID)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
