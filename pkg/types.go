package pkg

import "google.golang.org/api/drive/v3"

type VideoId string
type RepositoryId string
type RepositoryUrl string

type YtDlpDownloader struct{}
type DriveUploader struct {
	srv *drive.Service
}

type Downloader interface {
	downloadVideo(VideoId)
	downloadThumbnail(VideoId)
	embedThumbnail(VideoId) error
}

type Uploader interface {
	upload(string) RepositoryId
	getUrl(RepositoryId) RepositoryUrl
}

const TMP = "/tmp"
const CREDS = "creds/"
