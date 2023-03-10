package pkg

import (
	"log"
	"os/exec"
)

func (d YtDlpDownloader) downloadVideo(id VideoId) {
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

func (d YtDlpDownloader) downloadThumbnail(id VideoId) {
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

func (d YtDlpDownloader) embedThumbnail(id VideoId) error {
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
		getOutputFilePath(id),
	)

	if err := cmd.Run(); err != nil {
		log.Printf("Failed to embed thumbnail for %s: %v", id, err)
		return err
	}
	return nil
}
