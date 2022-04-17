package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "token")
	os.Setenv("CHANNEL_ID", "id")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))  //создаем новый slack.Client
	channelArr := []string{os.Getenv("CHANNEL_ID")}  //масив каналов
	fileArr := []string{"cert-25447957-1073.pdf"}  //масив файлов

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{  //задаем параметры загрузки
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)  //загружаем файл, получаем информацию о файле и ошибку(по умолчанию nil)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}