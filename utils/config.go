package utils

import (
	"log"
	"strings"
)

// 从指定配置文件中获取指定键的值
func GetValueByConfig(filePath, key string) (value string) {
	if exist := CheckFileExist(filePath); !exist {
		log.Fatalf("[GetValueByConfig] [CheckFileExist] file %s not exist!", filePath)
		return ""
	}
	for _, line := range ReadFileLines(filePath) {
		if line == "" {
			continue
		}
		subStrings := strings.Split(line, ":")
		if len(subStrings) < 2 || subStrings[0] != key {
			continue
		}
		return subStrings[1]
	}
	return ""
}
