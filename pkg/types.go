package pkg

import "google.golang.org/api/drive/v3"

type VideoId string
type RepositoryId string
type RepositoryUrl string

type YtDlpDownloader struct{}
type DriveUploader struct {
	srv *drive.Service
}

type DownloadService interface {
	DownloadVideo(VideoId)
	DownloadThumbnail(VideoId)
	EmbedThumbnail(VideoId)
}

type UploadService interface {
	Upload(string) RepositoryId
	GetUrl(RepositoryId) RepositoryUrl
}

const TMP = "/tmp"
const CREDS = "creds/"
