package services

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"
)

func SendDocumentByEmail(id, email string) error {
	//Postavke za Gmail SMTP server
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	username := "karla.fisic@gmail.com"
	password := "toyc ebbp eicb aqen"

	//Autentifikacija za SMTP
	auth := smtp.PlainAuth("", username, password, smtpHost)

	// Spremanje putanje datoteke
	filePath := "./uploads/" + id + ".pdf" // Koristi ID za traženje dokumenta
	fileName := filepath.Base(filePath)

	// Postoji li datoteka
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("File does not exist at path: %s", filePath)
			return fmt.Errorf("file does not exist at path: %s", filePath)
		}
		log.Printf("Failed to stat file: %v", err)
		return fmt.Errorf("failed to stat file: %w", err)
	}

	log.Println("File exists:", filePath)

	// Otvaranje datoteke
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Kreiraj MIME tijelo e-maila
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Dodaj zaglavlje za e-mail
	header := make(map[string][]string)
	header["From"] = []string{username}
	header["To"] = []string{email}
	header["Subject"] = []string{"Document"}
	header["Content-Type"] = []string{"multipart/mixed; boundary=" + writer.Boundary()}

	// Dodaj privitak (datoteku) sa sadržajem
	attachmentPart, err := writer.CreateFormFile("attachment", fileName)
	if err != nil {
		log.Printf("Neuspješno stvaranje dijela za privitak: %v", err)
		return fmt.Errorf("neuspješno stvaranje dijela za privitak: %w", err)
	}

	// Zapisivanje datoteke u privitak
	attachmentPart.Write(fileData)

	err = writer.Close()
	if err != nil {
		log.Printf("Neuspješno zatvaranje writer-a: %v", err)
		return fmt.Errorf("neuspješno zatvaranje writer-a: %w", err)
	}

	// Spremi cijeli e-mail u MIME formatu
	message := buf.Bytes()

	// Slanje e-maila putem SMTP-a
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, username, []string{email}, message)
	if err != nil {
		log.Printf("Neuspješno slanje e-maila: %v", err)
		return fmt.Errorf("neuspješno slanje e-maila: %w", err)
	}

	log.Println("E-mail uspješno poslan na", email)
	return nil
}
