package pkg

import "testing"

func TestRun(t *testing.T) {
	id := VideoId("fakeId")
	downloader := MockDownloader{}
	uploader := MockUploader{}
	expected := RepositoryUrl("fakeUrl")

	repoUrl := Run(id, downloader, uploader)

	if expected != repoUrl {
		t.Fatalf("Expected %q, got %q", expected, repoUrl)
	}
}
