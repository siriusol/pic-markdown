package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// 读取目录下的单层所有文件
func ReadDirFilesExcludeSubDir(dirPath string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// 检查某个文件是否存在
func CheckFileExist(filename string) bool {
	// TODO IsNotExist 与 IsExist
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func ReplaceOneLine(filename string, picDirPath string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls := GetUrlFromOneLine(scanner.Text())
		if len(urls) != 0 {
			for _, url := range urls {
				fmt.Println(url)
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
