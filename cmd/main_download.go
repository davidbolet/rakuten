package cmd

import (
	"log"
	"strconv"
	"time"

	"cursmedia.com/rakuten/utils"
)

const fileName = "./xml/result"

// Download downloads data from linkshare
func Download() {

	log.Printf("Processing Rakuten Udemy Affiliate API")
	log.Printf("1- Request token")
	rakutenClient := utils.NewRakutenClient()
	for i := 1; i < 1149; i++ {

		go callAndProcess(&rakutenClient, i)
		time.Sleep(time.Second * 65)
		log.Printf("Rakuten productSearch retrieved")
	}
}

func callAndProcess(client *utils.RakutenClient, i int) {
	content, err := client.MakeRequest(i)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fileN := fileName + time.Now().Format("2006-01-02 15:04:05.000000000") + strconv.Itoa(i) + ".xml"
	ProcessStringContent(content)
	utils.SaveToFile(fileN, &content)
}
