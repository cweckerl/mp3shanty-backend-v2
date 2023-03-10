package pkg

type MockDownloader struct{}
type MockUploader struct{}

func (d MockDownloader) DownloadVideo(id VideoId)     {}
func (d MockDownloader) DownloadThumbnail(id VideoId) {}
func (d MockDownloader) EmbedThumbnail(id VideoId)    {}

func (u MockUploader) Upload(path string) RepositoryId {
	return RepositoryId("testId")
}

func (u MockUploader) GetUrl(repoId RepositoryId) RepositoryUrl {
	return RepositoryUrl("fakeUrl")
}
