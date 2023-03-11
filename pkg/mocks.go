package pkg

type MockDownloader struct{}
type MockUploader struct{}

func (d MockDownloader) downloadVideo(id VideoId)        {}
func (d MockDownloader) downloadThumbnail(id VideoId)    {}
func (d MockDownloader) embedThumbnail(id VideoId) error { return nil }

func (u MockUploader) upload(path string) RepositoryId {
	return RepositoryId("testId")
}

func (u MockUploader) getUrl(repoId RepositoryId) RepositoryUrl {
	return RepositoryUrl("fakeUrl")
}
