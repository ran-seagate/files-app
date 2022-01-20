package api

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())
	files := router.Group("/files")
	{
		files.GET("/", GetFilesList)
		files.GET("/:file_name", GetFile)
		files.POST("/", UploadFile)
		files.DELETE("/:file_name", DeleteFile)
	}

	return router
}
