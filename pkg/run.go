package pkg

import "sync"

func Run(id VideoId, downloader DownloadService, uploader UploadService) RepositoryUrl {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		downloader.DownloadThumbnail(id)
	}()

	go func() {
		defer wg.Done()
		downloader.DownloadVideo(id)
	}()

	wg.Wait()

	downloader.EmbedThumbnail(id)

	path := GetOutputFilePath(id)
	repoId := uploader.Upload(path)
	return uploader.GetUrl(repoId)
}
