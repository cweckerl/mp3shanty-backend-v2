package pkg

import (
	"fmt"
	"path/filepath"
)

func getVideoUrl(id VideoId) string {
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", id)
}

func getThumbnailUrl(id VideoId) string {
	return fmt.Sprintf("https://img.youtube.com/vi/%s/maxresdefault.jpg", id)
}

func getTempFilePath(id VideoId) string {
	return filepath.Join(TMP, fmt.Sprintf("tmp_%s.mp3", id))
}

func GetOutputFilePath(id VideoId) string {
	return filepath.Join(TMP, fmt.Sprintf("%s.mp3", id))
}

func getThumbnailFilePath(id VideoId) string {
	return filepath.Join(TMP, fmt.Sprintf("%s.jpg", id))
}
