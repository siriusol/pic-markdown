package utils

import (
	"log"
	"regexp"
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
		urls = append(urls, url[1:len(url) - 1])
	}
	return urls
}
