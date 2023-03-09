package main

import (
	"mp3shanty/pkg"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(event pkg.EventInput) (pkg.EventOutput, error) {
	videoId := pkg.VideoId(event.VideoId)
	downloader := pkg.YtDlpDownloader{}
	uploader := pkg.NewDriveUploader()
	repoUrl := pkg.Run(videoId, downloader, uploader)
	return pkg.EventOutput{Url: string(repoUrl)}, nil
}
