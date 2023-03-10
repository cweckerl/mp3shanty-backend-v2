package pkg

import (
	"log"
	"os/exec"
)

func (d YtDlpDownloader) DownloadVideo(id VideoId) {
	cmd := exec.Command(
		"./bin/yt-dlp",
		"-x",
		"--audio-format", "mp3",
		"-o", getTempFilePath(id),
		"--add-metadata",
		getVideoUrl(id),
	)

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to download %s: %v", id, err)
	}
}

func (d YtDlpDownloader) DownloadThumbnail(id VideoId) {
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", getThumbnailUrl(id),
		"-vf", "crop=720:720:280:720",
		getThumbnailFilePath(id),
	)

	if err := cmd.Run(); err != nil {
		log.Printf("Failed to download thumbnail for %s: %v", id, err)
	}
}

func (d YtDlpDownloader) EmbedThumbnail(id VideoId) {
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", getTempFilePath(id),
		"-i", getThumbnailFilePath(id),
		"-map", "0:0",
		"-map", "1:0",
		"-c", "copy",
		"-id3v2_version", "3",
		"-metadata:s:v", "title=\"Album cover\"",
		"-metadata:s:v", "comment=\"Cover (front)\"",
		GetOutputFilePath(id),
	)

	if err := cmd.Run(); err != nil {
		log.Printf("Failed to embed thumbnail for %s: %v", id, err)
	}
}
