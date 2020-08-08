package utils

import (
	"log"
	"regexp"
	"strings"
)

func GetUrlFromOneLine(str string) []string {
	regStr := `!\[.*?\]\(.*?\)`
	reg := regexp.MustCompile(regStr)
	if reg == nil {
		log.Fatal("regexp error")
		return nil
	}
	result := reg.FindAllString(str, -1)
	var urls []string
	for _, value := range result {
		subRegStr := `\(.*?\)`
		subReg := regexp.MustCompile(subRegStr)
		if subReg == nil {
			log.Fatal("sub regexp error")
			return nil
		}
		url := subReg.FindString(value)
		urls = append(urls, url[1:len(url)-1])
	}
	return urls
}

func ReplaceUrlFromOneLine(url string, clouds []string) string {
	regStr := `!\[.*?\]\(.*?\)`
	reg := regexp.MustCompile(regStr)
	if reg == nil {
		log.Fatal("regexp error")
		return ""
	}
	result := reg.FindAllString(url, -1)
	for i := 0; i < len(result); i++ {
		var left, right int
		for i, char := range result[i] {
			if char == '(' {
				left = i
			}
			if char == ')' {
				right = i
			}
		}
		uploadFileUrl := result[i][left+1 : right]

		if strings.HasPrefix(uploadFileUrl, "http://") || strings.HasPrefix(uploadFileUrl, "http") {
			log.Println("File is already on Internet.")
			continue
		}
		if !CheckFileExist(uploadFileUrl) {
			log.Printf("File %s is not exist!", uploadFileUrl)
			continue
		}
		UploadFile(uploadFileUrl)

		repStr := strings.ReplaceAll(result[i], uploadFileUrl, "666")
		url = strings.ReplaceAll(url, result[i], repStr)
	}
	return url
}
