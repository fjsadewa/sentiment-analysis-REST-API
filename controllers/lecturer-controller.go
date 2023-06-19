package controllers

import (
	"net/http"
	"sentiment/dto"
	"sentiment/helper"
	"sentiment/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//LecturerController adalah interface yang menyediakan operasi-operasi untuk mengelola data dosen.
type LecturerController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAllData(ctx *gin.Context)
	GetDataByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

//lecturerController adalah implementasi dari LecturerController.
type lecturerController struct {
	lecturerService service.LecturerService
	jwtService      service.JWTService
}

/*
NewLecturerController digunakan untuk membuat instance baru dari lecturerController.
Fungsi ini menerima lecturerService dan jwtService yang digunakan oleh controller.
Fungsi ini mengembalikan instance baru dari LecturerController.
*/
func NewLecturerController(lecturerService service.LecturerService, jwtService service.JWTService) LecturerController {
	return &lecturerController{
		lecturerService: lecturerService,
		jwtService:      jwtService,
	}
}

/*
Create digunakan untuk membuat data dosen baru.
Fungsi ini menerima request body berupa LecturerCreateDTO dan mengembalikan respons berupa data dosen yang berhasil dibuat.
*/
func (c *lecturerController) Create(ctx *gin.Context) {
	var lecturerCreate dto.LecturerCreateDTO
	errDTO := ctx.ShouldBind(&lecturerCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	lecturer := c.lecturerService.InsertLecturer(lecturerCreate)
	res := helper.BuildResponse(true, "OK!", lecturer)
	ctx.JSON(http.StatusOK, res)
}

/*
Update digunakan untuk memperbarui data dosen yang sudah ada.
Fungsi ini menerima request body berupa LecturerUpdateDTO dan mengembalikan respons berupa data dosen yang berhasil diperbarui.
*/
func (c *lecturerController) Update(ctx *gin.Context) {
	var lecturerUpdateDTO dto.LecturerUpdateDTO
	errDTO := ctx.ShouldBind(&lecturerUpdateDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	lecturer := c.lecturerService.UpdateLecturer(lecturerUpdateDTO)
	res := helper.BuildResponse(true, "OK!", lecturer)
	ctx.JSON(http.StatusOK, res)
}

/*
GetAllData digunakan untuk mendapatkan semua data dosen.
Fungsi ini mengembalikan respons berupa daftar data dosen yang berhasil ditemukan.
*/
func (c *lecturerController) GetAllData(ctx *gin.Context) {
	lecturers := c.lecturerService.GetAllLecturer()
	res := helper.BuildResponse(true, "OK!", lecturers)
	ctx.JSON(http.StatusOK, res)
}

/*
GetDataByID digunakan untuk mendapatkan data dosen berdasarkan ID.
Fungsi ini menerima parameter ID sebagai path parameter dan mengembalikan respons berupa data dosen yang berhasil ditemukan.
*/
func (c *lecturerController) GetDataByID(ctx *gin.Context) {
	id := ctx.Param("id")
	lecturerID, _ := strconv.ParseUint(id, 0, 0)
	lecturer := c.lecturerService.GetLecturerByID(lecturerID)
	if lecturer.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "Data not found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", lecturer)
	ctx.JSON(http.StatusOK, res)
}

/*
Delete digunakan untuk menghapus data dosen berdasarkan ID.
Fungsi ini menerima parameter ID sebagai path parameter dan mengembalikan respons kosong jika penghapusan berhasil.
*/
func (c *lecturerController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	lecturerID, _ := strconv.ParseUint(id, 0, 0)
	err := c.lecturerService.DeleteLecturer(lecturerID)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
