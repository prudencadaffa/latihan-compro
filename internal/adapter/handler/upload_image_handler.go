package handler

import (
	"bytes"
	"fmt"
	"io"
	"latihan-compro/config"
	"latihan-compro/internal/adapter/handler/response"
	"latihan-compro/internal/adapter/storage"
	"latihan-compro/utils/middleware"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UploadImageInterface interface {
	UploadImage(c echo.Context) error
}

type uploadImage struct {
	storageService storage.SupabaseInterface
}

// UploadImage implements UploadImageInterface.
func (u *uploadImage) UploadImage(c echo.Context) error {
	var (
		respError = response.ErrorResponseDefault{}
		resp      = response.DefaultSuccessResponse{}
	)
	file, err := c.FormFile("file")
	if err != nil {
		log.Errorf("Error getting file: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(400, respError)
	}

	src, err := file.Open()
	if err != nil {
		log.Errorf("Error opening file: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(400, respError)
	}

	defer src.Close()

	fileBuffer := new(bytes.Buffer)
	_, err = io.Copy(fileBuffer, src)
	if err != nil {
		log.Errorf("Error copying file: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(400, respError)
	}

	newFileName := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), getExtension(file.Filename))

	uploadPath := fmt.Sprintf("public/uploads/%s", newFileName)
	url, err := u.storageService.UploadFile(uploadPath, fileBuffer)
	if err != nil {
		log.Errorf("Error uploading file: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(400, respError)
	}

	resp.Meta.Status = true
	resp.Meta.Message = "Success upload image"
	resp.Data = map[string]string{
		"url": url,
	}
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

func getExtension(fileName string) string {
	ext := "." + fileName[len(fileName)-3:] // Ambil 3 karakter terakhir untuk ekstensi
	if len(fileName) > 4 && fileName[len(fileName)-4] == '.' {
		ext = "." + fileName[len(fileName)-4:]
	}
	return ext
}

func NewUploadImage(e *echo.Echo, storageService storage.SupabaseInterface, cfg *config.Config) UploadImageInterface {
	res := &uploadImage{
		storageService: storageService,
	}

	mid := middleware.NewMiddleware(cfg)

	e.POST("/upload-image", res.UploadImage, mid.CheckToken())

	return res
}
