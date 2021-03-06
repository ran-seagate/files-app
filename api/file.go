package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ran-seagate/files-app/config"
	"github.com/ran-seagate/files-app/entities"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// CreateUploadFolder creates upload folder on the root dir
func CreateUploadFolder() error {
	if err := os.MkdirAll(filepath.Join(".", config.AppConfig.UploadFolder), os.ModePerm); err != nil {
		return err
	}

	fmt.Println("[CreateUploadFolder]: successfully created upload folder")
	return nil
}

// GetFilesList gets list of files on the upload dir
func GetFilesList(c *gin.Context) {
	files, err := ioutil.ReadDir(config.AppConfig.UploadFolder)
	if err != nil {
		fmt.Printf("[GetFilesList]: %s\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var filesDetails []*entities.FileDetails
	for _, file := range files {
		d := file.Sys().(*syscall.Stat_t)
		fileDetails := &entities.FileDetails{
			Name: file.Name(),
			Ext:  file.Name()[strings.LastIndex(file.Name(), "."):],
			// For mac:
			CreationDate: time.Unix(d.Ctimespec.Sec, d.Ctimespec.Nsec),
			// For linux:
			//CreationDate: time.Unix(int64(d.Ctim.Sec), int64(d.Ctim.Nsec)),
			Size: uint64(file.Size()),
		}
		filesDetails = append(filesDetails, fileDetails)
	}

	c.JSON(http.StatusOK, filesDetails)
}

// GetFile gets a file by its name as a path url
func GetFile(c *gin.Context) {
	fileName := c.Param("file_name")
	filePath := filepath.Join(".", config.AppConfig.UploadFolder, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("[UploadFile]: %s\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File doesn't exist", "file_name": fileName})
		return
	}

	c.FileAttachment(filePath, fileName)
}

// DeleteFile deletes a file from the local upload dir
func DeleteFile(c *gin.Context) {
	fileName := c.Param("file_name")
	filePath := filepath.Join(".", config.AppConfig.UploadFolder, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("[UploadFile]: %s\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File doesn't exist", "file_name": fileName})
		return
	}

	if err := os.Remove(filePath); err != nil {
		fmt.Printf("[UploadFile]: %s\n", err)
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"error": "Failed to remove file: " + err.Error(), "file_name": fileName})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// UploadFile uploads a file to the upload dir
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("[UploadFile]: %s\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from request: " + err.Error()})
		return
	}

	err = c.SaveUploadedFile(file, filepath.Join(".", config.AppConfig.UploadFolder, file.Filename))
	if err != nil {
		fmt.Printf("[UploadFile]: %s\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to save file from request: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File " + file.Filename + " Uploaded successfully"})
}
