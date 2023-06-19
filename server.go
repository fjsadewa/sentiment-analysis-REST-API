package main

import (
	"os"
	"sentiment/config"
	"sentiment/controllers"
	"sentiment/middleware"
	"sentiment/repository"
	"sentiment/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                     *gorm.DB                               = config.SetupDatabaseConnection()
	userRepository         repository.UserRepository              = repository.NewUserRepository(db)
	studyProgramRepository repository.StudyProgramRepository      = repository.NewStudyProgramRepository(db)
	lecturerRepository     repository.LecturerRepository          = repository.NewLecturerRepository(db)
	courseRepository       repository.CourseRepository            = repository.NewCourseRepository(db)
	commentRepository      repository.CommentRepository           = repository.NewCommentRepository(db)
	sentimentAnalysisRepo  repository.SentimenAnalysisRepository  = repository.NewSentimenAnalysisRepository(db)
	jwtService             service.JWTService                     = service.NewJWTService()
	authService            service.AuthService                    = service.NewAuthService(userRepository)
	userService            service.UserService                    = service.NewUserService(userRepository)
	studyProgramService    service.StudyProgramService            = service.NewStudyProgramService(studyProgramRepository)
	lecturerService        service.LecturerService                = service.NewLecturerService(lecturerRepository)
	courseService          service.CourseService                  = service.NewCourseService(courseRepository)
	commentService         service.CommentService                 = service.NewCommentServiceWithLecturer(commentRepository, lecturerRepository)
	sentimentAnalysisServ  service.SentimenAnalysisService        = service.NewSentimenAnalysisServiceWithComment(sentimentAnalysisRepo, commentRepository)
	authController         controllers.AuthController             = controllers.NewAuthController(authService, jwtService)
	userController         controllers.UserController             = controllers.NewUserController(userService, jwtService)
	studyProgramController controllers.StudyProgramController     = controllers.NewStudyProgramController(studyProgramService, jwtService)
	lecturerController     controllers.LecturerController         = controllers.NewLecturerController(lecturerService, jwtService)
	courseController       controllers.CourseController           = controllers.NewCourseController(courseService, jwtService)
	commentController      controllers.CommentController          = controllers.NewCommentController(commentService, jwtService)
	sentimentAnalysisCont  controllers.SentimenAnalysisController = controllers.NewSentimenAnalysisController(sentimentAnalysisServ, jwtService)
)

func main() {
	//Mengatur koneksi database
	defer config.CloseDatabaseConnection(db)

	//Mengatur port yang digunakan oleh server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	//Membuat instance router menggunakan gin framework
	router := gin.Default()

	//Rute untuk permintaan root
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	//Rute untuk permintaan otentikasi
	authRoute := router.Group("auth")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
	}

	//Rute untuk permintaan pengguna
	userRoutes := router.Group("users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	//Rute untuk permintaan program studi
	studyProgramRoutes := router.Group("study-program", middleware.AuthorizeJWT(jwtService))
	{
		studyProgramRoutes.POST("/", studyProgramController.Create)
		studyProgramRoutes.PUT("/", studyProgramController.Update)
		studyProgramRoutes.DELETE("/:id", studyProgramController.Delete)
		studyProgramRoutes.GET("/", studyProgramController.All)
		studyProgramRoutes.GET("/:id", studyProgramController.FindByID)
		studyProgramRoutes.GET("/code/:code", studyProgramController.FindByCode)
	}

	//Rute untuk permintaan dosen
	lecturerRoutes := router.Group("lecturer", middleware.AuthorizeJWT(jwtService))
	{
		lecturerRoutes.POST("/", lecturerController.Create)
		lecturerRoutes.PUT("/", lecturerController.Update)
		lecturerRoutes.GET("/", lecturerController.GetAllData)
		lecturerRoutes.GET("/:id", lecturerController.GetDataByID)
		lecturerRoutes.DELETE("/:id", lecturerController.Delete)
	}

	//Rute untuk permintaan kursus
	courseRoutes := router.Group("course", middleware.AuthorizeJWT(jwtService))
	{
		courseRoutes.POST("/", courseController.Create)
		courseRoutes.PUT("/", courseController.Update)
		courseRoutes.GET("/", courseController.GetAllData)
		courseRoutes.GET("/:id", courseController.GetDataByID)
		courseRoutes.DELETE("/:id", courseController.Delete)
	}

	//Rute untuk permintaan komentar
	commentRoutes := router.Group("comment", middleware.AuthorizeJWT(jwtService))
	{
		commentRoutes.POST("/", commentController.Create)
		commentRoutes.PUT("/", commentController.Update)
		commentRoutes.GET("/", commentController.GetAllData)
		commentRoutes.GET("/:id", commentController.GetDataByID)
		commentRoutes.DELETE("/:id", commentController.Delete)
	}

	//Rute untuk permintaan analisis sentimen
	sentimenAnalysisRoutes := router.Group("sentiment-analysis", middleware.AuthorizeJWT(jwtService))
	{
		sentimenAnalysisRoutes.POST("/", sentimentAnalysisCont.Create)
		sentimenAnalysisRoutes.PUT("/", sentimentAnalysisCont.Update)
		sentimenAnalysisRoutes.GET("/", sentimentAnalysisCont.GetAllData)
		sentimenAnalysisRoutes.GET("/:id", sentimentAnalysisCont.GetDataByID)
		sentimenAnalysisRoutes.DELETE("/:id", sentimentAnalysisCont.Delete)
	}

	//Menjalankan server pada port yang telah ditentukan
	router.Run(":" + port)
}
