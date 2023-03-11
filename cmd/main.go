package main

import (
	"log"
	"mp3shanty/pkg"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	videoId := pkg.VideoId(request.QueryStringParameters["videoId"])
	log.Printf("Received request for %s", videoId)
	downloader := pkg.YtDlpDownloader{}
	uploader := pkg.NewDriveUploader()
	repoUrl := pkg.Run(videoId, downloader, uploader)
	log.Printf("%s", repoUrl)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(repoUrl),
	}, nil
}
