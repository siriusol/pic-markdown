package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

// 读取目录下的所有文件和目录（单层）
func ReadDirFiles(dirPath string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// 检查某个文件是否存在
func CheckFileExist(filename string) bool {
	// TODO IsNotExist 与 IsExist
	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true
}

// 读取指定文件中所有的 URL
func GetAllUrlsFromFile(filename string) (urls []string) {
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
		oneLineUrls := GetUrlFromOneLine(scanner.Text())
		if len(oneLineUrls) != 0 {
			urls = append(urls, oneLineUrls...)
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}

// 读取文件，返回按行组织的字符串切片
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
