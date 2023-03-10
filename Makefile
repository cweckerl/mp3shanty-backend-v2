build:
	go build -o mp3shanty cmd/main.go

zip:
	zip mp3shanty.zip mp3shanty

clean:
	rm mp3shanty mp3shanty.zip

main: build zip
