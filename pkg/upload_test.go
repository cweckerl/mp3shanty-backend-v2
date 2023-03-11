package pkg

import "testing"

func TestUpload(t *testing.T) {
	path := "/tmp/song.mp3"
	uploader := MockUploader{}
	expected := RepositoryId("testId")

	repoId := uploader.upload(path)

	if expected != repoId {
		t.Fatalf("Expected %q, got %q", expected, repoId)
	}
}

func TestGetUrl(t *testing.T) {
	repoId := RepositoryId("testId")
	uploader := MockUploader{}
	expected := RepositoryUrl("fakeUrl")

	repoUrl := uploader.getUrl(repoId)

	if expected != repoUrl {
		t.Fatalf("Expected %q, got %q", expected, repoUrl)
	}
}
