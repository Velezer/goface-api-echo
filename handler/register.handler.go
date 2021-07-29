package handler

import (
	"context"
	"goface-api/helper"
	"goface-api/models"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/Kagami/go-face"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type inputValidation struct {
	id   string `validate:"required"`
	name string `validate:"required"`
}

func (h Handler) Register(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")

	validate := validator.New()
	err := validate.Struct(inputValidation{id: id, name: name})
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "fail",
			"detail": err.Error(),
		})
	}

	file, err := c.FormFile("file") //name=file in client html form
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "fail",
			"detail": err.Error(),
		})
	}

	content, err := file.Open()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "fail",
			"detail": err.Error(),
		})
	}
	folderSaved := filepath.Join(helper.ImagesDir, name+"_"+id)
	filename := time.Now().Local().String() + ".jpg"
	filename = strings.Replace(filename, ":", "_", -1)
	helper.SaveFile(folderSaved, filename, content)

	knownFaces, err := helper.RecognizeFile(h.Rec, folderSaved, filename)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "fail",
			"detail": err.Error(),
		})
	}

	dataFace := models.Face{
		Id:          id,
		Name:        name,
		Descriptors: []face.Descriptor{knownFaces[0].Descriptor},
	}

	res, err := dataFace.InsertOne(context.Background(), h.Coll, dataFace)
	if mongo.IsDuplicateKeyError(err) {
		log.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "id already created",
			"detail": "Sukses menambahkan wajah",
		})
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status": "fail",
			"detail": err.Error(),
		})
	}

	log.Println(res)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"data":   dataFace,
		"detail": "Sukses menambahkan wajah",
	})
}
