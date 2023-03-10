package pkg

type MockDownloader struct{}
type MockUploader struct{}

func (u MockUploader) upload(path string) RepositoryId {
	return RepositoryId("testId")
}

func (u MockUploader) getUrl(repoId RepositoryId) RepositoryUrl {
	return RepositoryUrl("fakeUrl")
}
