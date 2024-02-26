package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env variables")
	}
	files := make([]string, 0, 500)
	var ch string
	var path string
	for {
		fmt.Println("Enter the complete file path to upload.")
		fmt.Scanf("%s", &path)
		files = append(files, path)
		fmt.Println("Do You want to Upload more(Y/N)?")
		fmt.Scanf("%s", &ch)
		if ch == "N" || ch == "n" {
			break
		}
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channels := []string{os.Getenv("CHANNEL_ID")}
	for i := 0; i < len(files); i++ {
		vars := slack.FileUploadParameters{
			Channels: channels,
			File:     files[i],
		}
		file, err := api.UploadFile(vars)
		if err != nil {
			fmt.Printf("File Upload Error: %v\n", err)
			return
		}
		fmt.Printf("Name: %v,URL: %v\n", file.Name, file.URLPrivate)
	}
}
