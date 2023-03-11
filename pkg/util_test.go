package pkg

import "testing"

func TestGetVideoUrl(t *testing.T) {
	id := VideoId("abc")
	expected := "https://www.youtube.com/watch?v=abc"

	url := getVideoUrl(id)

	if expected != url {
		t.Fatalf("Expected %q, got %q", expected, url)
	}
}

func TestGetThumbnailUrl(t *testing.T) {
	id := VideoId("abc")
	expected := "https://img.youtube.com/vi/abc/maxresdefault.jpg"

	url := getThumbnailUrl(id)

	if expected != url {
		t.Fatalf("Expected %q, got %q", expected, url)
	}
}

func TestGetTempFilePath(t *testing.T) {
	id := VideoId("abc")
	expected := "/tmp/tmp_abc.mp3"

	path := getTempFilePath(id)

	if expected != path {
		t.Fatalf("Expected %q, got %q", expected, path)
	}
}

func TestGetOutputFilePath(t *testing.T) {
	id := VideoId("abc")
	expected := "/tmp/abc.mp3"

	path := getOutputFilePath(id)

	if expected != path {
		t.Fatalf("Expected %q, got %q", expected, path)
	}
}

func TestThumbnailFilePath(t *testing.T) {
	id := VideoId("abc")
	expected := "/tmp/abc.jpg"

	path := getThumbnailFilePath(id)

	if expected != path {
		t.Fatalf("Expected %q, got %q", expected, path)
	}
}
