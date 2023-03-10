package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func getCredentialsFilePath() string {
	return filepath.Join(CREDS, fmt.Sprintf("creds%d.json", rand.Intn(3)+1))
}

func getServiceAccount(path string) *http.Client {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open service account file: %v", err)
	}
	defer file.Close()
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading service account file: %v", err)
	}
	var creds = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(b, &creds)
	config := &jwt.Config{
		Email:      creds.Email,
		PrivateKey: []byte(creds.PrivateKey),
		Scopes:     []string{drive.DriveScope},
		TokenURL:   google.JWTTokenURL,
	}
	client := config.Client(context.Background())
	return client
}

func NewDriveUploader() DriveUploader {
	ctx := context.Background()
	credsPath := getCredentialsFilePath()
	client := getServiceAccount(credsPath)
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	return DriveUploader{srv: srv}
}

func (u DriveUploader) upload(path string) RepositoryId {
	_, id := filepath.Split(path)
	file := drive.File{
		Name:     id,
		MimeType: "audio/mpeg",
	}
	mp3, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open mp3 file: %v", err)
	}
	defer mp3.Close()

	driveFile, err := u.srv.Files.Create(&file).Media(mp3).Fields("id").Do()
	if err != nil {
		log.Fatalf("Failed to create drive file %v", err)
	}

	return RepositoryId(driveFile.Id)
}

func (u DriveUploader) getUrl(repoId RepositoryId) RepositoryUrl {
	perm := drive.Permission{
		Role: "reader",
		Type: "anyone",
	}

	_, err := u.srv.Permissions.Create(string(repoId), &perm).Do()
	if err != nil {
		log.Fatalf("Failed to update permissions for file %s", repoId)
	}
	driveFile, err := u.srv.Files.Get(string(repoId)).Fields("webContentLink").Do()
	if err != nil {
		log.Fatalf("Failed to get file %s", repoId)
	}

	return RepositoryUrl(driveFile.WebContentLink)
}
