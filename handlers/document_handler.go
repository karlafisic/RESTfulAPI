package handlers

import (
	"karla/restfulapi/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadDocument(c *gin.Context) {
	// Preuzimanje datoteke iz zahtjeva
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file found"})
		return
	}

	// Pohranjivanje datoteke
	fileID, err := services.StoreFile(c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "id": fileID})
}

func GetDocumentDetails(c *gin.Context) {
	id := c.Param("id")

	// Dohvaćanje detalja o dokumentu
	doc, err := services.GetDocumentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"document": doc})
}

func SendDocument(c *gin.Context) {
	id := c.Param("id") // ID dokumenta (UUID)
	email := c.Query("email")

	log.Printf("SendDocument called with id: %s, email: %s", id, email)

	if id == "" || email == "" {
		log.Println("Missing id or email")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id or email"})
		return
	}

	// Slanje dokumenta
	err := services.SendDocumentByEmail(id, email)
	if err != nil {
		log.Printf("Failed to send document: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send document", "details": err.Error()})
		return
	}

	log.Println("Document sent successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Document sent successfully"})
}
