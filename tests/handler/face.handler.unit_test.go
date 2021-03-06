package handler_test

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"goface-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Delete(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/face/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("23124112312")

	t.Run("Happy", func(t *testing.T) {
		mockRepoFace.On("DeleteId", "23124112312").Return(nil).Once()

		if assert.NoError(t, h.Delete(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("DBError", func(t *testing.T) {
		mockRepoFace.On("DeleteId", "23124112312").Return(errors.New("delete id error")).Once()

		// Assertions
		errHandler := h.Delete(c).(*echo.HTTPError)
		assert.Equal(t, http.StatusInternalServerError, errHandler.Code)
	})

}

func TestHandler_FaceAll(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/face/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	t.Run("Happy", func(t *testing.T) {
		mockRepoFace.On("FindAll").Return([]models.Face{faceData}, nil).Once()

		// Assertions
		if assert.NoError(t, h.FaceAll(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("FindAllErr", func(t *testing.T) {
		mockRepoFace.On("FindAll").Return([]models.Face{faceData}, errors.New("FindAll err")).Once()

		// Assertions
		errHandler := h.FaceAll(c).(*echo.HTTPError)
		assert.Equal(t, http.StatusInternalServerError, errHandler.Code)
	})

}

func TestHandler_FaceId(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/api/face/:id", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("13213213")


	t.Run("Happy", func(t *testing.T) {
		mockRepoFace.On("FindById", "13213213").Return([]models.Face{faceData}, nil).Once()
		// Assertions
		if assert.NoError(t, h.FaceId(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("FindByIdErr", func(t *testing.T) {
		mockRepoFace.On("FindById", "13213213").Return([]models.Face{faceData}, errors.New("FindByIdErr")).Once()
		// Assertions
		errHandler := h.FaceId(c).(*echo.HTTPError)
		assert.Equal(t, http.StatusInternalServerError, errHandler.Code)
	})

}
