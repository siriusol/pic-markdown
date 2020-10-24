package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Get(filePath, key string) (value string) {
	if exist := CheckFileExist(filePath); !exist {
		log.Printf("[Get] [CheckFileExist] file %s not exist!", filePath)
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

func ReadFileLines(filename string) []string {
	lines := make([]string, 0, 10)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return lines
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lines
}
