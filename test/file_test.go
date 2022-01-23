package test

import (
	"encoding/json"
	"files-app/entities"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"
)

const testFileName = "testFile.txt"

func TestGetFilesListNoFiles(t *testing.T) {
	router, w := setup()
	defer deleteTestUploadFolder()
	req := httptest.NewRequest("GET", "/files/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	respBody := []*entities.FileDetails{}
	err := json.Unmarshal(w.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, len(respBody), 0)
}

func TestGetFileNotExist(t *testing.T) {
	router, w := setup()
	defer deleteTestUploadFolder()
	req := httptest.NewRequest("GET", "/files/"+testFileName, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestGetFileExists(t *testing.T) {
	router, w := setup()
	defer deleteTestUploadFolder()
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer writer.Close()
		part, err := writer.CreateFormFile("file", testFileName)
		assert.NoError(t, err)
		file, _ := os.Open("../README.md")
		defer file.Close()
		_, err = io.Copy(part, file)
		if err != nil {
			t.Error(err)
		}
	}()

	w.Body.Reset()
	request := httptest.NewRequest("POST", "/files/", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	router.ServeHTTP(w, request)
	assert.Equal(t, 200, w.Code)

	w.Body.Reset()
	req := httptest.NewRequest("GET", "/files/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	respBody := []*entities.FileDetails{}
	err := json.Unmarshal(w.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.NotNil(t, respBody)
	assert.Equal(t, len(respBody), 1)
	assert.NotNil(t, respBody[0])
	assert.Equal(t, respBody[0].Name, testFileName)
}

func TestDeleteFileNotExist(t *testing.T) {
	router, w := setup()
	defer deleteTestUploadFolder()
	req := httptest.NewRequest("DELETE", "/files/"+testFileName, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestDeleteFileExists(t *testing.T) {
	router, w := setup()
	defer deleteTestUploadFolder()
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer writer.Close()
		part, err := writer.CreateFormFile("file", testFileName)
		assert.NoError(t, err)
		file, _ := os.Open("../README.md")
		defer file.Close()
		_, err = io.Copy(part, file)
		if err != nil {
			t.Error(err)
		}
	}()

	request := httptest.NewRequest("POST", "/files/", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	router.ServeHTTP(w, request)
	assert.Equal(t, 200, w.Code)

	w.Body.Reset()
	req := httptest.NewRequest("DELETE", "/files/"+testFileName, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
