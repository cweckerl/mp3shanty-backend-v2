build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main cmd/main.go

zip:
	zip main.zip main creds/creds1.json creds/creds2.json creds/creds3.json bin/yt-dlp

clean:
	rm main.zip main

main: build zip
