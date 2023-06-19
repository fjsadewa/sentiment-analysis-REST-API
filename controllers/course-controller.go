package controllers

import (
	"net/http"
	"sentiment/dto"
	"sentiment/helper"
	"sentiment/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//CourseController adalah interface yang mendefinisikan kontrak untuk controller kursus.
type CourseController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAllData(ctx *gin.Context)
	GetDataByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

//courseController adalah implementasi dari CourseController.
type courseController struct {
	courseService service.CourseService
	jwtService    service.JWTService
}

//NewCourseController digunakan untuk membuat instance baru dari courseController.
func NewCourseController(courseService service.CourseService, jwtService service.JWTService) CourseController {
	return &courseController{
		courseService: courseService,
		jwtService:    jwtService,
	}
}

//Create adalah implementasi dari method Create pada interface CourseController.
func (c *courseController) Create(ctx *gin.Context) {
	var courseCreate dto.CourseCreateDTO
	errDTO := ctx.ShouldBind(&courseCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	course := c.courseService.InsertCourse(courseCreate)
	res := helper.BuildResponse(true, "OK!", course)
	ctx.JSON(http.StatusOK, res)
}

//Update adalah implementasi dari method Update pada interface CourseController.
func (c *courseController) Update(ctx *gin.Context) {
	var courseUpdateDTO dto.CourseUpdateDTO
	errDTO := ctx.ShouldBind(&courseUpdateDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	course := c.courseService.UpdateCourse(courseUpdateDTO)
	res := helper.BuildResponse(true, "OK!", course)
	ctx.JSON(http.StatusOK, res)
}

//GetAllData adalah implementasi dari method GetAllData pada interface CourseController.
func (c *courseController) GetAllData(ctx *gin.Context) {
	courses := c.courseService.GetAllCourse()
	res := helper.BuildResponse(true, "OK!", courses)
	ctx.JSON(http.StatusOK, res)
}

//GetDataByID adalah implementasi dari method GetDataByID pada interface CourseController.
func (c *courseController) GetDataByID(ctx *gin.Context) {
	id := ctx.Param("id")
	courseID, _ := strconv.ParseUint(id, 0, 0)
	course := c.courseService.GetCourseByID(courseID)
	if course.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "Data not found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", course)
	ctx.JSON(http.StatusOK, res)
}

//Delete adalah implementasi dari method Delete pada interface CourseController.
func (c *courseController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	courseID, _ := strconv.ParseUint(id, 0, 0)
	err := c.courseService.DeleteCourse(courseID)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
