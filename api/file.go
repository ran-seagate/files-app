package api

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"newEmpTask/config"
	"newEmpTask/entities"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func GetFilesList(c *gin.Context) {
	files, err := ioutil.ReadDir(config.AppConfig.UploadFolder)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var filesDetails []*entities.FileDetails
	for _, file := range files {
		d := file.Sys().(*syscall.Win32FileAttributeData)
		fileDetails := &entities.FileDetails{
			Name:         file.Name(),
			Ext:          strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())),
			CreationDate: time.Unix(0, d.CreationTime.Nanoseconds()),
			Size:         uint64(file.Size()),
		}
		filesDetails = append(filesDetails, fileDetails)
	}

	c.JSON(http.StatusOK, filesDetails)
}

func GetFile(c *gin.Context) {
	fileName := c.Param("file_name")
	filePath := filepath.Join(".", config.AppConfig.UploadFolder, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File doesn't exist", "file_name": fileName})
		return
	}

	c.FileAttachment(filePath, fileName)
	c.JSON(http.StatusOK, nil)
}

func DeleteFile(c *gin.Context) {
	fileName := c.Param("file_name")
	filePath := filepath.Join(".", config.AppConfig.UploadFolder, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File doesn't exist", "file_name": fileName})
		return
	}

	err := os.Remove(filePath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"error": "Failed to remove file", "file_name": fileName})
		return
	}
	c.JSON(http.StatusOK, nil)
}
