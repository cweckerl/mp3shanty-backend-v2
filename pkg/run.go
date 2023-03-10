package pkg

import "sync"

func Run(id VideoId, downloader Downloader, uploader Uploader) RepositoryUrl {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		downloader.downloadThumbnail(id)
	}()

	go func() {
		defer wg.Done()
		downloader.downloadVideo(id)
	}()

	wg.Wait()

	err := downloader.embedThumbnail(id)

	var path string
	if err != nil {
		path = getOutputFilePath(id)
	} else {
		path = getTempFilePath(id)
	}
	repoId := uploader.upload(path)
	return uploader.getUrl(repoId)
}
