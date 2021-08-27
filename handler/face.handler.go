package handler

import (
	"context"
	"goface-api/models"
	"goface-api/response"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) Delete(c echo.Context) error {
	id := c.Param("id")

	modelFace := models.Face{Id: id}
	res,err := modelFace.Delete(context.Background(), h.DB)
	if err != nil {
		log.Println("delete error:", err)
		return c.JSON(http.StatusBadRequest, response.Response{Error: err.Error()})
	}
	log.Println("delete count:",res.DeletedCount)
	if res.DeletedCount>0 {
		return c.JSON(http.StatusOK, response.Response{Detail: "deleted"})
	}else {
		return c.JSON(http.StatusInternalServerError, response.Response{})
	}
}

func (h Handler) FaceAll(c echo.Context) error {
	modelFace := models.Face{}
	faces, err := modelFace.FindAll(context.Background(), h.DB)
	if err != nil {
		log.Println("FindAll error:", err)
		return c.JSON(http.StatusBadRequest, response.Response{Error: err.Error()})
	}

	resFaces := []response.FaceLenDesc{}
	for _, face := range faces {
		resFaces = append(resFaces, response.FaceLenDesc{Id: face.Id, Name: face.Name, Descriptors: len(face.Descriptors)})
	}
	log.Println("FindAll success!")
	return c.JSON(http.StatusOK, response.Response{Data: resFaces})
}