package services

import (
	"errors"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ImageServices(name string, c *gin.Context, image ...string) (string, error) {
	var BaseURL = os.Getenv("BASE_URL")
	file, err := c.FormFile("image")
	if err != nil {
		return "", errors.New("image is required")
	}

	if !isImage(file) {
		return "", errors.New("invalid image format")
	}

	if len(image) > 0 && image[0] != "" {
		if err := RemoveImage(image[0]); err != nil {
			return "", err
		}
	}

	imageFolder := "src/" + name + "/images/"
	uniqueName := uuid.New().String()
	imagePath := imageFolder + uniqueName + file.Filename

	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		return "", errors.New("failed to upload image")
	}

	imagePath = BaseURL + imagePath

	return imagePath, nil
}

func isImage(fileHeader *multipart.FileHeader) bool {
	contentType := fileHeader.Header.Get("Content-Type")
	return strings.HasPrefix(contentType, "image/")
}

func RemoveImage(image string) error {
	var BaseURL = os.Getenv("BASE_URL")
	fileName := strings.TrimPrefix(image, BaseURL)
	print("base url: ", BaseURL)
	if err := os.Remove(fileName); err != nil {
		return errors.New("failed to upload image")
	}
	return nil
}
