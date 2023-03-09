package pkg

import "testing"

func TestGetVideoUrl(t *testing.T) {
	id := VideoId("abc")
	expected := "https://www.youtube.com/watch?v=abc"

	url := getVideoUrl(id)

	if expected != url {
		t.Fatalf("getVideoUrl(abc) != https://www.youtube.com/watch?v=abc")
	}
}

func TestGetThumbnailUrl(t *testing.T) {
	id := VideoId("abc")
	expected := "https://img.youtube.com/vi/abc/maxresdefault.jpg"

	url := getThumbnailUrl(id)

	if expected != url {
		t.Fatalf("getThumbnailUrl(abc) != https://img.youtube.com/vi/abc/maxresdefault.jpg")
	}
}

func TestGetTempFilePath(t *testing.T) {
	id := VideoId("abc")
	expected := "/tmp/tmp_abc.mp3"

	path := getTempFilePath(id)

	if expected != path {
		t.Fatalf("getTempFilePath(abc) != /tmp/tmp_abc.mp3")
	}
}

func TestGetOutputFilePath(t *testing.T) {
	id := VideoId("abc")
	expected := "/tmp/abc.mp3"

	path := GetOutputFilePath(id)

	if expected != path {
		t.Fatalf("GetOutputFilePath(abc) != /tmp/abc.mp3")
	}
}

func TestThumbnailFilePath(t *testing.T) {
	id := VideoId("abc")
	expected := "/tmp/abc.jpg"

	path := getThumbnailFilePath(id)

	if expected != path {
		t.Fatalf("getThumbnailFilePath(abc) != /tmp/abc.jpg")
	}
}
