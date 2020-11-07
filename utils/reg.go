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

func ReplaceUrlFromOneLine(lineText string, clouds []string) string {
	regStr := `!\[.*?\]\(.*?\)`
	reg := regexp.MustCompile(regStr)
	if reg == nil {
		log.Fatal("regexp error")
		return ""
	}
	result := reg.FindAllString(lineText, -1)
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
		toUploadFilePath := result[i][left+1 : right]

		if strings.HasPrefix(toUploadFilePath, "http://") || strings.HasPrefix(toUploadFilePath, "http") {
			log.Printf("File %s is already on Internet", toUploadFilePath)
			continue
		}
		if !CheckFileExist(toUploadFilePath) {
			log.Printf("File %s is not exist!", toUploadFilePath)
			continue
		}
		repStr := strings.ReplaceAll(result[i], toUploadFilePath, "666")
		lineText = strings.ReplaceAll(lineText, result[i], repStr)
	}
	return lineText
}
