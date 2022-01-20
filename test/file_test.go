package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"newEmpTask/api"
	"testing"
)

func TestGetFilesList(t *testing.T) {
	defer deleteTestUploadFolder()
	router := api.InitRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/files/", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
