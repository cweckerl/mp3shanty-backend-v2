package pkg

import "testing"

func TestUpload(t *testing.T) {
	path := "/tmp/song.mp3"
	uploader := MockUploader{}
	expected := RepositoryId("testId")

	repoId := uploader.Upload(path)

	if expected != repoId {
		t.Fatalf("uploader.Upload(/tmp/song.mp3) != RepositoryId(testId)")
	}
}

func TestGetUrl(t *testing.T) {
	repoId := RepositoryId("testId")
	uploader := MockUploader{}
	expected := RepositoryUrl("fakeUrl")

	repoUrl := uploader.GetUrl(repoId)

	if expected != repoUrl {
		t.Fatalf("uploader.GetUrl(RepositoryId(testId)) != RepositoryUrl(fakeUrl)")
	}
}
