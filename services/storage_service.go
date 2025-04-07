package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var documentMap = make(map[string]string)
var mu sync.Mutex

func StoreFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	// Pohranjivanje u lokalni direktorij
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		return "", err
	}

	// Generiranje jedinstvenog ID-a koristeći UUID
	id := uuid.New().String()

	filePath := "./uploads/" + file.Filename

	// Pohranjivanje datoteke
	err := c.SaveUploadedFile(file, filePath)
	if err != nil {
		return "", err
	}

	// Pohranjivanje ID-a i puta datoteke u mapu
	mu.Lock()
	documentMap[id] = filePath
	mu.Unlock()

	return id, nil
}

// GetDocumentByID dohvaća ime dokumenta pomoću ID-a
func GetDocumentByID(id string) (map[string]string, error) {
	mu.Lock()
	defer mu.Unlock()

	filePath, exists := documentMap[id]
	if !exists {
		return nil, fmt.Errorf("document not found")
	}

	fileName := filepath.Base(filePath)

	// Vraćanje ID-a i imena datoteke
	return map[string]string{"id": id, "name": fileName}, nil
}
